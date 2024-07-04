package postgres

import (
	pb "Habits-Tracker/genproto/protos"
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

type HabitTracker struct {
	db *sql.DB
}

func NewHabitTracker(db *sql.DB) *HabitTracker {
	return &HabitTracker{db: db}
}

func (ht *HabitTracker) CreateHabit(ctx context.Context, req *pb.CreateHabitRequest) (*pb.Habit, error) {
	id := uuid.New().String()
	query := `INSERT INTO habits (id, user_id, name, description, frequency) VALUES ($1, $2, $3, $4, $5)`
	_, err := ht.db.ExecContext(ctx, query, id, req.UserId, req.Name, req.Description, req.Frequency.String())
	if err != nil {
		log.Printf("Error creating habit: %v", err)
		return nil, err
	}
	return &pb.Habit{
		Id:          id,
		UserId:      req.UserId,
		Name:        req.Name,
		Description: req.Description,
		Frequency:   req.Frequency,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}, nil
}

func (ht *HabitTracker) GetHabits(ctx context.Context, req *pb.GetHabitsRequest) (*pb.GetHabitsResponse, error) {
	query := `SELECT id, user_id, name, description, frequency, created_at FROM habits WHERE id = $1`
	rows, err := ht.db.QueryContext(ctx, query, req.UserId)
	if err != nil {
		log.Printf("Error getting habits: %v", err)
		return nil, err
	}
	defer rows.Close()

	var habits []*pb.Habit
	for rows.Next() {
		var habit pb.Habit
		var frequency string
		if err := rows.Scan(&habit.Id, &habit.UserId, &habit.Name, &habit.Description, &frequency, &habit.CreatedAt); err != nil {
			log.Printf("Error scanning habit: %v", err)
			return nil, err
		}
		habit.Frequency = pb.Frequency(pb.Frequency_value[frequency])
		habits = append(habits, &habit)
	}
	return &pb.GetHabitsResponse{Habits: habits}, nil
}

func (ht *HabitTracker) LogHabit(ctx context.Context, req *pb.LogHabitRequest) (*pb.HabitLog, error) {
	id := uuid.New().String()
	query := `INSERT INTO habit_logs (id, habit_id, notes) VALUES ($1, $2, $3)`
	_, err := ht.db.ExecContext(ctx, query, id, req.HabitId, req.Notes)
	if err != nil {
		log.Printf("Error logging habit: %v", err)
		return nil, err
	}
	return &pb.HabitLog{
		Id:       id,
		HabitId:  req.HabitId,
		LoggedAt: time.Now().Format(time.RFC3339),
		Notes:    req.Notes,
	}, nil
}

func (ht *HabitTracker) GetHabitLogs(ctx context.Context, req *pb.GetHabitLogsRequest) (*pb.GetHabitLogsResponse, error) {
    query := `SELECT id, habit_id, logged_at, notes FROM habit_logs WHERE id = $1 ORDER BY logged_at DESC`
    rows, err := ht.db.QueryContext(ctx, query, req.HabitId)
    if err != nil {
        log.Printf("Error getting habit logs: %v", err)
        return nil, err
    }
    defer rows.Close()
	fmt.Println(rows)
    var logs []*pb.HabitLog
    for rows.Next() {
        var lg pb.HabitLog
        if err := rows.Scan(&lg.Id, &lg.HabitId, &lg.LoggedAt, &lg.Notes); err != nil {
            log.Printf("Error scanning habit log: %v", err)
            return nil, err
        }
        logs = append(logs, &lg)
		// fmt.Println(lg)
    }
	fmt.Println(logs)
    return &pb.GetHabitLogsResponse{HabitLogs: logs}, nil
}

func (ht *HabitTracker) DeleteHabit(ctx context.Context, req *pb.DeleteHabitRequest) (*pb.DeleteHabitResponse, error) {
	query := `DELETE FROM habits WHERE id = $1`
	result, err := ht.db.ExecContext(ctx, query, req.HabitId)
	if err != nil {
		log.Printf("Error deleting habit: %v", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, fmt.Errorf("habit with id %s not found", req.HabitId)
	}

	return &pb.DeleteHabitResponse{Message: "Habit deleted"}, nil
}

func (ht *HabitTracker) UpdateHabit(ctx context.Context, req *pb.UpdateHabitRequest) (*pb.Habit, error) {
	// Frequency qiymatini kichik harflarga o'zgartirish
	frequency := strings.ToLower(req.Frequency.String())

	query := `UPDATE habits SET name = $1, description = $2, frequency = $3 WHERE id = $4 RETURNING id, user_id, name, description, frequency, created_at`
	row := ht.db.QueryRowContext(ctx, query, req.Name, req.Description, frequency, req.HabitId)

	var habit pb.Habit
	var dbFrequency string
	err := row.Scan(&habit.Id, &habit.UserId, &habit.Name, &habit.Description, &dbFrequency, &habit.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("habit with id %s not found", req.HabitId)
		}
		log.Printf("Error updating habit: %v", err)
		return nil, err
	}

	// Frequency qiymatini to'g'ri enum turiga o'tkazish
	habitFrequency, ok := pb.Frequency_value[strings.ToUpper(dbFrequency)]
	if !ok {
		return nil, fmt.Errorf("invalid frequency value: %s", dbFrequency)
	}

	habit.Frequency = pb.Frequency(habitFrequency)
	return &habit, nil
}


func (ht *HabitTracker) GetUserHabits(ctx context.Context, req *pb.GetUserHabitsRequest) (*pb.GetUserHabitsResponse, error) {
	query := `SELECT id, user_id, name, description, frequency, created_at FROM habits WHERE user_id = $1`
	rows, err := ht.db.QueryContext(ctx, query, req.UserId)
	if err != nil {
		log.Printf("Error getting user habits: %v", err)
		return nil, err
	}
	defer rows.Close()

	var habits []*pb.Habit
	for rows.Next() {
		var habit pb.Habit
		var frequency string
		if err := rows.Scan(&habit.Id, &habit.UserId, &habit.Name, &habit.Description, &frequency, &habit.CreatedAt); err != nil {
			log.Printf("Error scanning habit: %v", err)
			return nil, err
		}
		habit.Frequency = pb.Frequency(pb.Frequency_value[frequency])
		habits = append(habits, &habit)
	}

	return &pb.GetUserHabitsResponse{Habits: habits}, nil
}

func (ht *HabitTracker) GetHabitSuggestions(ctx context.Context, req *pb.GetHabitSuggestionsRequest) (*pb.GetHabitSuggestionsResponse, error) {
	// This is a placeholder. In a real application, you might use some algorithm or criteria to generate suggestions.
	query := `SELECT id, user_id, name, description, frequency, created_at FROM habits ORDER BY created_at DESC LIMIT 10`
	rows, err := ht.db.QueryContext(ctx, query)
	if err != nil {
		log.Printf("Error getting habit suggestions: %v", err)
		return nil, err
	}
	defer rows.Close()

	var habits []*pb.Habit
	for rows.Next() {
		var habit pb.Habit
		var frequency string
		if err := rows.Scan(&habit.Id, &habit.UserId, &habit.Name, &habit.Description, &frequency, &habit.CreatedAt); err != nil {
			log.Printf("Error scanning habit: %v", err)
			return nil, err
		}
		habit.Frequency = pb.Frequency(pb.Frequency_value[frequency])
		habits = append(habits, &habit)
	}
	return &pb.GetHabitSuggestionsResponse{Habits: habits}, nil
}
