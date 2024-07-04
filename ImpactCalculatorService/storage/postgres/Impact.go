package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	pb "Impact/Calculator/genproto/protos"
)

type ImpactCalculator struct {
	db *sql.DB
}

func NewImpactCalculator(db *sql.DB) *ImpactCalculator {
	return &ImpactCalculator{db: db}
}
// CalculateCarbonFootprint foydalanuvchi uchun Carbon Footprint qayd qiladi
func (i *ImpactCalculator) CalculateCarbonFootprint(ctx context.Context, request *pb.CalculateCarbonFootprintRequest) (*pb.CalculateCarbonFootprintResponse, error) {
	query := `INSERT INTO carbon_footprint_logs (user_id, category, amount, unit) VALUES ($1, $2, $3, $4) RETURNING id`
	_, err := i.db.ExecContext(ctx, query, request.UserId, request.Category, request.Amount, request.Unit)
	if err != nil {
		log.Printf("Error inserting carbon footprint log: %v", err)
		return nil, err
	}

	return &pb.CalculateCarbonFootprintResponse{
		CarbonFootprint: request.Amount,
		Unit:            request.Unit,
	}, nil
}

// GetUserImpact foydalanuvchi uchun umumiy uglerod izini oladi
func (i *ImpactCalculator) GetUserImpact(ctx context.Context, request *pb.GetUserImpactRequest) (*pb.GetUserImpactResponse, error) {
	query := `SELECT category, SUM(amount), unit FROM carbon_footprint_logs WHERE user_id = $1 GROUP BY category, unit`
	rows, err := i.db.QueryContext(ctx, query, request.UserId)
	if err != nil {
		log.Printf("Error querying user impact: %v", err)
		return nil, err
	}
	defer rows.Close()

	var totalCarbonFootprint float64
	var unit string
	var actions []string
	for rows.Next() {
		var category string
		var amount float64
		if err := rows.Scan(&category, &amount, &unit); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		totalCarbonFootprint += amount
		actions = append(actions, fmt.Sprintf("%s: %.2f %s", category, amount, unit))
	}

	return &pb.GetUserImpactResponse{
		CarbonFootprint: totalCarbonFootprint,
		Unit:            unit,
		Actions:         actions,
	}, nil
}

// GetGroupImpact guruh uchun umumiy uglerod izini oladi
func (i *ImpactCalculator) GetGroupImpact(ctx context.Context, request *pb.GetGroupImpactRequest) (*pb.GetGroupImpactResponse, error) {
	query := `
	SELECT u.id, SUM(c.amount), c.unit 
	FROM carbon_footprint_logs c
	JOIN users u ON c.user_id = u.id
	JOIN group_members gm ON u.id = gm.user_id
	WHERE gm.group_id = $1
	GROUP BY u.id, c.unit`
	rows, err := i.db.QueryContext(ctx, query, request.GroupId)
	if err != nil {
		log.Printf("Error querying group impact: %v", err)
		return nil, err
	}
	defer rows.Close()

	var totalCarbonFootprint float64
	var unit string
	userContributions := make(map[string]float64)
	for rows.Next() {
		var userID string
		var amount float64
		if err := rows.Scan(&userID, &amount, &unit); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		totalCarbonFootprint += amount
		userContributions[userID] = amount
	}

	return &pb.GetGroupImpactResponse{
		TotalCarbonFootprint: totalCarbonFootprint,
		Unit:                 unit,
		UserContributions:    userContributions,
	}, nil
}

// GetUserLeaderboard foydalanuvchilar uchun uglerod izlari asosida peshqadamlar jadvalini oladi
func (i *ImpactCalculator) GetUserLeaderboard(ctx context.Context, request *pb.GetLeaderboardRequest) (*pb.GetLeaderboardResponse, error) {
	query := `
	SELECT u.id, u.name, SUM(c.amount) AS total_amount, c.unit
	FROM carbon_footprint_logs c
	JOIN users u ON c.user_id = u.id
	GROUP BY u.id, u.name, c.unit
	ORDER BY total_amount DESC
	LIMIT $1`
	rows, err := i.db.QueryContext(ctx, query, request.Limit)
	if err != nil {
		log.Printf("Error querying user leaderboard: %v", err)
		return nil, err
	}
	defer rows.Close()

	var entries []*pb.LeaderboardEntry
	for rows.Next() {
		var entry pb.LeaderboardEntry
		if err := rows.Scan(&entry.Id, &entry.Name, &entry.CarbonFootprint, &entry.Unit); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		entries = append(entries, &entry)
	}

	return &pb.GetLeaderboardResponse{
		Entries: entries,
	}, nil
}

// GetGroupLeaderboard uglerod izlari asosida guruhlar uchun peshqadamlar jadvalini oladi
func (i *ImpactCalculator) GetGroupLeaderboard(ctx context.Context, request *pb.GetLeaderboardRequest) (*pb.GetLeaderboardResponse, error) {
	query := `
	SELECT g.id, g.name, SUM(c.amount) AS total_amount, c.unit
	FROM carbon_footprint_logs c
	JOIN group_members gm ON c.user_id = gm.user_id
	JOIN groups g ON gm.group_id = g.id
	GROUP BY g.id, g.name, c.unit
	ORDER BY total_amount DESC
	LIMIT $1`
	rows, err := i.db.QueryContext(ctx, query, request.Limit)
	if err != nil {
		log.Printf("Error querying group leaderboard: %v", err)
		return nil, err
	}
	defer rows.Close()

	var entries []*pb.LeaderboardEntry
	for rows.Next() {
		var entry pb.LeaderboardEntry
		if err := rows.Scan(&entry.Id, &entry.Name, &entry.CarbonFootprint, &entry.Unit); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		entries = append(entries, &entry)
	}

	return &pb.GetLeaderboardResponse{
		Entries: entries,
	}, nil
}

// DonateToCause foydalanuvchi tomonidan qilingan xayriyani qayd qiladi
func (i *ImpactCalculator) DonateToCause(ctx context.Context, request *pb.DonateToCauseRequest) (*pb.DonateToCauseResponse, error) {
	query := `INSERT INTO donations (user_id, cause, amount) VALUES ($1, $2, $3)`
	_, err := i.db.ExecContext(ctx, query, request.UserId, request.CauseId, request.Amount)
	if err != nil {
		log.Printf("Error inserting donation: %v", err)
		return nil, err
	}

	return &pb.DonateToCauseResponse{
		Success: true,
		Message: "Donation logged successfully",
	}, nil
}

// GetDonationCauses xayriya sabablari ro'yxatini oladi
func (i *ImpactCalculator) GetDonationCauses(ctx context.Context, request *pb.GetDonationCausesRequest) (*pb.GetDonationCausesResponse, error) {
	query := `SELECT id, name, description FROM donation_causes`
	rows, err := i.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error querying donation causes: %v", err)
		return nil, err
	}
	defer rows.Close()

	var causes []*pb.DonationCause
	for rows.Next() {
		var cause pb.DonationCause
		if err := rows.Scan(&cause.Id, &cause.Name, &cause.Description); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}
		causes = append(causes, &cause)
	}

	return &pb.GetDonationCausesResponse{
		Causes: causes,
	}, nil
}
