package app

import (
	"context"
	"log"

	chatApi "github.com/DenisCom3/m-chat-server/internal/api/chat"
	"github.com/DenisCom3/m-chat-server/internal/client/auth"
	userv1 "github.com/DenisCom3/m-chat-server/internal/client/auth/user_v1"
	"github.com/DenisCom3/m-chat-server/internal/client/db"
	"github.com/DenisCom3/m-chat-server/internal/client/db/postgres"
	"github.com/DenisCom3/m-chat-server/internal/client/db/transaction"
	"github.com/DenisCom3/m-chat-server/internal/config"
	"github.com/DenisCom3/m-chat-server/internal/repository"
	chatRepo "github.com/DenisCom3/m-chat-server/internal/repository/chat"
	msgRepo "github.com/DenisCom3/m-chat-server/internal/repository/message"
	"github.com/DenisCom3/m-chat-server/internal/service"
	chatServ "github.com/DenisCom3/m-chat-server/internal/service/chat"
)

type serviceProvider struct {
	pgConfig config.Postgres
	grpcConfig config.Grpc
	authConfig config.Auth

	dbClient db.Client
	txManager db.TxManager

	authClient auth.Auth

	chatRepository repository.ChatRepository
	messageRepository repository.MessageRepository
	chatService service.ChatService

	chatImpl *chatApi.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.Grpc {
	if s.grpcConfig == nil {
		cfg := config.GetGrpc()
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) PostgresConfig() config.Postgres {
	if s.pgConfig == nil {
		cfg := config.GetPostgres()

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) AuthConfig() config.Auth {
	if s.authConfig == nil {
		cfg := config.GetAuth()
		s.authConfig = cfg
	}
	return s.authConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := postgres.New(ctx, s.PostgresConfig().Dsn())

		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = client.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.New(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) AuthClient(ctx context.Context) auth.Auth {
	if s.authClient == nil {
		client, err := userv1.New(ctx, s.AuthConfig().Address())
		if err != nil {
			log.Fatalf("failed to create auth client: %v", err)
		}

		s.authClient = client
	}

	return s.authClient
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = chatRepo.New((s.DBClient(ctx)))
	}
	return s.chatRepository
}

func (s *serviceProvider) MessageRepository(ctx context.Context) repository.MessageRepository {
	if s.messageRepository == nil {
		s.messageRepository = msgRepo.New((s.DBClient(ctx)))
		}
	return s.messageRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		d := chatServ.Dependencies{
			ChatRepo:     s.ChatRepository(ctx),
			MessageRepo:  s.MessageRepository(ctx),
			AuthClient:   s.AuthClient(ctx),	
		}
		s.chatService = chatServ.New(&d)
	}
	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chatApi.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatApi.New(s.ChatService(ctx))
	}
	return s.chatImpl
}