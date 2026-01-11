package repo

type UserRepository struct {
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUser() User {
	return User{
		ID:   "1",
		Name: "John Doe",
	}
}
