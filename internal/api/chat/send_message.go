package chat

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/model"
	desc "github.com/DenisCom3/m-chat-server/pkg/chat_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {

	msg := model.CreateMessage{
		ChatId: req.GetToChatId(),
		UserId: req.GetFromUserId(),
		Text:   req.GetText(),
	}
	_,err := i.s.SendMessage(ctx, msg)

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}