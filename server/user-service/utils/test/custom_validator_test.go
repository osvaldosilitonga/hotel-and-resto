package test

import (
	"fmt"
	"testing"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/dto"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/utils"

	"github.com/stretchr/testify/assert"
)

func TestLoginValidator(t *testing.T) {
	type data struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	testCase := []map[string]any{
		{
			"description": "Email kurang dari 6",
			"data": data{
				Email:    "a@b.com",
				Password: "testpass",
			},
			"expected":     false,
			"expected_msg": "please insert email min length 8, password min length 6",
		},
		{
			"description": "Format email salah",
			"data": data{
				Email:    "ab.com",
				Password: "testpass",
			},
			"expected":     false,
			"expected_msg": "please insert correct email",
		},
		{
			"description": "Email kosong",
			"data": data{
				Email:    "",
				Password: "testpass",
			},
			"expected":     false,
			"expected_msg": "user or password cannot be null",
		},
		{
			"description": "Success",
			"data": data{
				Email:    "test@gmail.com",
				Password: "testpass",
			},
			"expected":     true,
			"expected_msg": "",
		},
	}

	for _, v := range testCase {
		d := v["data"].(data)
		isValid, _ := utils.LoginValidator(d.Email, d.Password)

		assert.Equal(t, v["expected"].(bool), isValid, "isValid must be False")
	}
}

func TestSaveValidator(t *testing.T) {

	testCase := []map[string]any{
		{
			"description": "Success",
			"data": dto.SaveUserReq{
				Email:    "test@mail.com",
				Password: "password",
				Name:     "test123",
				Phone:    "12345678",
				Birth:    "1992-10-20",
				Address:  "ragunan",
				Gender:   "male",
				RoleID:   1,
			},
			"expected":     true,
			"expected_msg": "",
		},
		{
			"description": "Fail - Payload Email empty",
			"data": dto.SaveUserReq{
				Password: "password",
				Name:     "test123",
				Phone:    "12345678",
				Birth:    "1992-10-20",
				Address:  "ragunan",
				Gender:   "male",
				RoleID:   1,
			},
			"expected":     false,
			"expected_msg": "email constraint",
		},
		{
			"description": "Fail - Payload Password empty",
			"data": dto.SaveUserReq{
				Email:   "test@mail.com",
				Name:    "test123",
				Phone:   "12345678",
				Birth:   "1992-10-20",
				Address: "ragunan",
				Gender:  "male",
				RoleID:  1,
			},
			"expected":     false,
			"expected_msg": "password constraint",
		},
		{
			"description": "Fail - Payload Name empty",
			"data": dto.SaveUserReq{
				Email:    "test@mail.com",
				Password: "password",
				Phone:    "12345678",
				Birth:    "1992-10-20",
				Address:  "ragunan",
				Gender:   "male",
				RoleID:   1,
			},
			"expected":     false,
			"expected_msg": "name constraint",
		},
		{
			"description": "Fail - Payload Phone empty",
			"data": dto.SaveUserReq{
				Email:    "test@mail.com",
				Password: "password",
				Name:     "test123",
				Birth:    "1992-10-20",
				Address:  "ragunan",
				Gender:   "male",
				RoleID:   1,
			},
			"expected":     false,
			"expected_msg": "phone constraint",
		},
		{
			"description": "Fail - Payload Birth empty",
			"data": dto.SaveUserReq{
				Email:    "test@mail.com",
				Password: "password",
				Name:     "test123",
				Phone:    "12345678",
				Address:  "ragunan",
				Gender:   "male",
				RoleID:   1,
			},
			"expected":     false,
			"expected_msg": "birth date constraint",
		},
		{
			"description": "Fail - Payload Address empty",
			"data": dto.SaveUserReq{
				Email:    "test@mail.com",
				Password: "password",
				Name:     "test123",
				Phone:    "12345678",
				Birth:    "1992-10-20",
				Gender:   "male",
				RoleID:   1,
			},
			"expected":     false,
			"expected_msg": "address constraint",
		},
		{
			"description": "Fail - Payload Gender empty",
			"data": dto.SaveUserReq{
				Email:    "test@mail.com",
				Password: "password",
				Name:     "test123",
				Phone:    "12345678",
				Birth:    "1992-10-20",
				Address:  "ragunan",
				RoleID:   1,
			},
			"expected":     false,
			"expected_msg": "gender constraint",
		},
	}

	for _, v := range testCase {
		d := v["data"].(dto.SaveUserReq)
		isValid, msg := utils.SaveValidator(d)

		assert.Equal(t, v["expected"].(bool), isValid, fmt.Sprintf("actual: %v, expected: %v, msg: %v", isValid, v["expected"].(bool), msg))
		assert.Contains(t, msg, v["expected_msg"].(string), fmt.Sprintf("actual: %v, expected: message must contains = %v", msg, v["expected_msg"].(string)))
	}

}
