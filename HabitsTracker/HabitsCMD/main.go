package main

import (
	"Habits-Tracker/Storage/postgres"
	pb "Habits-Tracker/genproto/HabitTracker"
	service "Habits-Tracker/Service"
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
	ht := service.NewHabitsTracker(postgres.NewHabitTracker(db))

	grpcServer := grpc.NewServer()
	pb.RegisterHabitTrackerServiceServer(grpcServer,ht)
	grpcServer.Serve(listener)

	log.Printf("Server started on port 50051")
}
