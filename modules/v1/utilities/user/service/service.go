package service

import (
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/user/models"
	"e-signature/modules/v1/utilities/user/repository"
)

type Service interface {
	Login(input models.LoginInput) (models.User, error)
}

type service struct {
	repository repository.Repository
	blockhain  *api.Api
}

func NewService(repository repository.Repository, blockhain *api.Api) *service {
	return &service{repository, blockhain}
}

func (s *service) Login(input models.LoginInput) (models.User, error) {
	// idsignature := input.IdSignature
	// password := input.Password
	return models.User{}, nil
}
