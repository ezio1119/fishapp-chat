package registry

import (
	"github.com/ezio1119/fishapp-chat/interfaces/controllers"
	"github.com/ezio1119/fishapp-chat/interfaces/controllers/chat_grpc"
	"github.com/ezio1119/fishapp-chat/interfaces/presenter"
	"github.com/ezio1119/fishapp-chat/interfaces/repository"
	"github.com/ezio1119/fishapp-chat/usecase/interactor"
)

func (r *registry) NewChatController() chat_grpc.ChatServiceServer {
	return controllers.NewChatController(
		interactor.NewChatInteractor(
			repository.NewRoomRepository(r.db),
			repository.NewQueuingRepository(r.kvs),
			presenter.NewChatPresenter(),
			r.ctxTimeout,
		),
	)
}
