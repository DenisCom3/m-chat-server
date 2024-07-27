package chat

import (
	"github.com/DenisCom3/m-chat-server/internal/client/auth"
	"github.com/DenisCom3/m-chat-server/internal/repository"
	"github.com/DenisCom3/m-chat-server/internal/service"
)

var _ service.ChatService = (*serv)(nil)

type serv struct {
	chatRepo    repository.ChatRepository
	messageRepo repository.MessageRepository

	authClient  auth.Auth
}

type Dependencies struct {
	ChatRepo    repository.ChatRepository
	MessageRepo repository.MessageRepository
	AuthClient  auth.Auth
}

func New(d *Dependencies) service.ChatService {
	return &serv{
		chatRepo:    d.ChatRepo,
		messageRepo: d.MessageRepo,

		authClient:  d.AuthClient,
	}
}