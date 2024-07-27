package repository

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, chat model.CreateChat) (*model.Chat, error)
	CheckUserInChat(ctx context.Context, chatID, userID int64) (bool, error)

}

type MessageRepository interface {
	Create(ctx context.Context, message model.CreateMessage) (*model.Message, error)
}
