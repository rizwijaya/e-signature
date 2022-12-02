package signatures

import (
	api "e-signature/app/contracts"
	"e-signature/modules/v1/utilities/signatures/repository"
	"e-signature/modules/v1/utilities/signatures/service"
	repoUser "e-signature/modules/v1/utilities/user/repository"
	serviceUser "e-signature/modules/v1/utilities/user/service"
	bl "e-signature/pkg/blockchain"
	pw "e-signature/pkg/crypto"
	docs "e-signature/pkg/document"
	images "e-signature/pkg/images"

	"github.com/ethereum/go-ethereum/ethclient"
	"go.mongodb.org/mongo-driver/mongo"
)

type signaturesHandler struct {
	serviceSignature service.Service
	serviceUser      serviceUser.Service
}

func NewSignaturesHandler(serviceSignature service.Service, serviceUser serviceUser.Service) *signaturesHandler {
	return &signaturesHandler{serviceSignature, serviceUser}
}

func Handler(db *mongo.Database, contracts *api.Api, client *ethclient.Client) *signaturesHandler {
	blockchain := bl.NewBlockchain(contracts, client)
	//Signatures
	Repository := repository.NewRepository(db, blockchain)
	documents := docs.NewDocuments()
	images := images.NewImages()
	Service := service.NewService(Repository, images, documents)
	//User
	RepoUser := repoUser.NewRepository(db, blockchain)
	crypto := pw.NewCrypto()
	serviceUser := serviceUser.NewService(RepoUser, crypto)
	return NewSignaturesHandler(Service, serviceUser)
}
