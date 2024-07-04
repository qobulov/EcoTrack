package handler

import (
	pb "api-getaway/genproto/protos"
)

type Handler struct {
	Habits pb.HabitTrackerServiceClient
	Impact pb.ImpactCalculatorServiceClient
}

func NewHandler(H pb.HabitTrackerServiceClient, I pb.ImpactCalculatorServiceClient) *Handler {
	return &Handler{
		Habits: H,
		Impact: I,
	}
}
