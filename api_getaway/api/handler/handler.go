package handler

import (
	pb "api-getaway/genproto/protos"
)

type Handler struct {
	Habits pb.HabitTrackerServiceClient
	Impact pb.ImpactCalculatorServiceClient
	Community pb.CommunityServiceClient
	User pb.UserServiceClient
	AuthService pb.AuthServiceClient
}

func NewHandler(H pb.HabitTrackerServiceClient, I pb.ImpactCalculatorServiceClient, C pb.CommunityServiceClient, U pb.UserServiceClient, A pb.AuthServiceClient) *Handler {
	return &Handler{
		Habits: H,
		Impact: I,
		Community: C,
		User: U,
		AuthService: A,
	}
}