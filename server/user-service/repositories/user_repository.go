package repositories

import (
	"context"
	"database/sql"
	"log"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/dto"
	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/entity"
)

type UserRepo interface {
	FindUserProfileByEmail(ctx context.Context, email string) (entity.UserProfile, error)
	Save(ctx context.Context, user dto.SaveUserReq) error
}

type userRepoImpl struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &userRepoImpl{
		DB: db,
	}
}

func (r *userRepoImpl) Save(ctx context.Context, user dto.SaveUserReq) error {
	// start transaction
	tx, err := r.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	defer tx.Commit()

	// insert to table users
	query := `
	INSERT INTO USERS(email, password, role_id)
	VALUES ($1, $2, $3)
	RETURNING ID;
	`

	var lastInsertedID string

	err = tx.QueryRowContext(ctx, query, user.Email, user.Password, user.RoleID).Scan(&lastInsertedID)
	if err != nil {
		tx.Rollback()
		return err
	}

	// insert to table user_details
	query = `
	INSERT INTO USER_DETAILS (user_id, name, phone, birth, address, gender)
	VALUES ($1, $2, $3, $4, $5, $6);
	`
	_, err = tx.ExecContext(ctx, query, lastInsertedID, user.Name, user.Phone, user.Birth, user.Address, user.Gender)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *userRepoImpl) FindUserProfileByEmail(ctx context.Context, email string) (entity.UserProfile, error) {
	userProfile := entity.UserProfile{}

	query := `
	SELECT U.ID, U.ROLE_ID, U.EMAIL, UD.NAME, UD.PHONE, UD.BIRTH, UD.ADDRESS, UD.GENDER
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
		log.Printf("error from QueryRowContext, \n[ERR] => %v", err)
		return userProfile, err
	}

	return userProfile, nil
}
