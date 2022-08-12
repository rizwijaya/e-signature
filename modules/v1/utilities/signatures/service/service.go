package service

import (
	"e-signature/modules/v1/utilities/signatures/repository"
)

type Service interface {
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
