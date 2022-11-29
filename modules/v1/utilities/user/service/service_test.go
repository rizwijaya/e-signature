package service

import (
	"e-signature/modules/v1/utilities/user/models"
	m_repo "e-signature/modules/v1/utilities/user/repository/mock"
	"errors"
	"os"
	"testing"
	"time"

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
		serv := NewService(nil)
		assert.NotNil(t, serv)
	})
}

func Test_service_Login(t *testing.T) {
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
		name     string
		input    models.LoginInput
		repoTest func(repo *m_repo.MockRepository)
	}{
		{
			name: "Login Service Case 1: Failed Login User Not Found",
			input: models.LoginInput{
				IdSignature: "rizwijaya",
				Password:    "123456",
			},
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("rizwijaya").Return(models.ProfileDB{}, nil).Times(1)
			},
		},
		{
			name: "Login Service Case 2: Failed Login Password Not Valid",
			input: models.LoginInput{
				IdSignature: "rizwijaya",
				Password:    "123",
			},
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("rizwijaya").Return(models.ProfileDB{
					Id:            primitive.NewObjectID(),
					Idsignature:   "rizwijaya",
					Name:          "Rizqi Wijaya",
					Email:         "smartsign@rizwijaya.com",
					Phone:         "081234567890",
					Identity_card: "a1ddfs2s3fes4s5s2aas3sdsdaad167890",
					Password:      "123456",
					PublicKey:     "sjn729nad6a9804jd7tnfs8bqvi3iaewc9y80c",
					Role_id:       2,
					Date_created:  times,
				}, nil).Times(1)
			},
		},
		{
			name: "Login Service Case 3: Login Success",
			input: models.LoginInput{
				IdSignature: "rizwijaya",
				Password:    "rizwijaya123",
			},
			repoTest: func(repo *m_repo.MockRepository) {
				repo.EXPECT().CheckUserExist("rizwijaya").Return(models.ProfileDB{
					Id:            primitive.NewObjectID(),
					Idsignature:   "rizwijaya",
					Name:          "Rizqi Wijaya",
					Email:         "smartsign@rizwijaya.com",
					Phone:         "081234567890",
					Identity_card: "a1ddfs2s3fes4s5s2aas3sdsdaad167890",
					Password:      "$2a$04$mL8CVXcMKPOINfUgBszbgupxCC9lj0eqPvnnz/iNng/CisnSjdYdu",
					PublicKey:     "sjn729nad6a9804jd7tnfs8bqvi3iaewc9y80c",
					Role_id:       2,
					Date_created:  times,
				}, nil).Times(1)
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			tt.repoTest(repo)
			s := NewService(repo)
			_, err := s.Login(tt.input)
			if err == errors.New("user not found") {
				assert.NoError(t, err)
			}
		})
	}
}

func Test_service_UploadIPFS(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Upload File IPFS Service Case 1: Upload File in IPFS", func(t *testing.T) {
		repo := m_repo.NewMockRepository(ctrl)
		s := &service{
			repository: repo,
		}
		cidr, _ := s.UploadIPFS("")
		assert.Equal(t, cidr, "")
	})
}
