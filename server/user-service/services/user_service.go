package services

import (
	"context"
	"database/sql"
	"errors"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/dto"
	pb "github.com/osvaldosilitonga/hotel-and-resto/user-service/internal/pb_user_service"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/repositories"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/utils"
	"golang.org/x/crypto/bcrypt"

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
	res := &pb.SaveRes{}

	user := dto.SaveUserReq{
		Email:    payload.Email,
		Password: payload.Password,
		Name:     payload.Name,
		Phone:    payload.Phone,
		Birth:    payload.Birth,
		Address:  payload.Birth,
		Gender:   payload.Gender,
		RoleID:   1,
	}

	isValid, msg := utils.SaveValidator(user)
	if !isValid {
		return res, status.Errorf(codes.InvalidArgument, msg)
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed when hashing password")
	}

	user.Password = string(hash)

	err = u.UserRepo.Save(ctx, user)
	if err != nil {
		return res, err
	}

	res.Code = 200
	res.Message = "ok"
	res.Data = &pb.UserData{
		Email:   user.Email,
		Name:    user.Name,
		Phone:   user.Phone,
		Birth:   user.Birth,
		Address: user.Address,
		Gender:  user.Gender,
	}

	return res, nil
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
