package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/dto"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/entity"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/helpers"
	pb "github.com/osvaldosilitonga/hotel-and-resto/user-service/internal/pb_user_service"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/repositories"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/utils"
	"golang.org/x/crypto/bcrypt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServer
	UserRepo    repositories.UserRepo
	SessionRepo repositories.SessionRepo
}

func NewUserService(ur repositories.UserRepo, sr repositories.SessionRepo) *UserService {
	return &UserService{
		UserRepo:    ur,
		SessionRepo: sr,
	}
}

func (u *UserService) Login(ctx context.Context, payload *pb.LoginReq) (*pb.LoginRes, error) {
	isValid, msg := utils.LoginValidator(payload.Email, payload.Password)
	if !isValid {
		return nil, status.Errorf(codes.InvalidArgument, msg)
	}

	user, err := u.UserRepo.FindUserProfileByEmail(ctx, payload.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "email not found")
		}
		return nil, status.Errorf(codes.Internal, "internal server error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.User.Password), []byte(payload.Password)); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "wrong password")
	}

	// Generate JWT Token
	tokenPair, err := helpers.GenerateTokenPair(&user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to generate token")
	}

	// Store Token to Database Session Table
	now := time.Now().Unix()
	sessionData := entity.Sessions{
		RefreshToken: tokenPair["refresh_token"].(string),
		AccessToken:  tokenPair["access_token"].(string),
		Email:        user.User.Email,
		RoleID:       user.User.RoleID,
		Exp:          tokenPair["access_token_exp"].(int64),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := u.SessionRepo.Save(ctx, &sessionData); err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return nil, status.Errorf(codes.AlreadyExists, "violates unique constraint")
		}
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("fail to save user session. [ERR]: %v", err))
	}

	// TODO: Store Token to Cache Redis

	return &pb.LoginRes{
		Code:         0,
		Message:      "ok",
		AccessToken:  tokenPair["access_token"].(string),
		RefreshToken: tokenPair["refresh_token"].(string),
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
