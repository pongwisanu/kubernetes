package services

type UserRequest struct {
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Role   string `json:"role"`
}

type UserResponse struct {
	UserId int    `json:"id"`
	Name   string `json:"name"`
	Detail string `json:"detail"`
	Role   string `json:"role"`
}

type UserService interface {
	GetUsers() ([]UserResponse, error)
	GetUser(int) (*UserResponse, error)
	AddUser(UserRequest) (*UserResponse, error)
	// EditUser(*UserResponse) error
	// DeleteUser(int) error
}
