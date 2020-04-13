package service

//UserRole ...
type UserRole struct {
	name     string
	isActive bool
}

//NewUserRole ...
func NewUserRole(role string, isActive bool) *UserRole {
	return &UserRole{
		name:     role,
		isActive: isActive,
	}
}
