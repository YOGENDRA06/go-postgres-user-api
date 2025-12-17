package models

type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2"`
	DOB  string `json:"dob"  validate:"required"`
}

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	DOB  string `json:"dob"`
	Age  int    `json:"age"`
}
