package service

import (
	"e-signature/modules/v1/utilities/signatures/models"
	m_repo "e-signature/modules/v1/utilities/signatures/repository/mock"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	m_docs "e-signature/pkg/document/mock"
	m_images "e-signature/pkg/images/mock"
	"errors"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tkuchiki/faketime"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestInit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Chdir("../../../../../")
}

func Test_service_NewService(t *testing.T) {
	t.Run("Case 1: Success New Service", func(t *testing.T) {
		serv := NewService(nil, nil, nil)
		assert.NotNil(t, serv)
	})
}

func Test_service_TimeFormating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	// location, err := time.LoadLocation("Asia/Jakarta")
	// assert.NoError(t, err)
	// times := time.Now().In(location)

	test := []struct {
		name       string
		time       string
		timeFormat string
	}{
		{
			name:       "Time Formatting Case 1: Success Convert Format Time 14 Digit",
			time:       "15040528022022",
			timeFormat: "Senin, 28 Feb 2022 | 15:04 WIB",
		},
		{
			name:       "Time Formatting Case 2: Success Convert Format Time 13 Digit",
			time:       "7040526022022",
			timeFormat: "Sabtu, 26 Feb 2022 | 07:04 WIB",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)
			s := NewService(repo, images, docs)
			timeformat := s.TimeFormating(tt.time)
			assert.Equal(t, tt.timeFormat, timeformat)
		})
	}
}

func Test_service_CreateImgSignature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name   string
		input  models.AddSignature
		output string
		test   func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			name: "Create Image Signature Case 1: Success Create Image Signature",
			input: models.AddSignature{
				Id:        "6380b5cbdc938c5fdf8e6bfe",
				Signature: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAApgAAAKYB3X3/OAAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAANCSURBVEiJtZZPbBtFFMZ/M7ubXdtdb1xSFyeilBapySVU8h8OoFaooFSqiihIVIpQBKci6KEg9Q6H9kovIHoCIVQJJCKE1ENFjnAgcaSGC6rEnxBwA04Tx43t2FnvDAfjkNibxgHxnWb2e/u992bee7tCa00YFsffekFY+nUzFtjW0LrvjRXrCDIAaPLlW0nHL0SsZtVoaF98mLrx3pdhOqLtYPHChahZcYYO7KvPFxvRl5XPp1sN3adWiD1ZAqD6XYK1b/dvE5IWryTt2udLFedwc1+9kLp+vbbpoDh+6TklxBeAi9TL0taeWpdmZzQDry0AcO+jQ12RyohqqoYoo8RDwJrU+qXkjWtfi8Xxt58BdQuwQs9qC/afLwCw8tnQbqYAPsgxE1S6F3EAIXux2oQFKm0ihMsOF71dHYx+f3NND68ghCu1YIoePPQN1pGRABkJ6Bus96CutRZMydTl+TvuiRW1m3n0eDl0vRPcEysqdXn+jsQPsrHMquGeXEaY4Yk4wxWcY5V/9scqOMOVUFthatyTy8QyqwZ+kDURKoMWxNKr2EeqVKcTNOajqKoBgOE28U4tdQl5p5bwCw7BWquaZSzAPlwjlithJtp3pTImSqQRrb2Z8PHGigD4RZuNX6JYj6wj7O4TFLbCO/Mn/m8R+h6rYSUb3ekokRY6f/YukArN979jcW+V/S8g0eT/N3VN3kTqWbQ428m9/8k0P/1aIhF36PccEl6EhOcAUCrXKZXXWS3XKd2vc/TRBG9O5ELC17MmWubD2nKhUKZa26Ba2+D3P+4/MNCFwg59oWVeYhkzgN/JDR8deKBoD7Y+ljEjGZ0sosXVTvbc6RHirr2reNy1OXd6pJsQ+gqjk8VWFYmHrwBzW/n+uMPFiRwHB2I7ih8ciHFxIkd/3Omk5tCDV1t+2nNu5sxxpDFNx+huNhVT3/zMDz8usXC3ddaHBj1GHj/As08fwTS7Kt1HBTmyN29vdwAw+/wbwLVOJ3uAD1wi/dUH7Qei66PfyuRj4Ik9is+hglfbkbfR3cnZm7chlUWLdwmprtCohX4HUtlOcQjLYCu+fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII=",
			},
			output: "public/images/signatures/signatures/signatures-example.png",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.AddSignature{
					Id:        "6380b5cbdc938c5fdf8e6bfe",
					Signature: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABgAAAAYCAYAAADgdz34AAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAApgAAAKYB3X3/OAAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAANCSURBVEiJtZZPbBtFFMZ/M7ubXdtdb1xSFyeilBapySVU8h8OoFaooFSqiihIVIpQBKci6KEg9Q6H9kovIHoCIVQJJCKE1ENFjnAgcaSGC6rEnxBwA04Tx43t2FnvDAfjkNibxgHxnWb2e/u992bee7tCa00YFsffekFY+nUzFtjW0LrvjRXrCDIAaPLlW0nHL0SsZtVoaF98mLrx3pdhOqLtYPHChahZcYYO7KvPFxvRl5XPp1sN3adWiD1ZAqD6XYK1b/dvE5IWryTt2udLFedwc1+9kLp+vbbpoDh+6TklxBeAi9TL0taeWpdmZzQDry0AcO+jQ12RyohqqoYoo8RDwJrU+qXkjWtfi8Xxt58BdQuwQs9qC/afLwCw8tnQbqYAPsgxE1S6F3EAIXux2oQFKm0ihMsOF71dHYx+f3NND68ghCu1YIoePPQN1pGRABkJ6Bus96CutRZMydTl+TvuiRW1m3n0eDl0vRPcEysqdXn+jsQPsrHMquGeXEaY4Yk4wxWcY5V/9scqOMOVUFthatyTy8QyqwZ+kDURKoMWxNKr2EeqVKcTNOajqKoBgOE28U4tdQl5p5bwCw7BWquaZSzAPlwjlithJtp3pTImSqQRrb2Z8PHGigD4RZuNX6JYj6wj7O4TFLbCO/Mn/m8R+h6rYSUb3ekokRY6f/YukArN979jcW+V/S8g0eT/N3VN3kTqWbQ428m9/8k0P/1aIhF36PccEl6EhOcAUCrXKZXXWS3XKd2vc/TRBG9O5ELC17MmWubD2nKhUKZa26Ba2+D3P+4/MNCFwg59oWVeYhkzgN/JDR8deKBoD7Y+ljEjGZ0sosXVTvbc6RHirr2reNy1OXd6pJsQ+gqjk8VWFYmHrwBzW/n+uMPFiRwHB2I7ih8ciHFxIkd/3Omk5tCDV1t+2nNu5sxxpDFNx+huNhVT3/zMDz8usXC3ddaHBj1GHj/As08fwTS7Kt1HBTmyN29vdwAw+/wbwLVOJ3uAD1wi/dUH7Qei66PfyuRj4Ik9is+hglfbkbfR3cnZm7chlUWLdwmprtCohX4HUtlOcQjLYCu+fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII=",
				}
				images.EXPECT().CreateImageSignature(input).Return("public/images/signatures/signatures/signatures-example.png").Times(1)
			},
		},
		{
			name: "Create Image Signature Case 2: Failed Create Image Signature Base64 Invalid",
			input: models.AddSignature{
				Id:        "6380b5cbdc938c5fdf8e6bfe",
				Signature: "fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII=",
			},
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.AddSignature{
					Id:        "6380b5cbdc938c5fdf8e6bfe",
					Signature: "fzGJH2QRKvP3UNz8bWk1qMxjGTOMThZ3kvgLI5AzFfo379UAAAAASUVORK5CYII=",
				}
				images.EXPECT().CreateImageSignature(input).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)
			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			ouput := s.CreateImgSignature(tt.input)
			assert.Equal(t, tt.output, ouput)
		})
	}
}

