package signatures

import (
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"
	repoUser "e-signature/modules/v1/utilities/user/repository"
	serviceUser "e-signature/modules/v1/utilities/user/service"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type signaturesHandler struct {
	signaturesService service.Service
	serviceUser       serviceUser.Service
}

func NewSignaturesHandler(signaturesService service.Service, serviceUser serviceUser.Service) *signaturesHandler {
	return &signaturesHandler{signaturesService, serviceUser}
}

func Handler(db *mongo.Database, blockhain *api.Api, client *ethclient.Client) *signaturesHandler {
	//Signatures
	Repository := repository.NewRepository(db, blockhain, client)
	Service := service.NewService(Repository)
	//User
	RepoUser := repoUser.NewRepository(db, blockhain, client)
	serviceUser := serviceUser.NewService(RepoUser)
	return NewSignaturesHandler(Service, serviceUser)
}
