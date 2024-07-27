package converter

import (
	model "github.com/DenisCom3/m-chat-server/internal/model"
	repo "github.com/DenisCom3/m-chat-server/internal/repository/chat/model"
)

func ToModel(u *repo.Chat) *model.Chat {
	return &model.Chat{
		ID:    u.ID,
		Name:  u.Name,
		UsersId: u.Users,
		IsActive: u.IsActive,
	}
}