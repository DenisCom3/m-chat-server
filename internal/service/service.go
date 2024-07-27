package service

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/model"
)


type ChatService interface {
	CreateChat(ctx context.Context, chat model.CreateChat) (*model.Chat, error)
	CheckUserInChat(ctx context.Context, chatID, userID int64) (bool, error)
	SendMessage(ctx context.Context, message model.CreateMessage) (*model.Message, error)
	AuthUsers(ctx context.Context, chatIDs []int64) ([]model.User, error)
}