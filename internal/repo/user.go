package repo

import (
	"errors"
	"fmt"
	"strconv"
)

type UserRepository struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUserByID(userID uint) (*User, *error) {
	validUserID := userID >= 10
	if !validUserID {
		err := errors.New("Invalid user ID")
		return nil, &err
	}
	return &User{
		ID:   strconv.FormatUint(uint64(userID), 10),
		Name: fmt.Sprintf("User %d", userID),
	}, nil
}
