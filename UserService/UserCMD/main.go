package main

import (
	"User-Service/Storage/postgres"
	pb "User-Service/genproto/user-proto"
	service "User-Service/Service"
	"net"
	"log"
	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	us := service.NewUserService(postgres.NewUser(db))	

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer,us)
	grpcServer.Serve(listener)

	log.Printf("Server started on port 50051")
}