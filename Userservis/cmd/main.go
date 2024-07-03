package main

import (
	"EcoTrack/UserServis/pkg/db"
	storage "EcoTrack/UserServis/storage/postgres"
	s "EcoTrack/UserServis/service"
	p "EcoTrack/UserServis/genproto/protos"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server := grpc.NewServer()

	db, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	storage := storage.NewUserRepo(db)
	service := s.UserService{Storage: *storage}

	p.RegisterUserServiceServer(server, &service)

	log.Println("server is running on :7070 ...")
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
