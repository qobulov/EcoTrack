package main

import (
	p "EcoTrack/UserServis/genproto/protos"
	"EcoTrack/UserServis/pkg/db"
	service "EcoTrack/UserServis/service"
	storage "EcoTrack/UserServis/storage/postgres"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()
	listener,err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()
	log.Printf("Server started on port 50051...")

	userStorage := storage.NewUserRepo(db)
	us := service.NewUserService(userStorage)

	s := grpc.NewServer()
	p.RegisterUserServiceServer(s, us)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
