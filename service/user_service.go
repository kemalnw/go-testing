package service

import (
	"go-testing/repository"
)

type UserService struct {
	repo repository.Repository
}

func NewUserService(userRepo repository.Repository) *UserService {
	return &UserService{repo: userRepo}
}

func (s *UserService) GetUsername(id int64) (string, error) {
	return s.repo.GetUserNameByID(id)
}
