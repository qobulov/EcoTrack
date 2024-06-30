package main

import (
	"Impact-Calculator-Service/Storage/postgres"
	pb "Impact-Calculator-Service/genproto/impact-proto"
	service "Impact-Calculator-Service/Sevice"
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
	ic := service.NewCalculatorService(postgres.NewImpactCalculator(db))
	
	grpcServer := grpc.NewServer()
	pb.RegisterImpactCalculatorServer(grpcServer,ic)
	grpcServer.Serve(listener)

	log.Printf("Server started on port 50051")
}
