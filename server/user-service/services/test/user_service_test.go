package test

import (
	"testing"

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

		assert.Equal(t, isValid, v["expected"].(bool), "isValid must be False")
	}
}
