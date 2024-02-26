package utils

import (
	"errors"
	"fmt"
	"net/mail"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/dto"
)

func LoginValidator(email, password string) (bool, string) {
	if email == "" || password == "" {
		return false, "user or password cannot be null"
	}

	if len(email) < 8 || len(password) < 6 {
		return false, "please insert email min length 8, password min length 6"
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return false, "please insert correct email"
	}

	return true, ""
}

func SaveValidator(user dto.SaveUserReq) (bool, string) {
	err := CheckMin(8, user.Email)
	if err != nil {
		return false, fmt.Sprintf("email constraint. %v", err)
	}

	err = CheckMin(6, user.Password)
	if err != nil {
		return false, fmt.Sprintf("password constraint. %v", err)
	}

	err = CheckMin(5, user.Name)
	if err != nil {
		return false, fmt.Sprintf("name constraint. %v", err)
	}

	err = CheckMin(8, user.Phone)
	if err != nil {
		return false, fmt.Sprintf("phone constraint. %v", err)
	}

	err = CheckMin(8, user.Birth)
	if err != nil {
		return false, fmt.Sprintf("birth date constraint. %v", err)
	}

	err = CheckMin(6, user.Address)
	if err != nil {
		return false, fmt.Sprintf("address constraint. %v", err)
	}

	err = CheckMin(3, user.Gender)
	if err != nil {
		return false, fmt.Sprintf("gender constraint. %v", err)
	}

	err = CheckEmailFormat(user.Email)
	if err != nil {
		return false, err.Error()
	}

	return true, ""

}

func CheckEmailFormat(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("please insert correct email")
	}
	return nil
}

func CheckMin(min int, str string) error {
	if len(str) < min {
		return errors.New(fmt.Sprintf("cannot less then %v character", min))
	}
	return nil
}

func CheckMax(max int, str string) error {
	if len(str) > max {
		return errors.New(fmt.Sprintf("max character is %v", max))
	}
	return nil
}
