package signatures

import (
	"bytes"
	"e-signature/modules/v1/utilities/signatures/models"
	m_serviceSignature "e-signature/modules/v1/utilities/signatures/service/mock"
	modelUser "e-signature/modules/v1/utilities/user/models"
	m_serviceUser "e-signature/modules/v1/utilities/user/service/mock"
	"e-signature/pkg/html"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tkuchiki/faketime"
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

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	tests := []struct {
		name        string
		statusCode  int
		request     string
		response    string
		serviceTest func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService)
	}{
		{
			name:       "Test Add Signatures Success",
			request:    `{"unique":"63322e432d405a140eb354e9","signature":"base64-pngimagefasnflanflasda"}`,
			statusCode: http.StatusOK,
			response:   `{"meta":{"message":"Success Add Signatures","code":200,"status":"success"},"data":null}`,
			serviceTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
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
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			serviceUser := m_serviceUser.NewMockService(ctrl)
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

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	tests := []struct {
		name        string
		statusCode  int
		sign_type   string
		sign_now    string
		pages       string
		serviceTest func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService)
	}{
		{
			name:       "Change Signatures to signature success",
			statusCode: 302,
			sign_type:  "signature",
			sign_now:   "<a href=\"/my-signatures\">Found</a>.\n\n",
			pages:      "/my-signatures",
			serviceTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
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
			serviceTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
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
			serviceTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
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
			serviceTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
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
			serviceTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceSignature.EXPECT().ChangeSignatures("latin", gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengganti tanda tangan ke signature-nothing", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			serviceUser := m_serviceUser.NewMockService(ctrl)
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

func CreateFilePDF(t *testing.T, w *multipart.Writer, filename string) (io.Writer, error) {
	t.Helper()
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", filename))
	h.Set("Content-Type", "application/pdf")
	return w.CreatePart(h)
}

func Test_signaturesHandler_SignDocuments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location).String()
	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	mysignature := models.MySignatures{
		Id:                 "1",
		Name:               "Rizqi Wijaya",
		User_id:            "rizwijaya",
		Signature:          "default.png",
		Signature_id:       "sign_type",
		Signature_data:     "default.png",
		Signature_data_id:  "sign_type",
		Latin:              "latin.png",
		Latin_id:           "sign_type",
		Latin_data:         "latin_data.png",
		Latin_data_id:      "sign_type",
		Signature_selected: "signature",
		Date_update:        times,
		Date_created:       times,
	}

	tests := []struct {
		name         string
		responseCode int
		docs         models.SignDocuments
		file         string
		pages        string
		serviceTest  func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService)
	}{
		{
			name:         "Test Sign Documents Input Invalid",
			responseCode: http.StatusFound,
			file:         "",
			docs: models.SignDocuments{
				SignPage:   1.0,
				X_coord:    1,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: false,
			},
			pages: "/sign-documents",
		},
		{
			name:         "Test Sign Documents Not File Request",
			responseCode: http.StatusFound,
			file:         "",
			docs: models.SignDocuments{
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: false,
			},
			pages: "/sign-documents",
		},
		{
			name:         "Test Sign Documents Not File PDF",
			responseCode: http.StatusFound,
			file:         "card_test.jpeg",
			docs: models.SignDocuments{
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: false,
			},
			pages: "/sign-documents",
		},
		{
			name:         "Test Sign Documents Failed to Input IPFS",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			docs: models.SignDocuments{
				Name:       "sample_test.pdf",
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: true,
				Email:      []string{"admin@smartsign.com"},
				Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id: "rizwijaya",
			},
			pages: "/sign-documents",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docs := models.SignDocuments{
					Name:       "sample_test.pdf",
					SignPage:   1.0,
					X_coord:    1.3,
					Y_coord:    1.2,
					Height:     4.2,
					Width:      5.3,
					Invite_sts: true,
					Email:      []string{"[61646d696e40736d6172747369676e2e636f6d]"},
				}

				path := "./public/temp/pdfsign/"
				serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				docs.Hash = docs.Hash_original
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("", errors.New("Failed to Input IPFS")).Times(1)
			},
		},
		{
			name:         "Test Sign Documents Mode Sign No Invite and Failed Add To Blockchain",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			docs: models.SignDocuments{
				Name:       "sample_test.pdf",
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: false,
				Email:      []string{""},
				Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id: "rizwijaya",
			},
			pages: "/sign-documents",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docs := models.SignDocuments{
					Name:       "sample_test.pdf",
					SignPage:   1.0,
					X_coord:    1.3,
					Y_coord:    1.2,
					Height:     4.2,
					Width:      5.3,
					Invite_sts: false,
					Email:      []string{"[]"},
				}

				path := "./public/temp/pdfsign/"
				serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Return(mysignature).Times(1)
				docs.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				docs.Hash = docs.Hash_original
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), "JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu").Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				docs.Address = append(docs.Address, common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.IdSignature = append(docs.IdSignature, docs.Creator_id)
				docs.Mode = "3"
				serviceSignature.EXPECT().AddToBlockhain(docs).Return(errors.New("Failed to Add to Blockchain")).Times(1)
			},
		},
		{
			name:         "Test Sign Documents Mode Sign with Invite and Failed to Add User Documents",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			docs: models.SignDocuments{
				Name:       "sample_test.pdf",
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: true,
				Email:      []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
				Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id: "rizwijaya",
			},
			pages: "/sign-documents",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docs := models.SignDocuments{
					Name:       "sample_test.pdf",
					SignPage:   1.0,
					X_coord:    1.3,
					Y_coord:    1.2,
					Height:     4.2,
					Width:      5.3,
					Invite_sts: true,
					Email:      []string{"[61646d696e4072697a77696a6179612e636f6d 736d6172747369676e4072697a77696a6179612e636f6d]"},
				}

				path := "./public/temp/pdfsign/"
				serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				docs.Hash = docs.Hash_original
				signDocs := models.SignDocs{
					Hash: docs.Hash,
				}
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), "JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu").Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				serviceUser.EXPECT().GetPublicKey(docs.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docs.Address = append(docs.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.IdSignature = append(docs.IdSignature, "signed_1")
				docs.IdSignature = append(docs.IdSignature, "signed2")
				docs.IdSignature = append(docs.IdSignature, docs.Creator_id)
				docs.Mode = "1"
				serviceSignature.EXPECT().AddToBlockhain(docs).Return(nil).Times(1)
				signDocs.Hash_original = docs.Hash_original
				signDocs.Creator = docs.Creator
				signDocs.IPFS = docs.IPFS
				serviceSignature.EXPECT().DocumentSigned(signDocs).Return(nil).Times(1)
				serviceUser.EXPECT().GetUserByEmail(docs.Email[0]).Times(1)
				serviceSignature.EXPECT().InvitePeople(docs.Email[0], docs, modelUser.User{}).Times(1)
				docs.Hash = signDocs.Hash
				serviceSignature.EXPECT().AddUserDocs(docs).Return(errors.New("Failed to insert data")).Times(1)
			},
		},
		{
			name:         "Test Sign Documents Success",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			docs: models.SignDocuments{
				Name:       "sample_test.pdf",
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: true,
				Email:      []string{"admin@rizwijaya.com", "smartsign@rizwijaya.com"},
				Creator:    "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
				Creator_id: "rizwijaya",
			},
			pages: "/download",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docs := models.SignDocuments{
					Name:       "sample_test.pdf",
					SignPage:   1.0,
					X_coord:    1.3,
					Y_coord:    1.2,
					Height:     4.2,
					Width:      5.3,
					Invite_sts: true,
					Email:      []string{"[61646d696e4072697a77696a6179612e636f6d 736d6172747369676e4072697a77696a6179612e636f6d]"},
				}

				path := "./public/temp/pdfsign/"
				serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				docs.Hash = docs.Hash_original
				signDocs := models.SignDocs{
					Hash: docs.Hash,
				}
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), "JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu").Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				serviceUser.EXPECT().GetPublicKey(docs.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docs.Address = append(docs.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.IdSignature = append(docs.IdSignature, "signed_1")
				docs.IdSignature = append(docs.IdSignature, "signed2")
				docs.IdSignature = append(docs.IdSignature, docs.Creator_id)
				docs.Mode = "1"
				serviceSignature.EXPECT().AddToBlockhain(docs).Return(nil).Times(1)
				signDocs.Hash_original = docs.Hash_original
				signDocs.Creator = docs.Creator
				signDocs.IPFS = docs.IPFS
				serviceSignature.EXPECT().DocumentSigned(signDocs).Return(nil).Times(1)
				serviceUser.EXPECT().GetUserByEmail(docs.Email[0]).Times(1)
				serviceSignature.EXPECT().InvitePeople(docs.Email[0], docs, modelUser.User{}).Times(1)
				docs.Hash = signDocs.Hash
				serviceSignature.EXPECT().AddUserDocs(docs).Return(nil).Times(1)
				serviceUser.EXPECT().Logging("Menandatangani dokumen "+docs.Name, "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			serviceUser := m_serviceUser.NewMockService(ctrl)
			//Testing Services Functions
			if tt.serviceTest != nil {
				tt.serviceTest(serviceUser, serviceSignature)
			}
			w := &signaturesHandler{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			got := w.SignDocuments
			router := NewRouter()
			router.POST("/sign-documents", got)
			//Testing Handler Functions
			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			err = writer.WriteField("signPage", fmt.Sprintf("%x", tt.docs.SignPage))
			assert.NoError(t, err)
			err = writer.WriteField("signX", fmt.Sprintf("%x", tt.docs.X_coord))
			assert.NoError(t, err)
			err = writer.WriteField("signY", fmt.Sprintf("%x", tt.docs.Y_coord))
			assert.NoError(t, err)
			err = writer.WriteField("signH", fmt.Sprintf("%x", tt.docs.Height))
			assert.NoError(t, err)
			err = writer.WriteField("signW", fmt.Sprintf("%x", tt.docs.Width))
			assert.NoError(t, err)
			err = writer.WriteField("invite_status", fmt.Sprintf("%v", tt.docs.Invite_sts))
			assert.NoError(t, err)
			err = writer.WriteField("email[]", fmt.Sprintf("%x", tt.docs.Email))
			assert.NoError(t, err)
			err = writer.WriteField("note", fmt.Sprintf("%x", tt.docs.Note))
			assert.NoError(t, err)
			err = writer.WriteField("judul", fmt.Sprintf("%x", tt.docs.Judul))
			assert.NoError(t, err)
			if tt.file != "" {
				path := "public/unit_testing/"
				file, errFile7 := os.Open(path + tt.file)
				assert.NoError(t, errFile7)
				defer file.Close()
				part7, errFile7 := CreateFilePDF(t, writer, path+tt.file)
				assert.NoError(t, errFile7)
				_, errFile7 = io.Copy(part7, file)
				assert.NoError(t, errFile7)
			}
			err := writer.Close()
			assert.Nil(t, err)

			req, err := http.NewRequest("POST", "/sign-documents", payload)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Cookie", cookies)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.responseCode, resp.Code)
			if tt.responseCode == http.StatusFound {
				location, err := resp.Result().Location()
				assert.NoError(t, err)
				assert.Equal(t, tt.pages, location.Path)
			} else {
				responseData, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Contains(t, string(responseData), "melakukan tanda tangan")
			}
		})
	}
}

