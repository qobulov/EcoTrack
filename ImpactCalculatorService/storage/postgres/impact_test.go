package postgres

import (
	"context"
	"testing"

	pb "Impact/Calculator/genproto/protos"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCalculateCarbonFootprint(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.CalculateCarbonFootprintRequest{
		UserId:   "1",
		Category: "transportation",
		Amount:   50.0,
		Unit:     "kg",
	}

	mock.ExpectExec("INSERT INTO carbon_footprint_logs \\(user_id, category, amount, unit\\) VALUES \\(\\$1, \\$2, \\$3, \\$4\\)").
		WithArgs(req.UserId, req.Category, req.Amount, req.Unit).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := calculator.CalculateCarbonFootprint(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.CarbonFootprint != req.Amount {
		t.Errorf("expected carbon footprint %v, got %v", req.Amount, resp.CarbonFootprint)
	}
	if resp.Unit != req.Unit {
		t.Errorf("expected unit %v, got %v", req.Unit, resp.Unit)
	}
}

func TestGetUserImpact(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetUserImpactRequest{UserId: "1"}

	rows := sqlmock.NewRows([]string{"category", "sum", "unit"}).
		AddRow("transportation", 50.0, "kg").
		AddRow("electricity", 30.0, "kg")

	mock.ExpectQuery("SELECT category, SUM\\(amount\\), unit FROM carbon_footprint_logs WHERE user_id = \\$1 GROUP BY category, unit").
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := calculator.GetUserImpact(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedCarbonFootprint := 80.0
	if resp.CarbonFootprint != expectedCarbonFootprint {
		t.Errorf("expected carbon footprint %v, got %v", expectedCarbonFootprint, resp.CarbonFootprint)
	}
	if resp.Unit != "kg" {
		t.Errorf("expected unit 'kg', got %v", resp.Unit)
	}
}

func TestGetGroupImpact(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetGroupImpactRequest{GroupId: "1"}

	rows := sqlmock.NewRows([]string{"user_id", "sum", "unit"}).
		AddRow("1", 50.0, "kg").
		AddRow("2", 30.0, "kg")

	mock.ExpectQuery(`
	SELECT u.id, SUM\\(c.amount\\), c.unit 
	FROM carbon_footprint_logs c
	JOIN users u ON c.user_id = u.id
	JOIN group_members gm ON u.id = gm.user_id
	WHERE gm.group_id = \\$1
	GROUP BY u.id, c.unit`).
		WithArgs(req.GroupId).
		WillReturnRows(rows)

	resp, err := calculator.GetGroupImpact(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedCarbonFootprint := 80.0
	if resp.TotalCarbonFootprint != expectedCarbonFootprint {
		t.Errorf("expected total carbon footprint %v, got %v", expectedCarbonFootprint, resp.TotalCarbonFootprint)
	}
	if resp.Unit != "kg" {
		t.Errorf("expected unit 'kg', got %v", resp.Unit)
	}
}

func TestGetUserLeaderboard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetLeaderboardRequest{Limit: 5}

	rows := sqlmock.NewRows([]string{"id", "name", "total_amount", "unit"}).
		AddRow("1", "User1", 50.0, "kg").
		AddRow("2", "User2", 30.0, "kg")

	mock.ExpectQuery(`
	SELECT u.id, u.name, SUM\\(c.amount\\) AS total_amount, c.unit
	FROM carbon_footprint_logs c
	JOIN users u ON c.user_id = u.id
	GROUP BY u.id, u.name, c.unit
	ORDER BY total_amount DESC
	LIMIT \\$1`).
		WithArgs(req.Limit).
		WillReturnRows(rows)

	resp, err := calculator.GetUserLeaderboard(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(resp.Entries) != 2 {
		t.Errorf("expected 2 entries, got %v", len(resp.Entries))
	}
	if resp.Entries[0].Name != "User1" {
		t.Errorf("expected first entry name 'User1', got %v", resp.Entries[0].Name)
	}
	if resp.Entries[0].CarbonFootprint != 50.0 {
		t.Errorf("expected first entry carbon footprint 50.0, got %v", resp.Entries[0].CarbonFootprint)
	}
}

func TestGetGroupLeaderboard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetLeaderboardRequest{Limit: 5}

	rows := sqlmock.NewRows([]string{"id", "name", "total_amount", "unit"}).
		AddRow("1", "Group1", 100.0, "kg").
		AddRow("2", "Group2", 80.0, "kg")

	mock.ExpectQuery(`
	SELECT g.id, g.name, SUM\\(c.amount\\) AS total_amount, c.unit
	FROM carbon_footprint_logs c
	JOIN group_members gm ON c.user_id = gm.user_id
	JOIN groups g ON gm.group_id = g.id
	GROUP BY g.id, g.name, c.unit
	ORDER BY total_amount DESC
	LIMIT \\$1`).
		WithArgs(req.Limit).
		WillReturnRows(rows)

	resp, err := calculator.GetGroupLeaderboard(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(resp.Entries) != 2 {
		t.Errorf("expected 2 entries, got %v", len(resp.Entries))
	}
	if resp.Entries[0].Name != "Group1" {
		t.Errorf("expected first entry name 'Group1', got %v", resp.Entries[0].Name)
	}
	if resp.Entries[0].CarbonFootprint != 100.0 {
		t.Errorf("expected first entry carbon footprint 100.0, got %v", resp.Entries[0].CarbonFootprint)
	}
}

func TestDonateToCause(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.DonateToCauseRequest{
		UserId:  "1",
		CauseId: "1",
		Amount:  100.0,
	}

	mock.ExpectExec("INSERT INTO donations \\(user_id, cause, amount\\) VALUES \\(\\$1, \\$2, \\$3\\)").
		WithArgs(req.UserId, req.CauseId, req.Amount).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := calculator.DonateToCause(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !resp.Success {
		t.Errorf("expected success true, got %v", resp.Success)
	}
	if resp.Message != "Donation logged successfully" {
		t.Errorf("expected message 'Donation logged successfully', got %v", resp.Message)
	}
}

func TestGetDonationCauses(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetDonationCausesRequest{}

	rows := sqlmock.NewRows([]string{"id", "name", "description"}).
		AddRow(1, "Cause1", "Description1").
		AddRow(2, "Cause2", "Description2")

	mock.ExpectQuery("SELECT id, name, description FROM donation_causes").
		WillReturnRows(rows)

	resp, err := calculator.GetDonationCauses(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(resp.Causes) != 2 {
		t.Errorf("expected 2 causes, got %v", len(resp.Causes))
	}
	if resp.Causes[0].Name != "Cause1" {
		t.Errorf("expected first cause name 'Cause1', got %v", resp.Causes[0].Name)
	}
	if resp.Causes[0].Description != "Description1" {
		t.Errorf("expected first cause description 'Description1', got %v", resp.Causes[0].Description)
	}
}
