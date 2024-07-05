package postgres

import (
	"context"
	"testing"

	pb "Impact/Calculator/genproto/protos"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

type FootprintCategory int32

const (
    FootprintCategory_TRANSPORTATION FootprintCategory = 0
    FootprintCategory_ELECTRICITY     FootprintCategory = 1
    // ... other categories
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
        Category: 6, // Now it should work
        Amount:   50.0,
        Unit:     "kg",
    }
	

	mock.ExpectExec("INSERT INTO carbon_footprint_logs (user_id, category, amount, unit) VALUES ($1, $2, $3, $4)").
		WithArgs(req.UserId, req.Category, req.Amount, req.Unit).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := calculator.CalculateCarbonFootprint(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, req.Amount, resp.CarbonFootprint)
	assert.Equal(t, req.Unit, resp.Unit)
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

	mock.ExpectQuery("SELECT category, SUM(amount), unit FROM carbon_footprint_logs WHERE user_id = $1 GROUP BY category, unit").
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := calculator.GetUserImpact(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedCarbonFootprint := 80.0
	assert.Equal(t, expectedCarbonFootprint, resp.CarbonFootprint)
	assert.Equal(t, "kg", resp.Unit)
	assert.Equal(t, []string{"transportation: 50.00 kg", "electricity: 30.00 kg"}, resp.Actions)
}

func TestGetGroupImpact(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetGroupImpactRequest{GroupId: "1"}

	rows := sqlmock.NewRows([]string{"username", "sum", "unit"}).
		AddRow("user1", 50.0, "kg").
		AddRow("user2", 30.0, "kg")

	mock.ExpectQuery(`
	SELECT u.username, SUM(c.amount), c.unit 
	FROM carbon_footprint_logs c
	JOIN users u ON c.user_id = u.id
	JOIN group_members gm ON u.id = gm.user_id
	WHERE gm.group_id = $1
	GROUP BY u.username, c.unit`).
		WithArgs(req.GroupId).
		WillReturnRows(rows)

	resp, err := calculator.GetGroupImpact(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedCarbonFootprint := 80.0
	assert.Equal(t, expectedCarbonFootprint, resp.TotalCarbonFootprint)
	assert.Equal(t, "kg", resp.Unit)
	assert.Equal(t, map[string]float32{"user1": 50.0, "user2": 30.0}, resp.UserContributions)
}

func TestGetUserLeaderboard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetLeaderboardRequest{}

	rows := sqlmock.NewRows([]string{"id", "username", "total_amount", "unit"}).
		AddRow("1", "User1", 50.0, "kg").
		AddRow("2", "User2", 30.0, "kg")

	mock.ExpectQuery(`
	SELECT u.id, u.username, SUM(c.amount) AS total_amount, c.unit
	FROM carbon_footprint_logs c
	JOIN users u ON c.user_id = u.id
	GROUP BY u.id, u.username, c.unit
	ORDER BY total_amount DESC`).
		WithArgs().
		WillReturnRows(rows)

	resp, err := calculator.GetUserLeaderboard(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, 2, len(resp.Entries))
	assert.Equal(t, "User1", resp.Entries[0].Name)
	assert.Equal(t, float32(50.0), resp.Entries[0].CarbonFootprint)
	assert.Equal(t, "kg", resp.Entries[0].Unit)
}

func TestGetGroupLeaderboard(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetLeaderboardRequest{}

	rows := sqlmock.NewRows([]string{"id", "name", "total_amount", "unit"}).
		AddRow("1", "Group1", 100.0, "kg").
		AddRow("2", "Group2", 80.0, "kg")

	mock.ExpectQuery(`
	SELECT g.id, g.name, SUM(c.amount) AS total_amount, c.unit
	FROM carbon_footprint_logs c
	JOIN group_members gm ON c.user_id = gm.user_id
	JOIN groups g ON gm.group_id = g.id
	GROUP BY g.id, g.name, c.unit
	ORDER BY total_amount DESC`).
		WithArgs().
		WillReturnRows(rows)

	resp, err := calculator.GetGroupLeaderboard(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, 2, len(resp.Entries))
	assert.Equal(t, "Group1", resp.Entries[0].Name)
	assert.Equal(t, float32(100.0), resp.Entries[0].CarbonFootprint)
	assert.Equal(t, "kg", resp.Entries[0].Unit)
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

	mock.ExpectExec("INSERT INTO donations (user_id, cause, amount) VALUES ($1, $2, $3)").
		WithArgs(req.UserId, req.CauseId, req.Amount).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := calculator.DonateToCause(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.True(t, resp.Success)
	assert.Equal(t, "Donation logged successfully", resp.Message)
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

	assert.Equal(t, 2, len(resp.Causes))
	assert.Equal(t, "Cause1", resp.Causes[0].Name)
	assert.Equal(t, "Description1", resp.Causes[0].Description)
}

func TestGetUserDonations(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	calculator := NewImpactCalculator(db)
	req := &pb.GetUserDonationsRequest{UserId: "1"}

	rows := sqlmock.NewRows([]string{"cause", "user_name", "amount"}).
		AddRow("Cause1", "User1", 100.0).
		AddRow("Cause2", "User1", 50.0)

	mock.ExpectQuery("SELECT d.cause, u.name AS user_name, d.amount FROM donations d JOIN users u ON d.user_id = u.id WHERE d.user_id = $1").
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := calculator.GetUserDonations(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, 2, len(resp.Donations))
	assert.Equal(t, "Cause1", resp.Donations[0].Cause)
	assert.Equal(t, "User1", resp.Donations[0].UserId)
	assert.Equal(t, float32(100.0), resp.Donations[0].Amount)
}
