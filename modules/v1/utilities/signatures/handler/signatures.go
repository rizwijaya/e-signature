package signatures

import (
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"

	"go.mongodb.org/mongo-driver/mongo"
)

// type SignaturesHandler interface {
// }

type signaturesHandler struct {
	signaturesService service.Service
}

func NewSignaturesHandler(signaturesService service.Service) *signaturesHandler {
	return &signaturesHandler{signaturesService}
}

func Handler(db *mongo.Database) *signaturesHandler {
	Repository := repository.NewRepository(db)
	Service := service.NewService(Repository)
	return NewSignaturesHandler(Service)
}
