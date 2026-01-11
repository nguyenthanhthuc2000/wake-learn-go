package service

import "golang/internal/repo"

type UserService struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUser() User {
	result := repo.NewUserRepository().GetUser()
	return User{
		ID:   result.ID,
		Name: result.Name,
	}
}
