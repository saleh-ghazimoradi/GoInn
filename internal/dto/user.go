package dto

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUser struct {
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}
