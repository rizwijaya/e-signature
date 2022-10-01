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
	//Prepare data for image signature and encryption
	pass := fmt.Sprintf("%v", sessions.Get("passph"))
	publickey := fmt.Sprintf("%v", sessions.Get("public_key"))
	name := fmt.Sprintf("%v", sessions.Get("name"))
	passph := string(h.serviceUser.Decrypt([]byte(pass), publickey))
	//Create base64 to image
	image := h.signaturesService.CreateImgSignature(input)
	//create signature with data
	imageData := h.signaturesService.CreateImgSignatureData(input, name)
	h.serviceUser.EncryptFile(image, passph)
	h.serviceUser.EncryptFile(imageData, passph)
	//Update Database MySignatures

	//Return Response API
	response := api.APIRespon("Success Add Signatures", 200, "success", nil)
	c.JSON(200, response)
}
