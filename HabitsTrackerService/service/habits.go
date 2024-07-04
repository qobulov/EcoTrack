package service

import (
	"context"
	pb "Habits-Tracker/genproto/protos"
	"Habits-Tracker/storage/postgres"
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

func (s *HabitsTracker) GetUserHabits(ctx context.Context, req *pb.GetUserHabitsRequest) (*pb.GetUserHabitsResponse, error) {
	habits, err := s.db.GetUserHabits(ctx, req)
	if err != nil {
		return nil, err
	}
	return habits, nil
}

func (s *HabitsTracker) UpdateHabit(ctx context.Context, req *pb.UpdateHabitRequest) (*pb.UpdateHabitResponse, error) {
	habit, err := s.db.UpdateHabit(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateHabitResponse{Habit: habit}, nil
}

func (s *HabitsTracker) GetHabitSuggestions(ctx context.Context, req *pb.GetHabitSuggestionsRequest) (*pb.GetHabitSuggestionsResponse, error) {
	habits, err := s.db.GetHabitSuggestions(ctx, req)
	if err != nil {
		return nil, err
	}
	return habits, nil
}

func (s *HabitsTracker) DeleteHabit(ctx context.Context, req *pb.DeleteHabitRequest) (*pb.DeleteHabitResponse, error) {
	mes,err := s.db.DeleteHabit(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteHabitResponse{Message: mes.Message}, nil
}
