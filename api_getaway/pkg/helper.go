package pkg

import (
	pb "api-getaway/genproto"

	"google.golang.org/grpc"
)

func CreateNewHabbitTrackerClient() pb.HabitTrackerServiceClient {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewHabitTrackerServiceClient(conn)
}

func CreateNewImpactCalculatorClient() pb.ImpactCalculatorClient{
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return pb.NewImpactCalculatorClient(conn)
}
