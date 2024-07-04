package postgres

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	pb "Habits-Tracker/genproto/protos"
)

func TestCreateHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	now := time.Now()
	req := &pb.CreateHabitRequest{
		UserId:      "user1",
		Name:        "Exercise",
		Description: "Daily exercise routine",
		Frequency:   pb.Frequency_DAILY,
	}

	// Expected query and arguments matcher
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO habits (id, user_id, name, description, frequency) VALUES ($1, $2, $3, $4, $5)`)).
		WithArgs(sqlmock.AnyArg(), req.UserId, req.Name, req.Description, req.Frequency.String()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	habit, err := habitTracker.CreateHabit(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, habit)
	assert.NotEmpty(t, habit.Id)
	assert.Equal(t, req.UserId, habit.UserId)
	assert.Equal(t, req.Name, habit.Name)
	assert.Equal(t, req.Description, habit.Description)
	assert.Equal(t, req.Frequency, habit.Frequency)
	assert.WithinDuration(t, now, time.Now(), 10*time.Second)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHabits(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	req := &pb.GetHabitsRequest{
		UserId: "user1",
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "frequency", "created_at"}).
		AddRow(uuid.New().String(), req.UserId, "Exercise", "Daily exercise routine", "DAILY", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, user_id, name, description, frequency, created_at FROM habits WHERE id = $1`)).
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := habitTracker.GetHabits(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Habits, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
