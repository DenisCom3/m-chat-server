package chat

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/model"
)

func (s *serv) SendMessage(ctx context.Context, message model.CreateMessage) (*model.Message, error) {

	msg, err := s.messageRepo.Create(ctx, message)

	if err != nil {
		return nil, err
	}

	return msg, nil

}