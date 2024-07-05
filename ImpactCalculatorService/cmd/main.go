package main

import (
	"Impact/Calculator/config"
	pb "Impact/Calculator/genproto/protos"
	"Impact/Calculator/pkg/db"
	"Impact/Calculator/service"
	"Impact/Calculator/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config := config.Load()

	dbConn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	listener, err := net.Listen("tcp", config.URL_PORT)
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	log.Printf("Server started on port "+config.URL_PORT)

	impactCalculatorStorage := postgres.NewImpactCalculator(dbConn)

	ics := service.NewImpactCalculatorService(impactCalculatorStorage)

	grpcServer := grpc.NewServer()

	pb.RegisterImpactCalculatorServiceServer(grpcServer, ics)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
