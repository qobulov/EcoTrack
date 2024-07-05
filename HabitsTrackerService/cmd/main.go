package main

import (
	"Habits-Tracker/pkg/db"
	"Habits-Tracker/config"
	pb "Habits-Tracker/genproto/protos"
	"Habits-Tracker/service"
	"Habits-Tracker/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	configr := config.Load()

	dbConn, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	listener, err := net.Listen("tcp", configr.URL_PORT)
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
