package signatures

import (
	"bytes"
	m_serviceSignature "e-signature/modules/v1/utilities/signatures/service/mock"
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
	gin.SetMode(gin.TestMode)
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
	router.Static("/landing/css", "./public/assets/landing/css")
	router.Static("/landing/js", "./public/assets/landing/js")
	router.Static("/form/vendor", "./public/assets/form/vendor")
	router.Static("/form/css", "./public/assets/form/css")
	router.Static("/form/js", "./public/assets/form/js")
	router.Static("/form/img", "./public/assets/form/img")
	router.Static("/signatures", "./public/images/signatures")
	router.Static("/file/documents", "./public/temp/pdfsign")
	router.HTMLRender = html.Render("./public/templates")
	return router
}

func Test_signaturesHandler_AddSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceSignature := m_serviceSignature.NewMockService(ctrl)
	serviceUser := m_serviceUser.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	tests := []struct {
		name        string
		statusCode  int
		request     string
		response    string
		serviceTest func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService)
	}{
		{
			name:       "Test Add Signatures Success",
			request:    `{"unique":"63322e432d405a140eb354e9","signature":"base64-pngimagefasnflanflasda"}`,
			statusCode: http.StatusOK,
			response:   `{"meta":{"message":"Success Add Signatures","code":200,"status":"success"},"data":null}`,
			serviceTest: func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService) {
				serviceSignature.EXPECT().CreateImgSignature(gomock.Any()).Return("public/images/signatures/signatures/signatures-63322e432d405a140eb354e9.png")
				serviceSignature.EXPECT().CreateImgSignatureData(gomock.Any(), gomock.Any()).Return("public/images/signatures/signatures_data/signaturesdata-63322e432d405a140eb354e9.png")
				serviceSignature.EXPECT().UpdateMySignatures(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
				serviceUser.EXPECT().Logging("Menambakan tanda tangan baru", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
		{
			name:       "Test Add Signatures, Input Not Valid",
			request:    `{"unique":"63322e432d405a140eb354e9","signature":"base64-pngimagefasnflanflasda"`,
			statusCode: 300,
			response:   `{"meta":{"message":"unexpected EOF","code":300,"status":"error"},"data":null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesHandler{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}

			//Testing Services Functions
			if tt.serviceTest != nil {
				tt.serviceTest(serviceSignature, serviceUser)
			}
			got := w.AddSignatures

			router := NewRouter()
			router.POST("/add-signatures", got)
			//Testing Handler Functions
			req, err := http.NewRequest("POST", "/add-signatures", bytes.NewReader([]byte(tt.request)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "text/plain")
			req.Header.Set("Cookie", cookies)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)
			responseData, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.statusCode, response.Code)
			assert.Equal(t, string(responseData), tt.response)
		})
	}
}

func Test_signaturesHandler_ChangeSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceSignature := m_serviceSignature.NewMockService(ctrl)
	serviceUser := m_serviceUser.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	tests := []struct {
		name        string
		statusCode  int
		sign_type   string
		sign_now    string
		pages       string
		serviceTest func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService)
	}{
		{
			name:       "Change Signatures to signature success",
			statusCode: 302,
			sign_type:  "signature",
			sign_now:   "<a href=\"/my-signatures\">Found</a>.\n\n",
			pages:      "/my-signatures",
			serviceTest: func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService) {
				serviceSignature.EXPECT().ChangeSignatures("signature", gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengganti tanda tangan ke signature", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
		{
			name:       "Change Signatures to signature data success",
			statusCode: 302,
			sign_type:  "signature_data",
			sign_now:   "<a href=\"/my-signatures\">Found</a>.\n\n",
			pages:      "/my-signatures",
			serviceTest: func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService) {
				serviceSignature.EXPECT().ChangeSignatures("signature_data", gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengganti tanda tangan ke signature_data", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
		{
			name:       "Change Signatures to latin success",
			statusCode: 302,
			sign_type:  "latin",
			sign_now:   "<a href=\"/my-signatures\">Found</a>.\n\n",
			pages:      "/my-signatures",
			serviceTest: func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService) {
				serviceSignature.EXPECT().ChangeSignatures("latin", gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengganti tanda tangan ke latin", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
		{
			name:       "Change Signatures to latin data success",
			statusCode: 302,
			sign_type:  "latin_data",
			sign_now:   "<a href=\"/my-signatures\">Found</a>.\n\n",
			pages:      "/my-signatures",
			serviceTest: func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService) {
				serviceSignature.EXPECT().ChangeSignatures("latin_data", gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengganti tanda tangan ke latin_data", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
		{
			name:       "Change Signatures, random sign type",
			statusCode: 302,
			sign_type:  "signature-nothing",
			sign_now:   "<a href=\"/my-signatures\">Found</a>.\n\n",
			pages:      "/my-signatures",
			serviceTest: func(signatureService *m_serviceSignature.MockService, signatureUser *m_serviceUser.MockService) {
				serviceSignature.EXPECT().ChangeSignatures("latin", gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengganti tanda tangan ke signature-nothing", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//Testing Services Functions
			if tt.serviceTest != nil {
				tt.serviceTest(serviceSignature, serviceUser)
			}
			w := &signaturesHandler{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			got := w.ChangeSignatures
			router := NewRouter()
			router.GET("/change-signatures/:sign_type", got)
			//Testing Handler Functions
			req, err := http.NewRequest("GET", "/change-signatures/"+tt.sign_type, nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)
			responseData, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)

			assert.Equal(t, tt.statusCode, response.Code)
			assert.Equal(t, string(responseData), tt.sign_now)

			location, err := response.Result().Location()
			assert.NoError(t, err)
			assert.Equal(t, location.Path, tt.pages)
		})
	}
}
