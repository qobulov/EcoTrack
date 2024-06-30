package service

import (
	"context"
	"Impact-Calculator-Service/Storage/postgres"
	pb "Impact-Calculator-Service/genproto/impact-proto"
)

type CalculatorService struct {
	pb.UnimplementedImpactCalculatorServer
	CalculatorRepo *postgres.ImpactCalculator
}

func NewCalculatorService(CalculatorRepo *postgres.ImpactCalculator) *CalculatorService {
	return &CalculatorService{
		CalculatorRepo: CalculatorRepo,
	}
}

func (s *CalculatorService) CalculateCarbonFootprint(ctx context.Context, req *pb.CalculateCarbonFootprintRequest) (*pb.CalculateCarbonFootprintResponse, error) {
	return s.CalculatorRepo.CalculateCarbonFootprint(ctx, req)
}

func (s *CalculatorService) GetUserImpact(ctx context.Context, req *pb.GetUserImpactRequest) (*pb.GetUserImpactResponse, error) {
	return s.CalculatorRepo.GetUserImpact(ctx, req)
}

func (s *CalculatorService) GetGroupImpact(ctx context.Context, req *pb.GetGroupImpactRequest) (*pb.GetGroupImpactResponse, error) {
	return s.CalculatorRepo.GetGroupImpact(ctx, req)
}

func (s *CalculatorService) GetUserLeaderboard(ctx context.Context, req *pb.GetLeaderboardRequest) (*pb.GetLeaderboardResponse, error) {
	return s.CalculatorRepo.GetUserLeaderboard(ctx, req)
}

func (s *CalculatorService) GetGroupLeaderboard(ctx context.Context, req *pb.GetLeaderboardRequest) (*pb.GetLeaderboardResponse, error) {
	return s.CalculatorRepo.GetGroupLeaderboard(ctx, req)
}

func (s *CalculatorService) DonateToCause(ctx context.Context, req *pb.DonateToCauseRequest) (*pb.DonateToCauseResponse, error) {
	return s.CalculatorRepo.DonateToCause(ctx, req)
}

func (s *CalculatorService) GetDonationCauses(ctx context.Context, req *pb.GetDonationCausesRequest) (*pb.GetDonationCausesResponse, error) {
	return s.CalculatorRepo.GetDonationCauses(ctx, req)
}