func Test_service_CreateImgSignatureData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name      string
		input     models.AddSignature
		name_sign string
		output    string
		test      func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			name: "Create Image Signature Data Case 1: Success Create Image Signature Data",
			input: models.AddSignature{
				Id: "6380b5cbdc938c5fdf8e6bfe",
			},
			name_sign: "Rizqi Wijaya",
			output:    "public/images/signatures/signatures_data/signaturesdata-6380b5cbdc938c5fdf8e6bfe.png",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.AddSignature{
					Id: "6380b5cbdc938c5fdf8e6bfe",
				}
				images.EXPECT().CreateImgSignatureData(input, "Rizqi Wijaya", "detail_data.ttf").Return("public/images/signatures/signatures_data/signaturesdata-6380b5cbdc938c5fdf8e6bfe.png").Times(1)
			},
		},
		{
			name: "Create Image Signature Data Case 2: Failed, Font Not Found",
			input: models.AddSignature{
				Id: "6380b5cbdc938c5fdf8e6bfe",
			},
			name_sign: "Rizqi Wijaya",
			output:    "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.AddSignature{
					Id: "6380b5cbdc938c5fdf8e6bfe",
				}
				images.EXPECT().CreateImgSignatureData(input, "Rizqi Wijaya", "detail_data.ttf").Return("").Times(1)
			},
		},
		{
			name: "Create Image Signature Data Case 3: Failed, Images Signatures Not Found",
			input: models.AddSignature{
				Id: "6380b5cbdc938cs",
			},
			name_sign: "Rizqi Wijaya",
			output:    "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.AddSignature{
					Id: "6380b5cbdc938cs",
				}
				images.EXPECT().CreateImgSignatureData(input, "Rizqi Wijaya", "detail_data.ttf").Return("").Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			ouput := s.CreateImgSignatureData(tt.input, tt.name_sign)
			assert.Equal(t, tt.output, ouput)
		})
	}
}

