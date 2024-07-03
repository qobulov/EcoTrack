package service

import ("context"
	"User-Service/Storage/postgres"
	pb "User-Service/genproto/user-proto"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	UserRepo *postgres.User
}

func NewUserService(userRepo *postgres.User) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	return s.UserRepo.RegisterUser(ctx, request)
}