package postgres

import (
	"context"
	"database/sql"

	pb "EcoTrack/UserServis/genproto/protos"

	_ "github.com/lib/pq"
)

type Server struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var id string
	err := s.db.QueryRowContext(ctx, "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id", req.Username, req.Email, req.Password).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{User: &pb.User{Id: id, Username: req.Username, Email: req.Email}, Message: "User Logged in successfully!"}, nil
}

func (s *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var username, email, createdAt, updatedAt string
	err := s.db.QueryRowContext(ctx, "SELECT username, email, created_at, updated_at FROM users WHERE id = $1", req.UserId).Scan(&username, &email, &createdAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{
		UserId:    req.UserId,
		Username:  username,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE users SET username = $1, email = $2 WHERE id = $3", req.Username, req.Email, req.UserId)
	if err != nil {
		return &pb.UpdateUserResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.UpdateUserResponse{Success: true, Message: "User updated successfully"}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	_, err := s.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", req.UserId)
	if err != nil {
		return &pb.DeleteUserResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.DeleteUserResponse{Success: true, Message: "User deleted successfully"}, nil
}

func (s *Server) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	var fullName, bio, location, avatarUrl string
	err := s.db.QueryRowContext(ctx, "SELECT full_name, bio, location, avatar_url FROM user_profiles WHERE user_id = $1", req.UserId).Scan(&fullName, &bio, &location, &avatarUrl)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserProfileResponse{
		UserId:    req.UserId,
		FullName:  fullName,
		Bio:       bio,
		Location:  location,
		AvatarUrl: avatarUrl,
	}, nil
}

func (s *Server) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	_, err := s.db.ExecContext(ctx, "UPDATE user_profiles SET full_name = $1, bio = $2, location = $3, avatar_url = $4 WHERE user_id = $5",
		req.FullName, req.Bio, req.Location, req.AvatarUrl, req.UserId)
	if err != nil {
		return &pb.UpdateUserProfileResponse{Success: false, Message: err.Error()}, nil
	}
	return &pb.UpdateUserProfileResponse{Success: true, Message: "User profile updated successfully"}, nil
}
