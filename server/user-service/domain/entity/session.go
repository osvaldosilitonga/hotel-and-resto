package entity

type Sessions struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Email        string `json:"email"`
	RoleID       int    `json:"role_id"`
	Exp          int64  `json:"exp"`
	CreatedAt    int64  `json:"created_at"`
	UpdatedAt    int64  `json:"updated_at"`
}
