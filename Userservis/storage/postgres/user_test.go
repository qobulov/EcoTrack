package postgres

import (
	"context"
	"testing"

	pb "EcoTrack/UserServis/genproto/protos"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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

	mock.ExpectQuery("SELECT username, email, created_at, updated_at FROM users WHERE id = $1").
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := repo.GetUser(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, "testuser", resp.Username)
	assert.Equal(t, "test@example.com", resp.Email)
	assert.Equal(t, "2024-01-01", resp.CreatedAt)
	assert.Equal(t, "2024-01-01", resp.UpdatedAt)
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	req := &pb.UpdateUserRequest{UserId: "1", Username: "newuser", Email: "new@example.com"}

	mock.ExpectExec("UPDATE users SET username = $1, email = $2 WHERE id = $3").
		WithArgs(req.Username, req.Email, req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.UpdateUser(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.True(t, resp.Success)
	assert.Equal(t, "User updated successfully", resp.Message)
}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error opening a stub database connection: %v", err)
	}
	defer db.Close()

	repo := NewUserRepo(db)
	req := &pb.DeleteUserRequest{UserId: "1"}

	mock.ExpectExec("DELETE FROM users WHERE id = $1").
		WithArgs(req.UserId).
		WillReturnResult(sqlmock.NewResult(1, 1))

	resp, err := repo.DeleteUser(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.True(t, resp.Success)
	assert.Equal(t, "User deleted successfully", resp.Message)
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

	mock.ExpectQuery("SELECT full_name, bio, location, avatar_url FROM user_profiles WHERE user_id = $1").
		WithArgs(req.UserId).
		WillReturnRows(rows)

	resp, err := repo.GetUserProfile(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.Equal(t, "Test User", resp.FullName)
	assert.Equal(t, "Bio", resp.Bio)
	assert.Equal(t, "Location", resp.Location)
	assert.Equal(t, "http://example.com/avatar.jpg", resp.AvatarUrl)
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

	resp, err := repo.UpdateUserProfile(context.Background(), req)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	assert.True(t, resp.Success)
	assert.Equal(t, "User profile updated successfully", resp.Message)
}
