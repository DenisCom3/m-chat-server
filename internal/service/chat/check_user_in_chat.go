package chat

import (
	"context"
)


func (s *serv) CheckUserInChat(ctx context.Context, chatID, userID int64) (bool, error) {


	isExists, err := s.chatRepo.CheckUserInChat(ctx, chatID, userID)

	if err != nil {
		return false, err
	}

	return isExists, nil
}