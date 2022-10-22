package signatures

import (
	"bytes"
	m_service "e-signature/modules/v1/utilities/signatures/service/mock"
	m_serviceUser "e-signature/modules/v1/utilities/user/service/mock"
	"e-signature/pkg/html"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	os.Chdir("../../../../../")
}

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	cookieStore := cookie.NewStore([]byte("JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu"))
	router.Use(sessions.Sessions("smartsign", cookieStore))

	router.Static("/landing/assets", "./public/assets/landing")
	router.Static("/landing/vendor", "./public/assets/landing/vendor")
	router.Static("/landing/swiper", "./public/assets/landing/vendor/swiper")
	router.Static("/landing/purecounter", "./public/assets/landing/vendor/purecounter")
	router.Static("/landing/img", "./public/assets/landing/img")
	router.Static("/landing/js", "./public/assets/landing/js")
	router.Static("/form/vendor", "./public/assets/form/vendor")
	router.Static("/form/css", "./public/assets/form/css")
	router.Static("/form/js", "./public/assets/form/js")
	router.Static("/form/img", "./public/assets/form/img")
	router.Static("/signatures", "./public/images/signatures")
	router.HTMLRender = html.Render("./public/templates")
	return router
}

func Test_signaturesHandler_AddSignatures(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	service := m_service.NewMockService(ctrl)
	serviceUser := m_serviceUser.NewMockService(ctrl)

	mock := &signaturesHandler{
		signaturesService: service,
		serviceUser:       serviceUser,
	}
	router := NewRouter()
	router.POST("/signatures", mock.AddSignatures)

	tests := []struct {
		name        string
		statusCode  int
		request     string
		response    string
		serviceTest func(signatureService *m_service.MockService, signatureUser *m_serviceUser.MockService)
	}{
		{
			name:       "Success Add Signatures",
			request:    `{"unique":"63322e432d405a140eb354e9","signature":"base64-pngimagefasnflanflasda"}`,
			statusCode: http.StatusOK,
			response:   `{"meta":{"message":"Success Add Signatures","code":200,"status":"success"},"data":null}`,
			serviceTest: func(signatureService *m_service.MockService, signatureUser *m_serviceUser.MockService) {
				service.EXPECT().CreateImgSignature(gomock.Any()).Return("public/images/signatures/signatures/signatures-63322e432d405a140eb354e9.png")
				service.EXPECT().CreateImgSignatureData(gomock.Any(), gomock.Any()).Return("public/images/signatures/signatures_data/signaturesdata-63322e432d405a140eb354e9.png")
				service.EXPECT().UpdateMySignatures(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			},
		},
		{
			name:       "Request Format Failed",
			request:    `{"unique":"63322e432d405a140eb354e9","signature":"base64-pngimagefasnflanflasda"`,
			statusCode: 300,
			response:   `{"meta":{"message":"unexpected EOF","code":300,"status":"error"},"data":null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Testing Services Functions
			if tt.serviceTest != nil {
				tt.serviceTest(service, serviceUser)
			}
			//Testing Handler Functions
			req, err := http.NewRequest("POST", "/signatures", bytes.NewReader([]byte(tt.request)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "text/plain")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)
			responseData, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.statusCode, response.Code)
			assert.Equal(t, string(responseData), tt.response)
		})
	}
}
