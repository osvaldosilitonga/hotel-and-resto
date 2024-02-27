package entity

type Sessions struct {
	RefreshToken string
	AccessToken  string
	Email        string
	RoleID       int
	Exp          int
	CreatedAt    int
	UpdatedAt    int
}
