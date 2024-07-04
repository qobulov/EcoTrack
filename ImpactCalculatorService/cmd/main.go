package main

import (
	"Impact/Calculator/pkg/db"
	pb "Impact/Calculator/genproto/protos"
	"Impact/Calculator/service"
	"Impact/Calculator/storage/postgres"
	"log"
	"net"
	"google.golang.org/grpc"
)

func main() {
	dbConn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	log.Printf("Server started on port 50051...")

	impactCalculatorStorage := postgres.NewImpactCalculator(dbConn)

	ics := service.NewImpactCalculatorService(impactCalculatorStorage)

	grpcServer := grpc.NewServer()

	pb.RegisterImpactCalculatorServer(grpcServer, ics)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
