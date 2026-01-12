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

func (us *UserService) GetUserByID(userID uint) (User, error) {
	result, err := repo.NewUserRepository().GetUserByID(userID)
	if err != nil {
		return User{}, *err
	}
	return User{
		ID:   result.ID,
		Name: result.Name,
	}, nil
}
