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

func TestLogHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	now := time.Now()
	req := &pb.LogHabitRequest{
		HabitId:  "habit1",
		Notes:    "Completed workout",
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO habit_logs (id, habit_id, notes) VALUES ($1, $2, $3)`)).
		WithArgs(sqlmock.AnyArg(), req.HabitId, req.Notes).
		WillReturnResult(sqlmock.NewResult(1, 1))

	habitLog, err := habitTracker.LogHabit(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, habitLog)
	assert.NotEmpty(t, habitLog.Id)
	assert.Equal(t, req.HabitId, habitLog.HabitId)
	assert.Equal(t, req.Notes, habitLog.Notes)
	assert.WithinDuration(t, now, time.Now(), 10*time.Second)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHabitLogs(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	req := &pb.GetHabitLogsRequest{
		HabitId: "habit1",
	}
	rows := sqlmock.NewRows([]string{"id", "habit_id", "logged_at", "notes"}).
		AddRow(uuid.New().String(), req.HabitId, time.Now().Format(time.RFC3339), "Completed workout")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, habit_id, logged_at, notes FROM habit_logs WHERE id = $1 ORDER BY logged_at DESC`)).
		WithArgs(req.HabitId).
		WillReturnRows(rows)

	resp, err := habitTracker.GetHabitLogs(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.HabitLogs, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	req := &pb.DeleteHabitRequest{
		HabitId: "habit1",
	}

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM habits WHERE id = $1`)).
		WithArgs(req.HabitId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := habitTracker.DeleteHabit(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Habit deleted", resp.Message)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateHabit(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	req := &pb.UpdateHabitRequest{
		HabitId:      "habit1",
		Name:        "Updated Exercise",
		Description: "Updated exercise routine",
		Frequency:   pb.Frequency_WEEKLY,
	}

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE habits SET name = $1, description = $2, frequency = $3 WHERE id = $4 RETURNING id, user_id, name, description, frequency, created_at`)).
		WithArgs(req.Name, req.Description, req.Frequency.String(), req.HabitId).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "name", "description", "frequency", "created_at"}).
			AddRow(req.HabitId, "user1", req.Name, req.Description, req.Frequency.String(), time.Now().Format(time.RFC3339)))

	updatedHabit, err := habitTracker.UpdateHabit(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, updatedHabit)
	assert.Equal(t, req.HabitId, updatedHabit.Id)
	assert.Equal(t, req.Name, updatedHabit.Name)
	assert.Equal(t, req.Description, updatedHabit.Description)
	assert.Equal(t, req.Frequency, updatedHabit.Frequency)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserHabits(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	req := &pb.GetUserHabitsRequest{
		UserId: "user1",
	}
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "frequency", "created_at"}).
		AddRow(uuid.New().String(), req.UserId, "Exercise", "Daily exercise routine", "DAILY", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, user_id, name, description, frequency, created_at FROM habits WHERE user_id = $1`)).
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := habitTracker.GetUserHabits(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Habits, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetHabitSuggestions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	habitTracker := NewHabitTracker(db)
	ctx := context.Background()
	req := &pb.GetHabitSuggestionsRequest{}
	rows := sqlmock.NewRows([]string{"id", "user_id", "name", "description", "frequency", "created_at"}).
		AddRow(uuid.New().String(), "user1", "Exercise", "Daily exercise routine", "DAILY", time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, user_id, name, description, frequency, created_at FROM habits ORDER BY created_at DESC LIMIT 10`)).
		WillReturnRows(rows)

	resp, err := habitTracker.GetHabitSuggestions(ctx, req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Habits, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
