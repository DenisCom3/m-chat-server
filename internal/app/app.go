package app

import (
	"context"
	"log"
	"net"

	"github.com/DenisCom3/m-chat-server/internal/closer"
	"github.com/DenisCom3/m-chat-server/internal/config"
	desc "github.com/DenisCom3/m-chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)


type App struct {

	serviceProvider *serviceProvider

	grpcServer *grpc.Server
}

func (a *App) Run() error {
	defer func () {
		closer.CloseAll()
	    closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	list, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return err
	}

	return nil
}


func New(ctx context.Context) (*App, error) {

	a := &App{}

	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}
	return a, nil
	
}

func (a *App) initDeps(ctx context.Context) error {

	 deps := [...]func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, dep := range deps {
		if err := dep(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.MustLoad()

	if err != nil {
		return err
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	desc.RegisterChatV1Server(a.grpcServer, a.serviceProvider.ChatImpl(ctx))

	return nil
}