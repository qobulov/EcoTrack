package main

import (
	"Habits-Tracker/pkg/db"
	pb "Habits-Tracker/genproto/protos"
	"Habits-Tracker/service"
	"Habits-Tracker/storage/postgres"
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

	habitTrackerStorage := postgres.NewHabitTracker(dbConn)

	ht := service.NewHabitsTracker(habitTrackerStorage)

	grpcServer := grpc.NewServer()

	pb.RegisterHabitTrackerServiceServer(grpcServer, ht)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
