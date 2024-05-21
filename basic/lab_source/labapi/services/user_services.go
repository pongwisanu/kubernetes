package services

import "labapi/repositories"

type userServices struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return userServices{userRepo: userRepo}
}

func (s userServices) GetUsers() ([]UserResponse, error) {

	users, err := s.userRepo.GetUsers()

	if err != nil {
		return nil, err
	}

	userResponses := []UserResponse{}

	for _, user := range users {
		userResponse := UserResponse{
			UserId: user.Id,
			Name:   user.Name,
			Detail: user.Detail,
			Role:   user.Detail,
		}
		userResponses = append(userResponses, userResponse)
	}

	return userResponses, nil
}

func (s userServices) GetUser(id int) (*UserResponse, error) {
	user, err := s.userRepo.GetUser(id)
	if err != nil {
		return nil, err
	}

	userResponse := UserResponse{
		UserId: user.Id,
		Name:   user.Name,
		Detail: user.Detail,
		Role:   user.Role,
	}

	return &userResponse, nil
}

func (s userServices) AddUser(request UserRequest) (*UserResponse, error) {

	user := repositories.NewUser{
		Name:   request.Name,
		Detail: request.Detail,
		Role:   request.Role,
	}

	newUser, err := s.userRepo.AddUser(user)
	if err != nil {
		return nil, err
	}

	response := UserResponse{
		UserId: newUser.Id,
		Name:   newUser.Name,
		Detail: newUser.Detail,
		Role:   newUser.Role,
	}

	return &response, nil
}
