package entity

type Employees struct {
	ID       string
	Username string
	Password string
	Role     int
}

type EmployeeDetails struct {
	EmployeeID     string
	IdentityCard   string
	Name           string
	Phone          string
	Address        string
	Gender         string
	Salary         int
	EmployeeStatus string
	IsActive       bool
	CreatedAt      string
	UpdatedAt      string
}
