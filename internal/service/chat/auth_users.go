package chat

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/model"
)


func (s *serv) AuthUsers(ctx context.Context, userIDs []int64) ([]model.User, error) {

	users, err := s.authClient.GetUsers(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	var result []model.User

	for _, user := range users {
		result = append(result, model.User{
			ID: user.ID,
			Login: user.Login,
		})
	}

	return result, nil
}