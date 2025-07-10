package userfx

type UpdateUserDTO struct {
	RoleID    string
	Email     string
	Password  string
	Name      string
	Image     string
	Contacts  string
	IsActive  *bool
	UpdatedAt int
}

type UpdatePasswordDTO struct {
	Password string
}
