package service

import (
    "context"
    "log"

	storage "EcoTrack/UserServis/storage/postgres"
    pb "EcoTrack/UserServis/genproto/protos"
)

type AuthService struct {
    pb.UnimplementedAuthServiceServer
    Storage *storage.AuthServer
}

func NewAuthService(storage *storage.AuthServer) *AuthService {
    return &AuthService{Storage: storage}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
    err := s.Storage.CreateUser(req.Username, req.Email, req.Password)
    if err != nil {
        log.Println("Foydalanuvchini yaratishda xatolik: ", err)
        return nil, err
    }
    return &pb.RegisterResponse{Message: "Foydalanuvchi muvaffaqiyatli ro'yxatdan o'tdi"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
    token, err := s.Storage.AuthenticateUser(req.Email, req.Password)
    if err != nil {
        log.Println("Tizimga kirishda xatolik: ", err)
        return nil, err
    }
    return &pb.LoginResponse{Token: token, Message: "Tizimga muvaffaqiyatli kirdingiz"}, nil
}

func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
    err := s.Storage.InvalidateToken(req.Token)
    if err != nil {
        log.Println("Tizimdan chiqishda xatolik: ", err)
        return nil, err
    }
    return &pb.LogoutResponse{Message: "Tizimdan muvaffaqiyatli chiqdingiz"}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenResponse, error) {
    newToken, err := s.Storage.RefreshToken(req.Token)
    if err != nil {
        log.Println("Tokenni yangilashda xatolik: ", err)
        return nil, err
    }
    return &pb.RefreshTokenResponse{NewToken: newToken, Message: "Token muvaffaqiyatli yangilandi"}, nil
}


