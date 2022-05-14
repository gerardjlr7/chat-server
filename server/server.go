package server

import (
	"context"
	"errors"
	"io"
	"reflect"
	"sync"

	"github.com/gerardjlr7/chat-server/protos"
)

type server struct {
	protos.UnimplementedChatServer
	connectedUsers map[string]bool //By username bool defines who is writing
	mutex          *sync.RWMutex
	userMessages   map[string]*QueueMessages
}

func NewChatServer() *server {
	return &server{
		connectedUsers: map[string]bool{},
		mutex:          &sync.RWMutex{},
		userMessages:   map[string]*QueueMessages{},
	}
}

func (s *server) Connect(ctx context.Context, input *protos.ConnectRequest) (*protos.ConnectReply, error) {
	s.mutex.Lock()

	if _, ok := s.connectedUsers[input.Name]; ok {
		s.mutex.Unlock()
		return nil, errors.New("User " + input.Name + " already connected")
	}

	s.connectedUsers[input.Name] = false
	if s.userMessages[input.Name] == nil {
		s.userMessages[input.Name] = NewQueue()
	}
	s.mutex.Unlock()

	return &protos.ConnectReply{
		Message: "User " + input.Name + " connected successfully",
	}, nil
}

func (s *server) Disconnect(ctx context.Context, input *protos.ConnectRequest) (*protos.ConnectReply, error) {
	s.mutex.Lock()

	if _, ok := s.connectedUsers[input.Name]; !ok {
		s.mutex.Unlock()
		return nil, errors.New("User " + input.Name + " isn't connected")
	}

	delete(s.connectedUsers, input.Name)
	s.mutex.Unlock()

	return &protos.ConnectReply{
		Message: "User " + input.Name + " disconnected successfully",
	}, nil
}

func (s *server) Writing(stream protos.Chat_WritingServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			reply := s.getConnectedUsers()
			return stream.SendAndClose(reply)
		}
		if err != nil {
			return err
		}
		s.mutex.Lock()
		s.connectedUsers[in.Name] = in.Writing
		s.mutex.Unlock()
	}
}

func (s *server) WhoIsWriting(in *protos.StatusRequest, stream protos.Chat_WhoIsWritingServer) error {
	previousStatus := map[string]bool{}
	for {
		s.mutex.Lock()
		if !reflect.DeepEqual(previousStatus, s.connectedUsers) {
			reply := s.getConnectedUsers()
			if err := stream.Send(reply); err != nil {
				s.mutex.Unlock()
				return err
			}
			s.refreshMap(previousStatus)
		}
		s.mutex.Unlock()
	}

}

func (s *server) SendMessage(stream protos.Chat_SendMessageServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			reply := &protos.MessageReply{}
			return stream.SendAndClose(reply)
		}
		if err != nil {
			return err
		}
		s.mutex.Lock()
		if _, ok := s.connectedUsers[in.Name]; !ok {
			s.mutex.Unlock()
			return errors.New("User " + in.Name + " isn't connected")
		}
		for username, queue := range s.userMessages {
			if username != in.Name {
				queue.Enqueue(in)
			}
		}
		s.mutex.Unlock()
	}
}

func (s *server) ReceiveMessage(in *protos.MessageRequest, stream protos.Chat_ReceiveMessageServer) error {
	for {
		s.mutex.Lock()
		if s.userMessages[in.Name] != nil {
			if msg := s.userMessages[in.Name].Dequeue(); msg != nil {
				if err := stream.Send(msg); err != nil {
					s.mutex.Unlock()
					return err
				}
			}
		}
		s.mutex.Unlock()
	}

}

func (s *server) refreshMap(previousStatus map[string]bool) {
	for key, _ := range previousStatus {
		delete(previousStatus, key)
	}
	for key, value := range s.connectedUsers {
		previousStatus[key] = value
	}
}

func (s *server) getConnectedUsers() *protos.StatusReply {
	usersWriting := []*protos.StatusUserReply{}
	for userName, isWriting := range s.connectedUsers {
		usersWriting = append(usersWriting, &protos.StatusUserReply{
			Name:    userName,
			Writing: isWriting,
		})
	}
	reply := &protos.StatusReply{
		UserStatus: usersWriting,
	}
	return reply
}
