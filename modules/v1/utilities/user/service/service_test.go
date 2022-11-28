package service

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
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
func Test_service_ConnectIPFS(t *testing.T) {
	//server := httptest.NewServer(http.HandlerFunc())
	defer gock.Off()

	gock.New("http://localhost:5001").
		MatchHeader("Accept", "application/json").
		Get("/").
		Reply(200).
		JSON(map[string]string{"value": "fixed"})

	t.Run("Case 1: Success Connect IPFS", func(t *testing.T) {
		w := &service{
			repository: nil,
		}
		sh := w.ConnectIPFS()
		assert.NotNil(t, sh)
		t.Log(sh)
	})
}
