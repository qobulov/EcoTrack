package postgres

import (
	// "database/sql"
	"testing"

	pb "EcoTrack/UserServis/genproto/protos"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	req := &pb.GetUserRequest{UserId: "1"}

	rows := sqlmock.NewRows([]string{"username", "email", "created_at", "updated_at"}).
		AddRow("testuser", "test@example.com", "2024-01-01", "2024-01-01")

	mock.ExpectQuery("SELECT username, email, created_at, updated_at FROM users WHERE id = \\$1").
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := repo.GetUser(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.Username != "testuser" {
		t.Errorf("expected username 'testuser', got %v", resp.Username)
	}
	if resp.Email != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got %v", resp.Email)
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	req := &pb.UpdateUserRequest{UserId: "1", Username: "newuser", Email: "new@example.com"}

	mock.ExpectExec("UPDATE users SET username = \\$1, email = \\$2 WHERE id = \\$3").
		WithArgs(req.Username, req.Email, req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.UpdateUser(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !resp.Success {
		t.Errorf("expected success true, got %v", resp.Success)
	}
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	req := &pb.DeleteUserRequest{UserId: "1"}

	mock.ExpectExec("DELETE FROM users WHERE id = \\$1").
		WithArgs(req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteUser(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !resp.Success {
		t.Errorf("expected success true, got %v", resp.Success)
	}
}

func TestGetUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	req := &pb.GetUserProfileRequest{UserId: "1"}

	rows := sqlmock.NewRows([]string{"full_name", "bio", "location", "avatar_url"}).
		AddRow("Test User", "Bio", "Location", "http://example.com/avatar.jpg")

	mock.ExpectQuery("SELECT full_name, bio, location, avatar_url FROM user_profiles WHERE user_id = \\$1").
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := repo.GetUserProfile(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if resp.FullName != "Test User" {
		t.Errorf("expected full_name 'Test User', got %v", resp.FullName)
	}
	if resp.Bio != "Bio" {
		t.Errorf("expected bio 'Bio', got %v", resp.Bio)
	}
}

func TestUpdateUserProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	req := &pb.UpdateUserProfileRequest{
		UserId:    "1",
		FullName:  "New User",
		Bio:       "New Bio",
		Location:  "New Location",
		AvatarUrl: "http://example.com/newavatar.jpg",
	}

	mock.ExpectExec("UPDATE user_profiles SET full_name = $1, bio = $2, location = $3, avatar_url = $4 WHERE user_id = $5").
		WithArgs(req.FullName, req.Bio, req.Location, req.AvatarUrl, req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.UpdateUserProfile(req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !resp.Success {
		t.Errorf("expected success true, got %v", resp.Success)
	}
}
