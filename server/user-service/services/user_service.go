package services

import (
	"context"

	pb "github.com/osvaldosilitonga/hotel-and-resto/user-service/internal/pb_user_service"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) Login(ctx context.Context, payload *pb.LoginReq) (*pb.LoginRes, error) {
	isValid, msg := utils.LoginValidator(payload.Email, payload.Password)
	if !isValid {
		return nil, status.Errorf(codes.Internal, msg)
	}

	return &pb.LoginRes{
		Code:        200,
		Message:     "ok",
		AccessToken: "access_token_test",
	}, nil
}

func (u *UserService) Save(ctx context.Context, payload *pb.SaveReq) (*pb.SaveRes, error) {
	return &pb.SaveRes{
		Code:    200,
		Message: "ok",
		Data: &pb.UserData{
			Email:   "test@mail.com",
			Name:    "test name",
			Phone:   "1234567",
			Birth:   "22/02/2024",
			Address: "test address",
			Gender:  "male",
		},
	}, nil
}

func (u *UserService) FindByEmail(ctx context.Context, payload *pb.FindByEmailReq) (*pb.FindByEmailRes, error) {
	return &pb.FindByEmailRes{
		Code:    200,
		Message: "ok",
		Data: &pb.UserData{
			Email:   "test@mail.com",
			Name:    "test name",
			Phone:   "1234567",
			Birth:   "22/02/2024",
			Address: "test address",
			Gender:  "male",
		},
	}, nil
}