func Test_service_CreateLatinSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name   string
		input  modelsUser.User
		id     string
		output string
		test   func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			name: "Create Latin Signatures Case 1: Success Create Latin Signatures",
			input: modelsUser.User{
				Name: "Rizqi Wijaya",
			},
			id:     "6380b5cbdc938c5fdf8e6bfe",
			output: "public/images/signatures/latin/latin-6380b5cbdc938c5fdf8e6bfe.png",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := modelsUser.User{
					Name: "Rizqi Wijaya",
				}
				images.EXPECT().CreateLatinSignatures(input, "6380b5cbdc938c5fdf8e6bfe", "latin.ttf").Return("public/images/signatures/latin/latin-6380b5cbdc938c5fdf8e6bfe.png").Times(1)
			},
		},
		{
			name: "Create Latin Signatures Case 2: Failed, Font Not Found",
			input: modelsUser.User{
				Name: "Rizqi Wijaya",
			},
			id:     "6380b5cbdc938c5fdf8e6bfe",
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := modelsUser.User{
					Name: "Rizqi Wijaya",
				}
				images.EXPECT().CreateLatinSignatures(input, "6380b5cbdc938c5fdf8e6bfe", "latin.ttf").Return("").Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			ouput := s.CreateLatinSignatures(tt.input, tt.id)
			assert.Equal(t, tt.output, ouput)
		})
	}
}

func Test_service_CreateLatinSignaturesData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name   string
		input  modelsUser.User
		id     string
		latin  string
		output string
		test   func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			name: "Create Latin Signatures Case 1: Success Create Latin Signatures",
			input: modelsUser.User{
				Name: "Rizqi Wijaya",
			},
			id:     "6380b5cbdc938c5fdf8e6bfe",
			latin:  "public/images/signatures/latin/latin-6380b5cbdc938c5fdf8e6bfe.png",
			output: "public/images/signatures/latin_data/latindata-6380b5cbdc938c5fdf8e6bfe.png",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := modelsUser.User{
					Name: "Rizqi Wijaya",
				}
				latin := "public/images/signatures/latin/latin-6380b5cbdc938c5fdf8e6bfe.png"
				images.EXPECT().CreateLatinSignaturesData(input, latin, "6380b5cbdc938c5fdf8e6bfe", "detail_data.ttf").Return("public/images/signatures/latin_data/latindata-6380b5cbdc938c5fdf8e6bfe.png").Times(1)
			},
		},
		{
			name: "Create Latin Signatures Case 2: Failed Images Not Found",
			input: modelsUser.User{
				Name: "Rizqi Wijaya",
			},
			id:     "6380b5cbdc938c5fdf8e6bfe",
			latin:  "public/images/signatures/latin/latin-e.png",
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := modelsUser.User{
					Name: "Rizqi Wijaya",
				}
				latin := "public/images/signatures/latin/latin-e.png"
				images.EXPECT().CreateLatinSignaturesData(input, latin, "6380b5cbdc938c5fdf8e6bfe", "detail_data.ttf").Return("").Times(1)
			},
		},
		{
			name: "Create Latin Signatures Case 3: Failed Font Not Found",
			input: modelsUser.User{
				Name: "Rizqi Wijaya",
			},
			id:     "6380b5cbdc938c5fdf8e6bfe",
			latin:  "public/images/signatures/latin/latin-6380b5cbdc938c5fdf8e6bfe.png",
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := modelsUser.User{
					Name: "Rizqi Wijaya",
				}
				latin := "public/images/signatures/latin/latin-6380b5cbdc938c5fdf8e6bfe.png"
				images.EXPECT().CreateLatinSignaturesData(input, latin, "6380b5cbdc938c5fdf8e6bfe", "detail_data.ttf").Return("").Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			ouput := s.CreateLatinSignaturesData(tt.input, tt.latin, tt.id)
			assert.Equal(t, tt.output, ouput)
		})
	}
}

