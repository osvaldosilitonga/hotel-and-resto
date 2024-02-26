package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "github.com/osvaldosilitonga/hotel-and-resto/user-service/internal/pb_user_service"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/services"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/services/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	type data struct {
		Email    string
		Password string
	}

	testCase := []map[string]any{
		{
			"description": "Success",
			"data": data{
				Email:    "test@mail.com",
				Password: "password",
			},
			"expected_code": 0,
		},
		{
			"description": "Email Empty",
			"data": data{
				Password: "password",
			},
			"expected_code": 13,
		},
		{
			"description": "Password Empty",
			"data": data{
				Email: "test@mail.com",
			},
			"expected_code": 13,
		},
	}

	// Mock
	mockUserRepo := mocks.UserRepoImplMock{
		Mock: mock.Mock{},
	}

	userService := services.NewUserService(&mockUserRepo)

	for _, v := range testCase {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
		defer cancel()

		res, err := userService.Login(ctx, &pb.LoginReq{
			Email:    v["data"].(data).Email,
			Password: v["data"].(data).Password,
		})
		if err != nil {
			assert.Contains(t, err.Error(), "code = InvalidArgument")
			continue
		}

		code := res.Code
		assert.Equal(t, v["expected_code"].(int), int(code), fmt.Sprintf("expected: %v, actual: %v", v["expected_code"].(int), int(code)))
	}
}
