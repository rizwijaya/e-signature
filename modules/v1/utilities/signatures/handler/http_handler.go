package signatures

import (
	"e-signature/modules/v1/utilities/signatures/models"
	api "e-signature/pkg/api_response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *signaturesHandler) AddSignatures(c *gin.Context) {
	var input models.AddSignature
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := api.APIRespon(err.Error(), 215, "error", nil)
		c.JSON(215, response)
		return
	}
	fmt.Println(input.Signature)
}