func Test_signaturesHandler_InviteSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	_ = time.Now().In(location).String()
	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	tests := []struct {
		name         string
		responseCode int
		file         string
		invite       models.InviteSignatures
		pages        string
		serviceTest  func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService)
	}{
		{
			name:         "Test Invite Signature Input Invalid",
			responseCode: http.StatusFound,
			file:         "",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
			},
			pages: "/invite-signatures",
		},
		{
			name:         "Test Invite Signature Not File Request",
			responseCode: http.StatusFound,
			file:         "",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
				Note:  "Test Note Invite Signature",
			},
			pages: "/invite-signatures",
		},
		{
			name:         "Test Invite Signature Not File PDF",
			responseCode: http.StatusFound,
			file:         "card_test.jpeg",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
				Note:  "Test Note Invite Signature",
			},
			pages: "/invite-signatures",
		},
		{
			name:         "Test Invite Signature Failed to Input IPFS",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
				Note:  "Test Note Invite Signature",
			},
			pages: "/invite-signatures",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docData := models.SignDocuments{
					Name:  "sample_test.pdf",
					Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
					Judul: "Test Judul Invite Signature",
					Note:  "Test Note Invite Signature",
				}
				path := fmt.Sprintf("./public/temp/pdfsign/%s", docData.Name)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("", errors.New("Error Upload IPFS")).Times(1)
			},
		},
		{
			name:         "Test Invite Signature Failed to Add in Blockchain",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
				Note:  "Test Note Invite Signature",
			},
			pages: "/invite-signatures",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docData := models.SignDocuments{
					Name: "sample_test.pdf",
				}
				path := fmt.Sprintf("./public/temp/pdfsign/%s", docData.Name)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("hdsa734ndbams9032k2l22das", nil).Times(1)
				docData.IPFS = "hdsa734ndbams9032k2l22das"
				serviceUser.EXPECT().Encrypt([]byte("hdsa734ndbams9032k2l22das"), "JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu").Return([]byte("sr23sdajdadsasdasr546fgfdsfs")).Times(1)
				docData.IPFS = "sr23sdajdadsasdasr546fgfdsfs"
				docData.Email = []string{"[72697a71694072697a77696a6179612e636f6d 736d6172747369676e4072697a77696a6179612e636f6d]"}
				serviceUser.EXPECT().GetPublicKey(docData.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docData.Address = append(docData.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docData.Address = append(docData.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docData.IdSignature = append(docData.IdSignature, "signed_1")
				docData.IdSignature = append(docData.IdSignature, "signed2")
				docData.Mode = "2"
				docData.Judul = "54657374204a7564756c20496e76697465205369676e6174757265"
				docData.Note = "54657374204e6f746520496e76697465205369676e6174757265"
				serviceSignature.EXPECT().AddToBlockhain(docData).Return(errors.New("Failed to Add to Blockchain")).Times(1)
			},
		},
		{
			name:         "Test Invite Signature Failed to Add Documents",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
				Note:  "Test Note Invite Signature",
			},
			pages: "/invite-signatures",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docData := models.SignDocuments{
					Name: "sample_test.pdf",
				}
				path := fmt.Sprintf("./public/temp/pdfsign/%s", docData.Name)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("hdsa734ndbams9032k2l22das", nil).Times(1)
				docData.IPFS = "hdsa734ndbams9032k2l22das"
				serviceUser.EXPECT().Encrypt([]byte("hdsa734ndbams9032k2l22das"), "JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu").Return([]byte("sr23sdajdadsasdasr546fgfdsfs")).Times(1)
				docData.IPFS = "sr23sdajdadsasdasr546fgfdsfs"
				docData.Email = []string{"[72697a71694072697a77696a6179612e636f6d 736d6172747369676e4072697a77696a6179612e636f6d]"}
				serviceUser.EXPECT().GetPublicKey(docData.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docData.Address = append(docData.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docData.Address = append(docData.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docData.IdSignature = append(docData.IdSignature, "signed_1")
				docData.IdSignature = append(docData.IdSignature, "signed2")
				docData.Mode = "2"
				docData.Judul = "54657374204a7564756c20496e76697465205369676e6174757265"
				docData.Note = "54657374204e6f746520496e76697465205369676e6174757265"
				serviceSignature.EXPECT().AddToBlockhain(docData).Times(1)
				serviceUser.EXPECT().GetUserByEmail(docData.Email[0]).Times(1)
				serviceSignature.EXPECT().InvitePeople(docData.Email[0], docData, modelUser.User{}).Times(1)
				docData.IdSignature = append(docData.IdSignature, docData.Creator_id)
				docData.Address = append(docData.Address, common.HexToAddress(docData.Creator))
				serviceSignature.EXPECT().AddUserDocs(docData).Return(errors.New("Failed to insert data")).Times(1)
			},
		},
		{
			name:         "Test Invite Signature Success",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
				Note:  "Test Note Invite Signature",
			},
			pages: "/download",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docData := models.SignDocuments{
					Name: "sample_test.pdf",
				}
				path := fmt.Sprintf("./public/temp/pdfsign/%s", docData.Name)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("hdsa734ndbams9032k2l22das", nil).Times(1)
				docData.IPFS = "hdsa734ndbams9032k2l22das"
				serviceUser.EXPECT().Encrypt([]byte("hdsa734ndbams9032k2l22das"), "JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu").Return([]byte("sr23sdajdadsasdasr546fgfdsfs")).Times(1)
				docData.IPFS = "sr23sdajdadsasdasr546fgfdsfs"
				docData.Email = []string{"[72697a71694072697a77696a6179612e636f6d 736d6172747369676e4072697a77696a6179612e636f6d]"}
				serviceUser.EXPECT().GetPublicKey(docData.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docData.Address = append(docData.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docData.Address = append(docData.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docData.IdSignature = append(docData.IdSignature, "signed_1")
				docData.IdSignature = append(docData.IdSignature, "signed2")
				docData.Mode = "2"
				docData.Judul = "54657374204a7564756c20496e76697465205369676e6174757265"
				docData.Note = "54657374204e6f746520496e76697465205369676e6174757265"
				serviceSignature.EXPECT().AddToBlockhain(docData).Times(1)
				serviceUser.EXPECT().GetUserByEmail(docData.Email[0]).Times(1)
				serviceSignature.EXPECT().InvitePeople(docData.Email[0], docData, modelUser.User{}).Times(1)
				docData.IdSignature = append(docData.IdSignature, docData.Creator_id)
				docData.Address = append(docData.Address, common.HexToAddress(docData.Creator))
				serviceSignature.EXPECT().AddUserDocs(docData).Times(1)
				serviceUser.EXPECT().Logging("Mengundang orang lain untuk tanda tangan", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			serviceUser := m_serviceUser.NewMockService(ctrl)
			//Testing Services Functions
			if tt.serviceTest != nil {
				tt.serviceTest(serviceUser, serviceSignature)
			}
			w := &signaturesHandler{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			got := w.InviteSignatures
			router := NewRouter()
			router.POST("/invite-signatures", got)
			//Testing Handler Functions
			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			err = writer.WriteField("email[]", fmt.Sprintf("%x", tt.invite.Email))
			assert.NoError(t, err)
			err = writer.WriteField("judul", fmt.Sprintf("%x", tt.invite.Judul))
			assert.NoError(t, err)
			err = writer.WriteField("note", fmt.Sprintf("%x", tt.invite.Note))
			assert.NoError(t, err)
			if tt.file != "" {
				path := "public/unit_testing/"
				file, errFile7 := os.Open(path + tt.file)
				assert.NoError(t, errFile7)
				defer file.Close()
				part7, errFile7 := CreateFilePDF(t, writer, path+tt.file)
				assert.NoError(t, errFile7)
				_, errFile7 = io.Copy(part7, file)
				assert.NoError(t, errFile7)
			}
			err := writer.Close()
			assert.Nil(t, err)

			req, err := http.NewRequest("POST", "/invite-signatures", payload)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Cookie", cookies)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.responseCode, resp.Code)
			if tt.responseCode == http.StatusFound {
				location, err := resp.Result().Location()
				assert.NoError(t, err)
				assert.Equal(t, tt.pages, location.Path)
			}
		})
	}
}

func Test_signaturesHandler_Document(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location).String()
	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	mysignature := models.MySignatures{
		Id:                 "1",
		Name:               "Rizqi Wijaya",
		User_id:            "rizwijaya",
		Signature:          "default.png",
		Signature_id:       "sign_type",
		Signature_data:     "default.png",
		Signature_data_id:  "sign_type",
		Latin:              "latin.png",
		Latin_id:           "sign_type",
		Latin_data:         "latin_data.png",
		Latin_data_id:      "sign_type",
		Signature_selected: "signature",
		Date_update:        times,
		Date_created:       times,
	}

	tests := []struct {
		name         string
		responseCode int
		docs         models.SignDocuments
		hash         string
		pages        string
		serviceTest  func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService)
	}{
		{
			name:         "Test Document Invalid input",
			responseCode: http.StatusFound,
			hash:         "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
			docs: models.SignDocuments{
				Name:     "sample_test.pdf",
				SignPage: 1.0,
				X_coord:  1.3,
				Y_coord:  1.2,
				Height:   4.2,
			},
			pages: "/request-signatures",
		},
		{
			name:         "Test Document Failed Upload to IPFS",
			responseCode: http.StatusFound,
			hash:         "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
			docs: models.SignDocuments{
				// Name:     "sample_test.pdf",
				SignPage: 1.0,
				X_coord:  1.3,
				Y_coord:  1.2,
				Height:   4.2,
				Width:    5.3,
			},
			pages: "/request-signatures",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docs := models.SignDocuments{
					Name:          "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf",
					SignPage:      1.0,
					X_coord:       1.3,
					Y_coord:       1.2,
					Height:        4.2,
					Width:         5.3,
					Hash_original: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				}
				path := "./public/temp/pdfsign/"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf")
				serviceUser.EXPECT().UploadIPFS(path+"signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Return("", errors.New("Failed Upload to IPFS")).Times(1)
			},
		},
		{
			name:         "Test Document Signed Success",
			responseCode: http.StatusFound,
			hash:         "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
			docs: models.SignDocuments{
				// Name:     "sample_test.pdf",
				SignPage: 1.0,
				X_coord:  1.3,
				Y_coord:  1.2,
				Height:   4.2,
				Width:    5.3,
			},
			pages: "/download",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docs := models.SignDocuments{
					Name:          "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf",
					SignPage:      1.0,
					X_coord:       1.3,
					Y_coord:       1.2,
					Height:        4.2,
					Width:         5.3,
					Hash_original: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
				}
				path := "./public/temp/pdfsign/"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Return("js63hd9asn32nasddy783en9djas933").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Return("2c3idnsymuia7n8sb7dl92j63onsfdf", nil).Times(1)
				docs.Hash = "js63hd9asn32nasddy783en9djas933"
				docs.IPFS = "2c3idnsymuia7n8sb7dl92j63onsfdf"
				serviceUser.EXPECT().Encrypt([]byte("2c3idnsymuia7n8sb7dl92j63onsfdf"), "JWT_DAS3443HBOARDD_TAMS_RIZ_SK4343_343_KEJNF00975SDISu").Return([]byte("2dj3d1d6a34323ds4d4as43asda4sr5456fgfsdsfs")).Times(1)
				docs.IPFS = "2dj3d1d6a34323ds4d4as43asda4sr5456fgfsdsfs"
				signDocs := models.SignDocs{
					Hash_original: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Creator:       "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Hash:          docs.Hash,
					IPFS:          docs.IPFS,
				}
				serviceSignature.EXPECT().DocumentSigned(signDocs).Return(nil).Times(1)
				serviceUser.EXPECT().Logging("Melakukan tanda tangan dari permintaan tanda tangan", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			serviceUser := m_serviceUser.NewMockService(ctrl)
			//Testing Services Functions
			if tt.serviceTest != nil {
				tt.serviceTest(serviceUser, serviceSignature)
			}
			w := &signaturesHandler{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			got := w.Document
			router := NewRouter()
			router.POST("/document/:hash", got)
			//Testing Handler Functions
			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			err = writer.WriteField("signPage", fmt.Sprintf("%x", tt.docs.SignPage))
			assert.NoError(t, err)
			err = writer.WriteField("signX", fmt.Sprintf("%x", tt.docs.X_coord))
			assert.NoError(t, err)
			err = writer.WriteField("signY", fmt.Sprintf("%x", tt.docs.Y_coord))
			assert.NoError(t, err)
			err = writer.WriteField("signH", fmt.Sprintf("%x", tt.docs.Height))
			assert.NoError(t, err)
			err = writer.WriteField("signW", fmt.Sprintf("%x", tt.docs.Width))
			assert.NoError(t, err)
			err := writer.Close()
			assert.Nil(t, err)

			req, err := http.NewRequest("POST", "/document/"+tt.hash, payload)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Cookie", cookies)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.responseCode, resp.Code)
			if tt.responseCode == http.StatusFound {
				location, err := resp.Result().Location()
				assert.NoError(t, err)
				assert.Equal(t, tt.pages, location.Path)
			} else {
				responseData, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Contains(t, string(responseData), "melakukan tanda tangan")
			}
		})
	}
}
