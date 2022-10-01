package signatures

import (
	"e-signature/modules/v1/utilities/signatures/models"
	api "e-signature/pkg/api_response"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (h *signaturesHandler) AddSignatures(c *gin.Context) {
	sessions := sessions.Default(c)
	var input models.AddSignature
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := api.APIRespon(err.Error(), 300, "error", nil)
		c.JSON(215, response)
		return
	}

	sign := fmt.Sprintf("%v", sessions.Get("sign"))
	h.signaturesService.CreateImgSignature(input)
	h.signaturesService.CreateImgSignatureData(input, sign)
	//Update Database MySignatures
	err = h.signaturesService.UpdateMySignatures(fmt.Sprintf("signatures-%s.png", input.Id), fmt.Sprintf("signaturesdata-%s.png", input.Id), sign)
	if err != nil {
		response := api.APIRespon(err.Error(), 300, "error", nil)
		c.JSON(215, response)
		return
	}
	//Return Response API
	response := api.APIRespon("Success Add Signatures", 200, "success", nil)
	c.JSON(200, response)
}
