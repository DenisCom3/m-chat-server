package userv1

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/client/auth"
	"github.com/DenisCom3/m-chat-server/internal/client/auth/user_v1/converter"
	"github.com/DenisCom3/m-chat-server/internal/model"
	desc "github.com/DenisCom3/m-chat-server/pkg/user_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var _ auth.Auth = (*UserV1Auth)(nil)

type UserV1Auth struct{
	client desc.UserV1Client
}


func New(ctx context.Context, address string) (auth.Auth, error) {

	conn, err := grpc.NewClient("127.0.0.1:4200", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}
	defer conn.Close()


	client := desc.NewUserV1Client(conn)

	return &UserV1Auth{
		client: client,
	}, nil
}

func (a *UserV1Auth) GetUser(ctx context.Context, id int64) (*model.User, error) {
	
	user, err := a.client.Get(ctx, &desc.GetRequest{Id: id})

	if err != nil {
		return nil, err
	}

	return converter.ToModel(user), nil

}

func (a *UserV1Auth) GetUsers(ctx context.Context, ids []int64) ([]*model.User, error) {

	users := make([]*model.User, 2)

	for _, id := range ids {
		user, err := a.GetUser(ctx, id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}