func Test_service_DefaultSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name  string
		input modelsUser.User
		id    string
		err   error
		test  func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			name: "Default Signatures Case 1: Success Insert Default Signatures Users",
			input: modelsUser.User{
				Idsignature: "rizwijaya",
			},
			id:  "6380b5cbdc938c5fdf8e6bfe",
			err: nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := modelsUser.User{
					Idsignature: "rizwijaya",
				}
				repo.EXPECT().DefaultSignatures(input, "6380b5cbdc938c5fdf8e6bfe").Return(nil).Times(1)
			},
		},
		{
			name: "Default Signatures Case 2: Error Failed Insert Default Signatures Users",
			input: modelsUser.User{
				Idsignature: "rizwijaya",
			},
			id:  "6380b5cbd",
			err: errors.New("Failed Insert Default Signatures Users to Database"),
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := modelsUser.User{
					Idsignature: "rizwijaya",
				}
				repo.EXPECT().DefaultSignatures(input, "6380b5cbd").Return(errors.New("Failed Insert Default Signatures Users to Database")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			err := s.DefaultSignatures(tt.input, tt.id)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_service_UpdateMySignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name          string
		signature     string
		signaturedata string
		sign          string
		err           error
		test          func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			name:          "Update My Signatures Case 1: Success Update My Signatures Users",
			signature:     "signatures-6380b5cbdc938c5fdf8e6bfe.png",
			signaturedata: "signaturesdata-6380b5cbdc938c5fdf8e6bfe.png",
			sign:          "6380b5cbdc938c5fdf8e6bfe",
			err:           nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				repo.EXPECT().UpdateMySignatures("signatures-6380b5cbdc938c5fdf8e6bfe.png", "signaturesdata-6380b5cbdc938c5fdf8e6bfe.png", "6380b5cbdc938c5fdf8e6bfe").Return(nil).Times(1)
			},
		},
		{
			name:          "Update My Signatures Case 2: Error Failed Update My Signatures Users",
			signature:     "signatures-e.png",
			signaturedata: "signaturesdata-e.png",
			sign:          "6380b5cbdc938c5fdf8e6bfe",
			err:           errors.New("Error Failed Update My Signatures Users to Database"),
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				repo.EXPECT().UpdateMySignatures("signatures-e.png", "signaturesdata-e.png", "6380b5cbdc938c5fdf8e6bfe").Return(errors.New("Error Failed Update My Signatures Users to Database")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			err := s.UpdateMySignatures(tt.signature, tt.signaturedata, tt.sign)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_service_GetMySignature(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	id, err := primitive.ObjectIDFromHex("6380b5cbdc938c5fdf8e6bfe")
	assert.NoError(t, err)
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)

	test := []struct {
		nameTest string
		sign     string
		id       string
		name     string
		output   models.MySignatures
		test     func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			nameTest: "Update My Signatures Case 1: Success Update My Signatures Users",
			sign:     "rizwijaya",
			id:       "6380b5cbdc938c5fdf8e6bfe",
			name:     "Rizqi Wijaya",
			output: models.MySignatures{
				Id:                 "6380b5cbdc938c5fdf8e6bfe",
				Name:               "Rizqi Wijaya",
				User_id:            "6380b5cbdc938c5fdf8e6bfe",
				Signature:          "signatures/sign-6380b5cbdc938c5fdf8e6bfe.png",
				Signature_id:       "sign-6380b5cbdc938c5fdf8e6bfe",
				Signature_data:     "signatures_data/signatures_data-6380b5cbdc938c5fdf8e6bfe.png",
				Signature_data_id:  "sign_data-6380b5cbdc938c5fdf8e6bfe",
				Latin:              "latin/latin-6380b5cbdc938c5fdf8e6bfe.png",
				Latin_id:           "latin-6380b5cbdc938c5fdf8e6bfe",
				Latin_data:         "latin_data/latin_data-6380b5cbdc938c5fdf8e6bfe.png",
				Latin_data_id:      "latin_data-6380b5cbdc938c5fdf8e6bfe",
				Signature_selected: "latin",
				Date_update:        "27 Nopember 2022 | 18:30 WIB",
				Date_created:       "27 Nopember 2022 | 18:30 WIB",
			},
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				sign := "rizwijaya"
				repo.EXPECT().GetMySignature(sign).Return(models.Signatures{
					Id:                 id,
					User:               "rizwijaya",
					Signature:          "sign-6380b5cbdc938c5fdf8e6bfe.png",
					Signature_data:     "signatures_data-6380b5cbdc938c5fdf8e6bfe.png",
					Latin:              "latin-6380b5cbdc938c5fdf8e6bfe.png",
					Latin_data:         "latin_data-6380b5cbdc938c5fdf8e6bfe.png",
					Signature_selected: "latin",
					Date_update:        times,
					Date_created:       times,
				}, nil).Times(1)
			},
		},
		{
			nameTest: "Update My Signatures Case 2: Error Failed Update My Signatures Users",
			sign:     "rizwijaya",
			id:       "6380b5cbdc938c5fdf8e6bfe",
			name:     "Rizqi Wijaya",
			output: models.MySignatures{
				Id:                 "6380b5cbdc938c5fdf8e6bfe",
				Name:               "Rizqi Wijaya",
				User_id:            "6380b5cbdc938c5fdf8e6bfe",
				Signature:          "signatures/sign-6380b5cbdc938c5fdf8e6bfe.png",
				Signature_id:       "sign-6380b5cbdc938c5fdf8e6bfe",
				Signature_data:     "signatures_data/signatures_data-6380b5cbdc938c5fdf8e6bfe.png",
				Signature_data_id:  "sign_data-6380b5cbdc938c5fdf8e6bfe",
				Latin:              "latin/latin-6380b5cbdc938c5fdf8e6bfe.png",
				Latin_id:           "latin-6380b5cbdc938c5fdf8e6bfe",
				Latin_data:         "latin_data/latin_data-6380b5cbdc938c5fdf8e6bfe.png",
				Latin_data_id:      "latin_data-6380b5cbdc938c5fdf8e6bfe",
				Signature_selected: "latin",
				Date_update:        "27 Nopember 2022 | 18:30 WIB",
				Date_created:       "27 Nopember 2022 | 18:30 WIB",
			},
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				sign := "rizwijaya"
				repo.EXPECT().GetMySignature(sign).Return(models.Signatures{
					Id:                 id,
					User:               "rizwijaya",
					Signature:          "sign-6380b5cbdc938c5fdf8e6bfe.png",
					Signature_data:     "signatures_data-6380b5cbdc938c5fdf8e6bfe.png",
					Latin:              "latin-6380b5cbdc938c5fdf8e6bfe.png",
					Latin_data:         "latin_data-6380b5cbdc938c5fdf8e6bfe.png",
					Signature_selected: "latin",
					Date_update:        times,
					Date_created:       times,
				}, errors.New("Error Failed Update My Signatures Users")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			output := s.GetMySignature(tt.sign, tt.id, tt.name)
			assert.Equal(t, tt.output, output)
		})
	}
}

