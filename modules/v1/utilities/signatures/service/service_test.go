package service

import (
	"e-signature/modules/v1/utilities/signatures/models"
	m_repo "e-signature/modules/v1/utilities/signatures/repository/mock"
	modelsUser "e-signature/modules/v1/utilities/user/models"
	m_docs "e-signature/pkg/document/mock"
	m_images "e-signature/pkg/images/mock"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tkuchiki/faketime"
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
