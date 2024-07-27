package chat

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/model"
)

	func (s *serv) CreateChat(ctx context.Context, chat model.CreateChat) (*model.Chat, error) {

		chatCreated, err := s.chatRepo.Create(ctx, chat)

		if err != nil {
			return nil, err
		}

		return chatCreated, nil

	}