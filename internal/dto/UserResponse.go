package dto

type UserResponse struct {
	FirstName  string `json:"first_name,omitempty"`
	SecondName string `json:"second_name,omitempty"`
	Surname    string `json:"surname,omitempty"`
}
