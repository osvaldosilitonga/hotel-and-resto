package dto

type SaveUserReq struct {
	Email    string `json:"email"`
	Password string `json:"-"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Birth    string `json:"birth"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	RoleID   int    `json:"role_id"`
}