func Test_service_ChangeSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		nameTest    string
		sign_type   string
		idsignature string
		err         error
		test        func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			nameTest:    "Change Signatures Case 1: Success Change Signatures User to Latin",
			sign_type:   "latin",
			idsignature: "rizwijaya",
			err:         nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				repo.EXPECT().ChangeSignature("latin", "rizwijaya").Return(nil).Times(1)
			},
		},
		{
			nameTest:    "Change Signatures Case 2: Success Change Signatures User to Latin Data",
			sign_type:   "latin_data",
			idsignature: "rizwijaya",
			err:         nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				repo.EXPECT().ChangeSignature("latin_data", "rizwijaya").Return(nil).Times(1)
			},
		},
		{
			nameTest:    "Change Signatures Case 3: Success Change Signatures User to Signature",
			sign_type:   "signature",
			idsignature: "rizwijaya",
			err:         nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				repo.EXPECT().ChangeSignature("signature", "rizwijaya").Return(nil).Times(1)
			},
		},
		{
			nameTest:    "Change Signatures Case 4: Success Change Signatures User to Signature Data",
			sign_type:   "signature_data",
			idsignature: "rizwijaya",
			err:         nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				repo.EXPECT().ChangeSignature("signature_data", "rizwijaya").Return(nil).Times(1)
			},
		},
		{
			nameTest:    "Change Signatures Case 5: Error Failed Change Signatures User",
			sign_type:   "signature",
			idsignature: "rizwijaya",
			err:         errors.New("Error Failed Change Signatures User"),
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				repo.EXPECT().ChangeSignature("signature", "rizwijaya").Return(errors.New("Error Failed Change Signatures User")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			err := s.ChangeSignatures(tt.sign_type, tt.idsignature)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_service_ResizeImages(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		nameTest string
		mysign   models.MySignatures
		signDocs models.SignDocuments
		output   string
		test     func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			nameTest: "Resize Images Case 1: Success Resize Images Signatures",
			mysign: models.MySignatures{
				Id:                 "6380b5cbdc938c5fdf8e6bfe",
				Name:               "Rizqi Wijaya",
				User_id:            "6380b5cbdc938c5fdf8e6bfe",
				Signature:          "signatures/sign-6380b5cbdc938c5fdf8e6bfe.png",
				Signature_id:       "sign-6380b5cbdc938c5fdf8e6bfe",
				Signature_data:     "signatures_data/signatures_data-6380b5cbdc938c5fdf8e6bfe.png",
				Signature_data_id:  "sign_data-6380b5cbdc938c5fdf8e6bfe",
				Latin:              "latin/latin-6380b5cbdc938c5fdf8e6bfe.png",
				Latin_id:           "latin-6380b5cbdc938c5fdf8e6bfe",
				Latin_data:         "latin_data/latin_data-6380b5cbdc938c5fdf8e6bfe.png",
				Latin_data_id:      "latin_data-6380b5cbdc938c5fdf8e6bfe",
				Signature_selected: "latin",
				Date_update:        "27 Nopember 2022 | 18:30 WIB",
				Date_created:       "27 Nopember 2022 | 18:30 WIB",
			},
			signDocs: models.SignDocuments{
				Height: 200.3,
				Width:  143.6,
			},
			output: "./public/temp/sizes-latin.png",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				mysign := models.MySignatures{
					Id:                 "6380b5cbdc938c5fdf8e6bfe",
					Name:               "Rizqi Wijaya",
					User_id:            "6380b5cbdc938c5fdf8e6bfe",
					Signature:          "signatures/sign-6380b5cbdc938c5fdf8e6bfe.png",
					Signature_id:       "sign-6380b5cbdc938c5fdf8e6bfe",
					Signature_data:     "signatures_data/signatures_data-6380b5cbdc938c5fdf8e6bfe.png",
					Signature_data_id:  "sign_data-6380b5cbdc938c5fdf8e6bfe",
					Latin:              "latin/latin-6380b5cbdc938c5fdf8e6bfe.png",
					Latin_id:           "latin-6380b5cbdc938c5fdf8e6bfe",
					Latin_data:         "latin_data/latin_data-6380b5cbdc938c5fdf8e6bfe.png",
					Latin_data_id:      "latin_data-6380b5cbdc938c5fdf8e6bfe",
					Signature_selected: "latin",
					Date_update:        "27 Nopember 2022 | 18:30 WIB",
					Date_created:       "27 Nopember 2022 | 18:30 WIB",
				}

				signDocs := models.SignDocuments{
					Height: 200.3,
					Width:  143.6,
				}

				images.EXPECT().ResizeImages(mysign, signDocs).Return("./public/temp/sizes-latin.png").Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			output := s.ResizeImages(tt.mysign, tt.signDocs)
			assert.Equal(t, tt.output, output)
		})
	}
}

