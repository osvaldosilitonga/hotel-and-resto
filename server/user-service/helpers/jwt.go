package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/entity"
)

func GenerateTokenPair(user *entity.UserProfile) (map[string]any, error) {
	// Generate Access Token - 15 minute
	accessExp := time.Now().Add(time.Minute * 15).Unix()
	accessToken, err := GenerateToken(user.User.Email, user.User.RoleID, accessExp)
	if err != nil {
		return nil, err
	}

	// Generate Resfresh Token - 24 hour
	refreshExp := time.Now().Add(time.Hour * 24).Unix()
	refreshToken, err := GenerateToken("", 0, refreshExp)
	if err != nil {
		return nil, err
	}

	res := map[string]any{
		"access_token":     accessToken,
		"refresh_token":    refreshToken,
		"access_token_exp": accessExp,
	}

	return res, nil
}

func GenerateToken(email string, roleId int, exp int64) (string, error) {
	// Refresh Token
	if email == "" || roleId == 0 {
		refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": exp,
		})

		refreshToken, err := token.SignedString([]byte(refreshTokenSecret))
		if err != nil {
			return "", err
		}

		return refreshToken, nil
	}

	// Access Token
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  roleId,
		"exp":   time.Now().Add(time.Minute * 15).Unix(),
	})

	accessToken, err := token.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
