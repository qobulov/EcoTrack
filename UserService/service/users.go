package service

import (
	"context"
	"fmt"
	"log"

	pb "EcoTrack/UserServis/genproto/protos"
	storage "EcoTrack/UserServis/storage/postgres"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Storage *storage.Server
}

func NewUserService(storage *storage.Server) *UserService {
	return &UserService{Storage: storage}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := u.Storage.CreateUser(ctx, req)
	if err != nil {
		log.Println("error while creating user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := u.Storage.GetUser(ctx, req)
	if err != nil {
		log.Println("error while getting user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res, err := u.Storage.UpdateUser(ctx, req)
	if err != nil {
		log.Println("error while updating user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := u.Storage.DeleteUser(ctx, req)
	if err != nil {
		log.Println("error while deleting user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	res, err := u.Storage.GetUserProfile(ctx, req)
	if err != nil {
		log.Println("error while getting user profile: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	res, err := u.Storage.UpdateUserProfile(ctx, req)
	if err != nil {
		log.Println("error while updating user profile: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
