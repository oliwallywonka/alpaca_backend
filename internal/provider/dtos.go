package providerfx

type UpdateProviderDTO struct {
	Name        string
	Description string
	Contacts    string
	IsActive    *bool
	UpdatedAt   int
}

type UpdatePasswordDTO struct {
	Password string
}
