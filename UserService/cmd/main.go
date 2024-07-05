package main

import (
	"EcoTrack/UserServis/config"
	p "EcoTrack/UserServis/genproto/protos"
	"EcoTrack/UserServis/pkg/db"
	service "EcoTrack/UserServis/service"
	storage "EcoTrack/UserServis/storage/postgres"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config := config.Load()
	db, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()
	listener, err := net.Listen("tcp", config.URL_PORT )
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()
	log.Printf("Server started on port "+config.URL_PORT)

	userStorage := storage.NewUserRepo(db)
	us := service.NewUserService(userStorage)

	s := grpc.NewServer()
	p.RegisterUserServiceServer(s, us)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
