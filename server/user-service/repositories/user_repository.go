package repositories

import (
	"context"
	"database/sql"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/entity"
)

type UserRepo interface {
	FindUserProfileByEmail(ctx context.Context, email string) (entity.UserProfile, error)
}

type userRepoImpl struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		DB: db,
	}
}

func (r *userRepoImpl) FindUserProfileByEmail(ctx context.Context, email string) (entity.UserProfile, error) {
	userProfile := entity.UserProfile{}

	query := `
	SELECT U.ID, U.ROLE, U.EMAIL, UD.NAME, UD.PHONE, UD.BIRTH, UD.ADDRESS, UD.GENDER
	FROM USERS AS U
	JOIN USER_DETAILS AS UD ON U.ID = UD.USER_ID
	WHERE U.EMAIL = $1
	LIMIT 1
	`

	if err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&userProfile.User.ID,
		&userProfile.User.RoleID,
		&userProfile.User.Email,
		&userProfile.UserDetails.Name,
		&userProfile.UserDetails.Phone,
		&userProfile.UserDetails.Birth,
		&userProfile.UserDetails.Address,
		&userProfile.UserDetails.Gender,
	); err != nil {
		return userProfile, err
	}

	return userProfile, nil
}
