package service

import (
	"context"
	"Impact/Calculator/storage/postgres"
	pb "Impact/Calculator/genproto/protos"
)

type ImpactCalculatorService struct {
	pb.UnimplementedImpactCalculatorServer
	db *postgres.ImpactCalculator
}

func NewImpactCalculatorService(db *postgres.ImpactCalculator) *ImpactCalculatorService {
	return &ImpactCalculatorService{db: db}
}

func (s *ImpactCalculatorService) CalculateCarbonFootprint(ctx context.Context, req *pb.CalculateCarbonFootprintRequest) (*pb.CalculateCarbonFootprintResponse, error) {
	carbonFootprint, err := s.db.CalculateCarbonFootprint(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.CalculateCarbonFootprintResponse{CarbonFootprint: carbonFootprint.CarbonFootprint, Unit: carbonFootprint.Unit}, nil
}

func (s *ImpactCalculatorService) GetUserImpact(ctx context.Context, req *pb.GetUserImpactRequest) (*pb.GetUserImpactResponse, error) {
	userImpact, err := s.db.GetUserImpact(ctx, req)
	if err != nil {
		return nil, err
	}
	return userImpact, nil
}

func (s *ImpactCalculatorService) GetGroupImpact(ctx context.Context, req *pb.GetGroupImpactRequest) (*pb.GetGroupImpactResponse, error) {
	groupImpact, err := s.db.GetGroupImpact(ctx, req)
	if err != nil {
		return nil, err
	}
	return groupImpact, nil
}

func (s *ImpactCalculatorService) GetUserLeaderboard(ctx context.Context, req *pb.GetLeaderboardRequest) (*pb.GetLeaderboardResponse, error) {
	leaderboard, err := s.db.GetUserLeaderboard(ctx, req)
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}

func (s *ImpactCalculatorService) GetGroupLeaderboard(ctx context.Context, req *pb.GetLeaderboardRequest) (*pb.GetLeaderboardResponse, error) {
	leaderboard, err := s.db.GetGroupLeaderboard(ctx, req)
	if err != nil {
		return nil, err
	}
	return leaderboard, nil
}

func (s *ImpactCalculatorService) DonateToCause(ctx context.Context, req *pb.DonateToCauseRequest) (*pb.DonateToCauseResponse, error) {
	donation, err := s.db.DonateToCause(ctx, req)
	if err != nil {
		return nil, err
	}
	return &pb.DonateToCauseResponse{Success: donation.Success, Message: donation.Message}, nil
}

func (s *ImpactCalculatorService) GetDonationCauses(ctx context.Context, req *pb.GetDonationCausesRequest) (*pb.GetDonationCausesResponse, error) {
	causes, err := s.db.GetDonationCauses(ctx, req)
	if err != nil {
		return nil, err
	}
	return causes, nil
}