func Test_service_SignDocuments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		nameTest string
		imgpath  string
		signDocs models.SignDocuments
		output   string
		test     func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			nameTest: "Sign Documents Case 1: Success Sign Documents",
			imgpath:  "./public/temp/sizes-latin.png",
			signDocs: models.SignDocuments{
				Name:    "sample_test.pdf",
				X_coord: 350,
				Y_coord: 310,
				Height:  200.3,
				Width:   143.6,
			},
			output: "./public/temp/pdfsign/signed_sample_test.pdf",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				signDocs := models.SignDocuments{
					Name:    "sample_test.pdf",
					X_coord: 350,
					Y_coord: 310,
					Height:  200.3,
					Width:   143.6,
				}
				docs.EXPECT().SignDocuments("./public/temp/sizes-latin.png", signDocs).Return("./public/temp/pdfsign/signed_sample_test.pdf").Times(1)
			},
		},
		{
			nameTest: "Sign Documents Case 2: Failed Sign Documents because image path is empty",
			imgpath:  "",
			signDocs: models.SignDocuments{
				Name:    "sample_test.pdf",
				X_coord: 350,
				Y_coord: 310,
				Height:  200.3,
				Width:   143.6,
			},
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				signDocs := models.SignDocuments{
					Name:    "sample_test.pdf",
					X_coord: 350,
					Y_coord: 310,
					Height:  200.3,
					Width:   143.6,
				}
				docs.EXPECT().SignDocuments("", signDocs).Return("").Times(1)
			},
		},
		{
			nameTest: "Sign Documents Case 3: Failed Sign Documents because cannot read and write document pdf",
			imgpath:  "./public/temp/sizes-latin.png",
			signDocs: models.SignDocuments{
				Name:    "sample_test.pdf",
				X_coord: 350,
				Y_coord: 310,
				Height:  200.3,
				Width:   143.6,
			},
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				signDocs := models.SignDocuments{
					Name:    "sample_test.pdf",
					X_coord: 350,
					Y_coord: 310,
					Height:  200.3,
					Width:   143.6,
				}
				docs.EXPECT().SignDocuments("./public/temp/sizes-latin.png", signDocs).Return("").Times(1)
			},
		},
		{
			nameTest: "Sign Documents Case 4: Failed Sign Documents because cannot get pages from document pdf",
			imgpath:  "./public/temp/sizes-latin.png",
			signDocs: models.SignDocuments{
				Name:    "sample_test.pdf",
				X_coord: 350,
				Y_coord: 310,
				Height:  200.3,
				Width:   143.6,
			},
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				signDocs := models.SignDocuments{
					Name:    "sample_test.pdf",
					X_coord: 350,
					Y_coord: 310,
					Height:  200.3,
					Width:   143.6,
				}
				docs.EXPECT().SignDocuments("./public/temp/sizes-latin.png", signDocs).Return("").Times(1)
			},
		},
		{
			nameTest: "Sign Documents Case 5: Failed Sign Documents because cannot add images signatures to document pdf",
			imgpath:  "./public/temp/sizes-latin.png",
			signDocs: models.SignDocuments{
				Name:    "sample_test.pdf",
				X_coord: 350,
				Y_coord: 310,
				Height:  200.3,
				Width:   143.6,
			},
			output: "",
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				signDocs := models.SignDocuments{
					Name:    "sample_test.pdf",
					X_coord: 350,
					Y_coord: 310,
					Height:  200.3,
					Width:   143.6,
				}
				docs.EXPECT().SignDocuments("./public/temp/sizes-latin.png", signDocs).Return("").Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			if tt.test != nil {
				tt.test(repo, images, docs)
			}

			s := NewService(repo, images, docs)
			output := s.SignDocuments(tt.imgpath, tt.signDocs)
			assert.Equal(t, tt.output, output)
		})
	}
}

