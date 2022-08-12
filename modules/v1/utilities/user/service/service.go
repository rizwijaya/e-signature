package service

import (
	"e-signature/modules/v1/utilities/user/models"
	"e-signature/modules/v1/utilities/user/repository"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) Login(input models.LoginInput) (models.User, error) {
	// idsignature := input.IdSignature
	// password := input.Password
	return models.User{}, nil
}
