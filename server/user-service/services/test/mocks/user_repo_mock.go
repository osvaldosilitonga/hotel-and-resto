package mocks

import (
	"context"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/dto"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/entity"
	"github.com/stretchr/testify/mock"
)

type UserRepoImplMock struct {
	mock.Mock
}

func (u *UserRepoImplMock) FindUserProfileByEmail(ctx context.Context, email string) (entity.UserProfile, error) {
	return entity.UserProfile{}, nil
}

func (u *UserRepoImplMock) Save(ctx context.Context, user dto.SaveUserReq) error {
	return nil
}