func Test_service_InvitePeople(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		nameTest string
		email    string
		signDocs models.SignDocuments
		users    modelsUser.User
		err      error
	}{
		{
			nameTest: "Sign Documents Case 1: Success Sign Documents",
			email:    "member@rizwijaya.com",
			signDocs: models.SignDocuments{
				Judul:         "Invite People Test",
				Creator_id:    "rizwijaya",
				Note:          "Note Invite People Test",
				Hash_original: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				Name:          "sample_test.pdf",
			},
			users: modelsUser.User{
				Name: "Rizqi Wijaya",
			},
			err: nil,
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			s := NewService(repo, images, docs)
			err := s.InvitePeople(tt.email, tt.signDocs, tt.users)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_service_GenerateHashDocument(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		nameTest string
		input    string
		output   string
	}{
		{
			nameTest: "Generate Hash Document Case 1: Success Generate Hash Document",
			input:    "./public/unit_testing/sample_test.pdf",
			output:   "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			nameTest: "Generate Hash Document Case 2: Error Failed Generate Hash Document PDF file not found",
			input:    "samp.pdf",
			output:   "",
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)

			s := NewService(repo, images, docs)
			output := s.GenerateHashDocument(tt.input)
			if tt.output == "" {
				assert.Equal(t, tt.output, output)
			} else {
				assert.NotEqual(t, tt.output, output)
			}
		})
	}
}

func Test_service_AddToBlockhain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)

	test := []struct {
		nameTest string
		input    models.SignDocuments
		err      error
		test     func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			nameTest: "Add To Blockchain Case 1: Success Add Data Signatures and Document To Blockchain",
			input: models.SignDocuments{
				Name:          "sample_test.pdf",
				SignPage:      1.0,
				X_coord:       1.3,
				Y_coord:       1.2,
				Height:        4.2,
				Width:         5.3,
				Invite_sts:    true,
				Email:         []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
				Note:          "Note Test Add To Blockchain",
				Judul:         "Judul Test Add To Blockchain",
				Mode:          "1",
				IPFS:          "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
				Hash:          "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				Hash_original: "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				Creator:       "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id:    "rizwijaya",
				Address:       []common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")},
				IdSignature:   []string{"signed_1", "signed2"},
			},
			err: nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.SignDocuments{
					Name:          "sample_test.pdf",
					SignPage:      1.0,
					X_coord:       1.3,
					Y_coord:       1.2,
					Height:        4.2,
					Width:         5.3,
					Invite_sts:    true,
					Email:         []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
					Note:          "Note Test Add To Blockchain",
					Judul:         "Judul Test Add To Blockchain",
					Mode:          "1",
					IPFS:          "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
					Hash:          "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Hash_original: "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Creator:       "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Creator_id:    "rizwijaya",
					Address:       []common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")},
					IdSignature:   []string{"signed_1", "signed2"},
				}
				timeSign := new(big.Int)
				timeFormat := times.Format("15040502012006")
				timeSign, _ = timeSign.SetString(timeFormat, 10)
				repo.EXPECT().AddToBlockhain(input, timeSign).Return(nil).Times(1)
			},
		},
		{
			nameTest: "Add To Blockchain Case 2: Error Failed Add Data Signatures and Document To Blockchain",
			input: models.SignDocuments{
				Name:       "sample_test.pdf",
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: true,
				Email:      []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
				Note:       "Note Test Add To Blockchain",
				Judul:      "Judul Test Add To Blockchain",
				Mode:       "1",
				IPFS:       "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
			},
			err: errors.New("Failed Insert Data to Blockchain"),
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.SignDocuments{
					Name:       "sample_test.pdf",
					SignPage:   1.0,
					X_coord:    1.3,
					Y_coord:    1.2,
					Height:     4.2,
					Width:      5.3,
					Invite_sts: true,
					Email:      []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
					Note:       "Note Test Add To Blockchain",
					Judul:      "Judul Test Add To Blockchain",
					Mode:       "1",
					IPFS:       "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
				}
				timeSign := new(big.Int)
				timeFormat := times.Format("15040502012006")
				timeSign, _ = timeSign.SetString(timeFormat, 10)
				repo.EXPECT().AddToBlockhain(input, timeSign).Return(errors.New("Failed Insert Data to Blockchain")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)
			if tt.test != nil {
				tt.test(repo, images, docs)
			}
			s := NewService(repo, images, docs)
			err := s.AddToBlockhain(tt.input)
			assert.Equal(t, tt.err, err)
		})
	}
}

