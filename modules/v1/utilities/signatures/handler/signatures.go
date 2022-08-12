package signatures

import (
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"

	"gorm.io/gorm"
)

type SignaturesHandler interface {
}

type signaturesHandler struct {
	signaturesService service.Service
}

func NewSignaturesHandler(signaturesService service.Service) *signaturesHandler {
	return &signaturesHandler{signaturesService}
}

func Handler(db *gorm.DB) *signaturesHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewSignaturesHandler(Service)
}
