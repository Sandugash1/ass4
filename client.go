package main

import (
	"context"
	"log"

	pb "ass4/protofile"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create a new UserServiceClient
	client := pb.NewUserServiceClient(conn)

	// Test AddUser RPC
	user := &pb.User{
		Id:    1,
		Name:  "Sandu",
		Email: "sandu@example.com",
	}
	addUserResponse, err := client.AddUser(context.Background(), user)
	if err != nil {
		log.Fatalf("Failed to add user: %v", err)
	}
	log.Printf("User added: %v", addUserResponse)

	// Test GetUser RPC
	userID := int32(1)
	getUserResponse, err := client.GetUser(context.Background(), &pb.UserId{Id: userID})
	if err != nil {
		log.Fatalf("Failed to get user: %v", err)
	}
	log.Printf("User retrieved: %v", getUserResponse)

	// Test ListUsers RPC
	listUsersResponse, err := client.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}
	users := listUsersResponse.GetUsers()
	log.Printf("Users list:")
	for _, u := range users {
		log.Printf("- ID: %d, Name: %s, Email: %s", u.Id, u.Name, u.Email)
	}
}