func Test_service_AddUserDocs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		nameTest string
		input    models.SignDocuments
		err      error
		test     func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			nameTest: "Add User and Documents Case 1: Success Add Data User and Documents",
			input: models.SignDocuments{
				Name:          "sample_test.pdf",
				SignPage:      1.0,
				X_coord:       1.3,
				Y_coord:       1.2,
				Height:        4.2,
				Width:         5.3,
				Invite_sts:    true,
				Email:         []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
				Note:          "Note Test",
				Judul:         "Judul Test",
				Mode:          "1",
				IPFS:          "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
				Hash:          "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				Hash_original: "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				Creator:       "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id:    "rizwijaya",
				Address:       []common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")},
				IdSignature:   []string{"signed_1", "signed2"},
			},
			err: nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.SignDocuments{
					Name:          "sample_test.pdf",
					SignPage:      1.0,
					X_coord:       1.3,
					Y_coord:       1.2,
					Height:        4.2,
					Width:         5.3,
					Invite_sts:    true,
					Email:         []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
					Note:          "Note Test",
					Judul:         "Judul Test",
					Mode:          "1",
					IPFS:          "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
					Hash:          "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Hash_original: "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Creator:       "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Creator_id:    "rizwijaya",
					Address:       []common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")},
					IdSignature:   []string{"signed_1", "signed2"},
				}
				repo.EXPECT().AddUserDocs(input).Return(nil).Times(1)
			},
		},
		{
			nameTest: "Add User and Documents Case 2: Error Failed Add Data User and Documents",
			input: models.SignDocuments{
				Name:       "sample_test.pdf",
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: true,
				Email:      []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
				Note:       "Note Test",
				Judul:      "Judul Test",
				Mode:       "1",
				IPFS:       "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
			},
			err: errors.New("Failed Add User and Documents Data"),
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				input := models.SignDocuments{
					Name:       "sample_test.pdf",
					SignPage:   1.0,
					X_coord:    1.3,
					Y_coord:    1.2,
					Height:     4.2,
					Width:      5.3,
					Invite_sts: true,
					Email:      []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
					Note:       "Note Test",
					Judul:      "Judul Test",
					Mode:       "1",
					IPFS:       "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
				}
				repo.EXPECT().AddUserDocs(input).Return(errors.New("Failed Add User and Documents Data")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)
			if tt.test != nil {
				tt.test(repo, images, docs)
			}
			s := NewService(repo, images, docs)
			err := s.AddUserDocs(tt.input)
			assert.Equal(t, tt.err, err)
		})
	}

}

func Test_service_DocumentSigned(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location)

	test := []struct {
		nameTest string
		sign     models.SignDocs
		err      error
		test     func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments)
	}{
		{
			nameTest: "Document Signed Case 1: Success Sign Document in Blockchain",
			sign: models.SignDocs{
				Hash_original: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				Creator:       "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Hash:          "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				IPFS:          "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
			},
			err: nil,
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				sign := models.SignDocs{
					Hash_original: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Creator:       "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Hash:          "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					IPFS:          "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
				}
				timeSign := new(big.Int)
				timeFormat := times.Format("15040502012006")
				timeSign, _ = timeSign.SetString(timeFormat, 10)
				repo.EXPECT().DocumentSigned(sign, timeSign).Return(nil).Times(1)
			},
		},
		{
			nameTest: "Document Signed Case 2: Failed Sign Document Because No Hash Original Data",
			sign: models.SignDocs{
				Creator: "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Hash:    "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				IPFS:    "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
			},
			err: errors.New("Failed Sign Document in Blockchain"),
			test: func(repo *m_repo.MockRepository, images *m_images.MockImages, docs *m_docs.MockDocuments) {
				sign := models.SignDocs{
					Creator: "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Hash:    "u798sc537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					IPFS:    "d9sj84msl02ndm93d8df4d2u43soj3bdsds",
				}
				timeSign := new(big.Int)
				timeFormat := times.Format("15040502012006")
				timeSign, _ = timeSign.SetString(timeFormat, 10)
				repo.EXPECT().DocumentSigned(sign, timeSign).Return(errors.New("Failed Sign Document in Blockchain")).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)
			if tt.test != nil {
				tt.test(repo, images, docs)
			}
			s := NewService(repo, images, docs)
			err := s.DocumentSigned(tt.sign)
			assert.Equal(t, tt.err, err)
		})
	}

}
