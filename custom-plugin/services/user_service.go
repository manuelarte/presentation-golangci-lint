package services

import "custom-module/models"

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

//nolint:revive // to be implemented
func (s UserService) GetAll(page, size int) []models.User {
	// TODO(manuelarte): handle page and size
	return []models.User{
		{
			ID:       1,
			Nickname: "johndoe",
		},
		{
			ID:       2,
			Nickname: "alice",
		},
		{
			ID:       3,
			Nickname: "bob",
		},
	}
}
