package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/osvaldosilitonga/hotel-and-resto/user-service/internal/pb_user_service"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/services"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Can't load .env file.\n[ERR]: %v", err)
	}
}

func main() {
	PORT := os.Getenv("PORT")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatalf("Server can't listen. \n[ERR]: %v", err)
	}

	userService := services.NewUserService()

	grpcServer := grpc.NewServer()
	pb.RegisterUserServer(grpcServer, userService)

	log.Printf("Starting User Service gRPC listener on port : %v\n", PORT)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to running User gRPC Server. \n[ERR]: %v", err)
	}
}