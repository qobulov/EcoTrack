package handler

import (
	pb "api-getaway/genproto"
)

type Handler struct {
	Habits pb.HabitTrackerServiceClient
	Impact pb.ImpactCalculatorClient
}

func NewHandler(H pb.HabitTrackerServiceClient, I pb.ImpactCalculatorClient) *Handler {
	return &Handler{
		Habits: H,
		Impact: I,
	}
}
