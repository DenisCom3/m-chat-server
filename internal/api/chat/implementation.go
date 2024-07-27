package chat

import (
	"github.com/DenisCom3/m-chat-server/internal/service"
	desc "github.com/DenisCom3/m-chat-server/pkg/chat_v1"
)

type Implementation struct {
	s service.ChatService
	desc.UnimplementedChatV1Server 
}

func New( s service.ChatService) *Implementation {
	return &Implementation{
		s: s,
	}
}