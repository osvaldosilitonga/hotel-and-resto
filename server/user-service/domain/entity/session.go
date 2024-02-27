package entity

type Sessions struct {
	RefreshToken string
	AccessToken  string
	Email        string
	RoleID       int
	Exp          int64
	CreatedAt    int64
	UpdatedAt    int64
}
