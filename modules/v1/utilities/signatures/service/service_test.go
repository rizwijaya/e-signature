package service

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
