package services

import "custom-module/models"

type UserService struct{}

func NewUserService() UserService {
	return UserService{}
}

func (s UserService) GetAll() []models.User {
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
