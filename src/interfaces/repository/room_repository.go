package repository

import (
	"context"
	"database/sql"

	"github.com/ezio1119/fishapp-chat/domain"
	"github.com/ezio1119/fishapp-chat/usecase/repository"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type roomRepository struct {
	conn *gorm.DB
}

func NewRoomRepository(c *gorm.DB) repository.RoomRepository {
	return &roomRepository{c}
}

// リポジトリ内でトランザクションしちゃってる
func (r *roomRepository) CreateRoomAddMember(ctx context.Context, rm *domain.Room, m *domain.Member) error {
	tx := r.conn.BeginTx(ctx, &sql.TxOptions{})
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return status.Error(codes.Aborted, err.Error())
	}
	result := tx.Create(rm)
	if err := result.Error; err != nil {
		tx.Rollback()
		e, ok := err.(*mysql.MySQLError)
		if ok {
			if e.Number == 1062 {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return err
	}
	if rows := result.RowsAffected; rows != 1 {
		return status.Errorf(codes.Internal, "%d rows affected", rows)
	}
	m.RoomID = rm.ID
	result = tx.Create(m)
	if err := result.Error; err != nil {
		tx.Rollback()
		return err
	}
	if rows := result.RowsAffected; rows != 1 {
		return status.Errorf(codes.Internal, "%d rows affected", rows)
	}
	if err := tx.Commit().Error; err != nil {
		return status.Error(codes.Aborted, err.Error())
	}
	return nil
}

func (r *roomRepository) AddMember(ctx context.Context, m *domain.Member) error {
	result := r.conn.Create(m)
	if err := result.Error; err != nil {
		e, ok := err.(*mysql.MySQLError)
		if ok {
			if e.Number == 1062 {
				err = status.Error(codes.AlreadyExists, err.Error())
			}
		}
		return err
	}
	if rows := result.RowsAffected; rows != 1 {
		return status.Errorf(codes.Internal, "%d rows affected", rows)
	}
	return nil
}

func (r *roomRepository) GetMemberByUserIDAndRoomID(ctx context.Context, uID int64, rID int64) (*domain.Member, error) {
	m := &domain.Member{}
	if err := r.conn.Where("user_id = ? AND room_id = ?", uID, rID).First(m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = status.Errorf(codes.NotFound, "user_id=%d is not in room_id=%d", uID, rID)
		}
		return nil, err
	}
	return m, nil
}

func (r *roomRepository) GetRoomByID(ctx context.Context, id int64) (*domain.Room, error) {
	rm := &domain.Room{}
	if err := r.conn.Where("id = ?", id).First(rm).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			err = status.Errorf(codes.NotFound, "room_id=%d is not found", id)
		}
		return nil, err
	}
	return rm, nil
}

func (r *roomRepository) CreateMessage(ctx context.Context, m *domain.Message) error {
	if err := r.conn.Create(m).Error; err != nil {
		return err
	}
	return nil
}
