package handler

import (
	pb "api-getaway/genproto/protos"
)

type Handler struct {
	Habits pb.HabitTrackerServiceClient
	Impact pb.ImpactCalculatorClient
    Community pb.CommunityServiceClient
}

func NewHandler(H pb.HabitTrackerServiceClient, I pb.ImpactCalculatorClient, C pb.CommunityServiceClient) *Handler {
	return &Handler{
		Habits: H,
		Impact: I,
		Community: C,
	}
}
