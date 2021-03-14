package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"go-grpc-example/data"
	userpb "go-grpc-example/proto"

	"google.golang.org/grpc"
)

const portNum = "9000"

type userServer struct {
	userpb.UserServer
}

func (s *userServer) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	userID := req.UserId

	var userMessage *userpb.UserMessage
	for _, u := range data.UserData {
		if u.UserId != userID {
			continue
		}
		userMessage = u
		break
	}
	return &userpb.GetUserResponse{
		UserMessage: userMessage,
	}, nil
}

func (s *userServer) ListUser(ctx context.Context, req *userpb.ListUsersRequest) (*userpb.ListUsersResponse, error) {
	userMessage := make([]*userpb.UserMessage, len(data.UserData))
	for i, u := range data.UserData {
		userMessage[i] = u
	}
	fmt.Println(userMessage)

	return &userpb.ListUsersResponse{
		UserMessages: userMessage,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":"+portNum)
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServer(grpcServer, &userServer{})

	log.Printf("start gRPC Server on %s port", portNum)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to server: %s", err)
	}
}
