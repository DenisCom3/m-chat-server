package converter

import (
	"github.com/DenisCom3/m-chat-server/internal/model"
	desc "github.com/DenisCom3/m-chat-server/pkg/user_v1"
)

func ToModel(u *desc.GetResponse) *model.User {

	return &model.User{
		ID: u.GetId(),
		Login: u.GetName(),
	}
}