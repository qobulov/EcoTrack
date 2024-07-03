package postgres

import (
	"context"
	"database/sql"

	pb "User-Service/genproto/user-proto"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3)"

	_, err := u.db.ExecContext(ctx, query, request.Username, request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterUserResponse{Message: "User created successfully"}, nil
}

func (u *User) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	query := "SELECT id, name, email, password FROM users WHERE name = $1"

	row := u.db.QueryRowContext(ctx, query, request.Username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var id, name, email, password string

	err := row.Scan(&id, &name, &email, &password)
	if err != nil {
		return nil, err
	}
	if password != request.Password {
		return &pb.LoginUserResponse{UserId: id ,Message: "Invalid Password ,Can`t login"}, err
	} else {
		return &pb.LoginUserResponse{UserId: id, Message: "User logged in successfully"}, nil
	}
}

func (u *User) GetUserProfile(ctx context.Context, request *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	query := "SELECT id, name, email, full_name, bio, location, created_at FROM users WHERE id = $1"

	row := u.db.QueryRowContext(ctx, query, request.UserId)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var res pb.GetUserProfileResponse
	err := row.Scan(&res.UserId, &res.Username, &res.Email, &res.FullName, &res.Bio, &res.Location)
	if err != nil {
		return nil, err
	}
	return &res, nil
}


func (u *User) UpdateUserProfile(ctx context.Context, request *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	query := "UPDATE users SET full_name = $1, bio = $2, location = $3 WHERE id = $4"

	_, err := u.db.ExecContext(ctx, query, request.FullName, request.Bio, request.Location, request.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserProfileResponse{Message: "User profile updated successfully"}, nil
}


