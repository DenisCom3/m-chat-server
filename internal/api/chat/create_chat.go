package chat

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/model"
	desc "github.com/DenisCom3/m-chat-server/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error)  {
	
	users, err := i.s.AuthUsers(ctx, r.GetUsersId())

	if err != nil {
		return nil, err
	}
	
	chat := model.CreateChat{
		Name: r.GetChatName(),
		Users: users,
	}

	chatCreated, err := i.s.CreateChat(ctx, chat)
	if err != nil {
		return nil, err
	}	
	return &desc.CreateResponse{Id: chatCreated.ID}, nil
}