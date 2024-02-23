package entity

type UserDetails struct {
	UserID    string
	Name      string
	Phone     string
	Birth     string
	Address   string
	Gender    string
	CreatedAt string
	UpdatedAt string
}

type Users struct {
	ID       string
	RoleID   int
	Email    string
	Password string
}

type UserProfile struct {
	User        Users
	UserDetails UserDetails
}
