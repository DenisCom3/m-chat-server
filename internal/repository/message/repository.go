package message

import (
	"github.com/DenisCom3/m-chat-server/internal/client/db"
	"github.com/DenisCom3/m-chat-server/internal/repository"
)

type repo struct {
	db db.Client
}

func New(db db.Client) repository.MessageRepository {
	return &repo{
		db: db,
	}
}