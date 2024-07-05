package main

import (
	"EcoTrack/community/config"
	p "EcoTrack/community/genproto/protos"
	"EcoTrack/community/pkg/db"
	s "EcoTrack/community/service"
	storage "EcoTrack/community/storage/postgres"

	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config := config.Load()

	db, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", config.URL_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	server := grpc.NewServer()

	storage := storage.NewCommunityRepo(db)
	service := s.CommunityService{Storage: *storage}

	p.RegisterCommunityServiceServer(server, &service)

	server.Serve(listener)

	log.Println("server is running on " + config.URL_PORT)
}
