package interactor

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	"github.com/ezio1119/fishapp-chat/usecase/presenter"
	"github.com/ezio1119/fishapp-chat/usecase/repository"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ChatInteractor interface {
	CreateChatRoom(ctx context.Context, rm *domain.Room, m *domain.Member) (*chat_grpc.Room, error)
	SendMessage(ctx context.Context, m *domain.Message) (*chat_grpc.Message, error)
	AddMember(ctx context.Context, m *domain.Member) (*chat_grpc.Member, error)
	StreamMessage(rIDs []int64, uID int64, stream chat_grpc.ChatService_StreamMessageServer) error
}

type chatInteractor struct {
	roomRepository    repository.RoomRepository
	queuingRepository repository.QueuingRepository
	chatPresenter     presenter.ChatPresenter
	ctxTimeout        time.Duration
}

func NewChatInteractor(rr repository.RoomRepository, qr repository.QueuingRepository, cp presenter.ChatPresenter, t time.Duration) ChatInteractor {
	return &chatInteractor{rr, qr, cp, t}
}

func (i *chatInteractor) CreateChatRoom(ctx context.Context, rm *domain.Room, m *domain.Member) (*chat_grpc.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	if err := i.roomRepository.CreateRoomAddMember(ctx, rm, m); err != nil {
		return nil, err
	}
	rID := strconv.FormatInt(rm.ID, 10)
	uID := strconv.FormatInt(m.UserID, 10)
	if err := i.queuingRepository.XGroupCreateMkStream(rID, uID); err != nil {
		return nil, err
	}
	return i.chatPresenter.TransformRoomProto(rm)
}

func (i *chatInteractor) AddMember(ctx context.Context, m *domain.Member) (*chat_grpc.Member, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	if _, err := i.roomRepository.GetRoomByID(ctx, m.RoomID); err != nil {
		return nil, err
	}
	if err := i.roomRepository.AddMember(ctx, m); err != nil {
		return nil, err
	}
	rID := strconv.FormatInt(m.RoomID, 10)
	uID := strconv.FormatInt(m.UserID, 10)
	if err := i.queuingRepository.XGroupCreateMkStream(rID, uID); err != nil {
		return nil, err
	}
	return i.chatPresenter.TransformMemberProto(m)
}

func (i *chatInteractor) SendMessage(ctx context.Context, m *domain.Message) (*chat_grpc.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	if _, err := i.roomRepository.GetMemberByUserIDAndRoomID(ctx, m.UserID, m.RoomID); err != nil {
		return nil, err
	}
	if err := i.roomRepository.CreateMessage(ctx, m); err != nil {
		return nil, err
	}
	if err := i.queuingRepository.XADD(m); err != nil {
		return nil, err
	}
	return i.chatPresenter.TransformMessageProto(m)
}

// redisのXREADGROUPでブロック処理を行い、gRPCのstreamで返す
func (i *chatInteractor) StreamMessage(rIDs []int64, uID int64, stream chat_grpc.ChatService_StreamMessageServer) error {
	ctx := stream.Context()

	strUID := strconv.FormatInt(uID, 10)
	if cltID := i.queuingRepository.GetBlockingID(strUID); cltID != 0 {
		return status.Errorf(codes.Unavailable, "user_id=%d is already streaming", uID)
	}

	s := make([]string, len(rIDs)*2)
	for n, rID := range rIDs {
		fmt.Println(uID, rID)
		if _, err := i.roomRepository.GetMemberByUserIDAndRoomID(ctx, uID, rID); err != nil {
			fmt.Println(err)
			return err
		}
		s[n*2] = strconv.FormatInt(rID, 10) // eq. ["5", ">", "12", ">", "32", ">"]
		s[n*2+1] = ">"
	}
	fmt.Println("commmmm")
	eg, ctx := errgroup.WithContext(ctx)
	msgChan := make(chan *domain.Message)

	eg.Go(func() error {
		<-ctx.Done()
		cltID := i.queuingRepository.GetBlockingID(strUID)
		if cltID == 0 {
			return status.Errorf(codes.Internal, "blocking clientID is empty")
		}
		if err := i.queuingRepository.ClientUnblock(cltID); err != nil {
			return err
		}
		i.queuingRepository.DeleteBlockingID(strUID)
		close(msgChan)
		return nil
	})

	eg.Go(func() error {
		for {
			log.Println("NumGoroutine: ", runtime.NumGoroutine())
			msgs, err := i.queuingRepository.XREADGROUP(ctx, strUID, s)
			log.Printf("%#v", len(msgs))
			if err != nil {
				return err
			}
			for _, m := range msgs {
				msgChan <- m
			}
		}
	})

	eg.Go(func() error {
		for m := range msgChan {
			if err := i.chatPresenter.StreamMessage(m, stream); err != nil {
				return err
			}
		}
		return nil
	})

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}
