package auth

import (
	"context"
	"github.com/DenisCom3/m-chat-server/internal/model"
	
)

type Auth interface {
	GetUser(ctx context.Context, id int64) (*model.User, error)
	GetUsers(ctx context.Context, ids []int64) ([]*model.User, error)
}


