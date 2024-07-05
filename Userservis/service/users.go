package service

import (
	storage "EcoTrack/UserServis/storage/postgres"
	pb "EcoTrack/UserServis/genproto/protos"
	"context"
	"fmt"
	"log"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Storage *storage.Server
}

func NewUserService(storage *storage.Server) *UserService {
	return &UserService{Storage: storage}
}

func(u *UserService) Login(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	res, err := u.Storage.Login(req)
	if err != nil {
		log.Println("error while creating user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	res, err := u.Storage.GetUser(req)
	if err != nil {
		log.Println("error while getting user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	res, err := u.Storage.UpdateUser(req)
	if err != nil {
		log.Println("error while updating user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	res, err := u.Storage.DeleteUser(req)
	if err != nil {
		log.Println("error while deleting user: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) GetUserProfile(ctx context.Context, req *pb.GetUserProfileRequest) (*pb.GetUserProfileResponse, error) {
	res, err := u.Storage.GetUserProfile(req)
	if err != nil {
		log.Println("error while getting user profile: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}

func (u *UserService) UpdateUserProfile(ctx context.Context, req *pb.UpdateUserProfileRequest) (*pb.UpdateUserProfileResponse, error) {
	res, err := u.Storage.UpdateUserProfile(req)
	if err != nil {
		log.Println("error while updating user profile: ", err)
		return nil, err
	}
	fmt.Println(res)
	return res, nil
}
