package services

import (
	"context"
	"database/sql"
	"errors"

	pb "github.com/osvaldosilitonga/hotel-and-resto/user-service/internal/pb_user_service"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/repositories"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServer
	UserRepo repositories.UserRepo
}

func NewUserService(ur repositories.UserRepo) *UserService {
	return &UserService{
		UserRepo: ur,
	}
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
	profile, err := u.UserRepo.FindUserProfileByEmail(ctx, payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "email not found")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.FindByEmailRes{
		Code:    200,
		Message: "ok",
		Data: &pb.UserData{
			Email:   profile.User.Email,
			Name:    profile.UserDetails.Name,
			Phone:   profile.UserDetails.Phone,
			Birth:   profile.UserDetails.Birth,
			Address: profile.UserDetails.Address,
			Gender:  profile.UserDetails.Gender,
		},
	}, nil
}
