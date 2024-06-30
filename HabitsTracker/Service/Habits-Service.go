package service

import (
	"context"
	"Habits-Tracker/Storage/postgres"
	pb "Habits-Tracker/genproto/HabitTracker"
)

type HabitsTracker struct {
	pb.UnimplementedHabitTrackerServiceServer
	db *postgres.HabitTracker
}

func NewHabitsTracker(db *postgres.HabitTracker) *HabitsTracker {
	return &HabitsTracker{db: db}
}

func (s *HabitsTracker) CreateHabit(ctx context.Context, req *pb.CreateHabitRequest) (*pb.CreateHabitResponse, error) {
	habit, err := s.db.CreateHabit(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CreateHabitResponse{Habit: habit}, nil
}

func (s *HabitsTracker) GetHabits(ctx context.Context, req *pb.GetHabitsRequest) (*pb.GetHabitsResponse, error) {
	habits, err := s.db.GetHabits(ctx, req)
	if err != nil {
		return nil, err
	}
	return habits, nil
}

func (s *HabitsTracker) LogHabit(ctx context.Context, req *pb.LogHabitRequest) (*pb.LogHabitResponse, error) {
	habitLog, err := s.db.LogHabit(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.LogHabitResponse{HabitLog: habitLog}, nil
}

func (s *HabitsTracker) GetHabitLogs(ctx context.Context, req *pb.GetHabitLogsRequest) (*pb.GetHabitLogsResponse, error) {
	habitLogs, err := s.db.GetHabitLogs(ctx, req)
	if err != nil {
		return nil, err
	}
	return habitLogs, nil
}