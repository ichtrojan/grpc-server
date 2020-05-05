package chat

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("New Message: %s", message.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}
