package repositories

import (
	"context"
	"database/sql"

	"github.com/osvaldosilitonga/hotel-and-resto/user-service/domain/entity"
)

type SessionRepo interface {
	Save(ctx context.Context, session *entity.Sessions) error
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
