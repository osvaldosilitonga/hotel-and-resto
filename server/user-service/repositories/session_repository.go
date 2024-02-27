package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/entity"
)

type SessionRepo interface {
	Save(ctx context.Context, session *entity.Sessions) error
	Delete(ctx context.Context, sessionID string) error
}

type sessionRepoImpl struct {
	DB *sql.DB
}

func NewSessionRepo(db *sql.DB) SessionRepo {
	return &sessionRepoImpl{
		DB: db,
	}
}

func (s *sessionRepoImpl) Save(ctx context.Context, session *entity.Sessions) error {
	query := `
	INSERT INTO sessions (refresh_token, access_token, email, role_id, exp, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7);
	`

	_, err := s.DB.ExecContext(ctx, query, session.RefreshToken, session.AccessToken, session.Email, session.RoleID, session.Exp, session.CreatedAt, session.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (s *sessionRepoImpl) Delete(ctx context.Context, refreshToken string) error {
	query := `DELETE FROM sessions WHERE refresh_token = $1;`

	res, err := s.DB.ExecContext(ctx, query, refreshToken)
	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("no affected row")
	}

	return nil
}
