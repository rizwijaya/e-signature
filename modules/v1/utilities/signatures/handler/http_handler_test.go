package signatures

import (
	"bytes"
	"e-signature/app/config"
	"e-signature/modules/v1/utilities/signatures/models"
	m_serviceSignature "e-signature/modules/v1/utilities/signatures/service/mock"
	modelUser "e-signature/modules/v1/utilities/user/models"
	m_serviceUser "e-signature/modules/v1/utilities/user/service/mock"
	"e-signature/pkg/html"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path/filepath"
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
	conf, _ := config.Init()

	cookieStore := cookie.NewStore([]byte(conf.App.Secret_key))
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

	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="

	tests := []struct {
		name        string
		statusCode  int
		request     string
		response    string
		serviceTest func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService)
	}{
		{
			name:       "Add Signatures Case 1: Success",
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
			name:       "Add Signatures Case 2: Input Not Valid",
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

	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="

	tests := []struct {
		name        string
		statusCode  int
		sign_type   string
		sign_now    string
		pages       string
		serviceTest func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService)
	}{
		{
			name:       "Change Signatures Case 1: Change to Signature Success",
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
			name:       "Change Signatures Case 2: Change to Signature Data Success",
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
			name:       "Change Signatures Case 3: Change to Latin Success",
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
			name:       "Change Signatures Case 4: Change to Latin Data Success",
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
			name:       "Change Signatures Case 5: Input Random Sign Type",
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
	conf, _ := config.Init()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location).String()
	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="

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
			name:         "Sign Documents Case 1: Input Invalid",
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
			name:         "Sign Documents Case 2: Not File Document Request",
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
			name:         "Sign Documents Case 3:  Not File format PDF in Request",
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
			name:         "Sign Documents Case 4: Failed to Input IPFS",
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
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
				//serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				//docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				docs.Hash = docs.Hash_original
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("", errors.New("Failed to Input IPFS")).Times(1)
			},
		},
		{
			name:         "Sign Documents Case 5: No Invite Signers and Failed Add To Blockchain",
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
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
				//serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Hash = docs.Hash_original
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), conf.App.Secret_key).Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				docs.Address = append(docs.Address, common.HexToAddress("0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"))
				docs.IdSignature = append(docs.IdSignature, docs.Creator_id)
				docs.Mode = "3"
				serviceSignature.EXPECT().AddToBlockhain(docs).Return(errors.New("Failed to Add to Blockchain")).Times(1)
			},
		},
		{
			name:         "Sign Documents Case 6: Invite Signers and Failed to Add User Documents",
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
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
				// serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Hash = docs.Hash_original
				signDocs := models.SignDocs{
					Hash: docs.Hash,
				}
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), conf.App.Secret_key).Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				serviceUser.EXPECT().GetPublicKey(docs.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docs.Address = append(docs.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"))
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
			name:         "Sign Documents Case 7: Success Signed Document",
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
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
				// serviceSignature.EXPECT().GenerateHashDocument(path + docs.Name).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				// docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Hash = docs.Hash_original
				signDocs := models.SignDocs{
					Hash: docs.Hash,
				}
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), conf.App.Secret_key).Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				serviceUser.EXPECT().GetPublicKey(docs.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docs.Address = append(docs.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"))
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
				var part7 io.Writer
				if tt.file[len(tt.file)-3:] == "pdf" {
					part7, errFile7 = CreateFilePDF(t, writer, path+tt.file)
				} else {
					part7, errFile7 = writer.CreateFormFile("file", filepath.Base(path+tt.file))
				}
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
	conf, _ := config.Init()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="

	tests := []struct {
		name         string
		responseCode int
		file         string
		invite       models.InviteSignatures
		pages        string
		serviceTest  func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService)
	}{
		{
			name:         "Invite Signature Case 1: Input Invalid",
			responseCode: http.StatusFound,
			file:         "",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature",
			},
			pages: "/invite-signatures",
		},
		{
			name:         "Invite Signature Case 2: Not File Document in Request",
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
			name:         "Invite Signature Case 3: Not File PDF Format in Request",
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
			name:         "Invite Signature Case 4: Failed Create Watermark in Document",
			responseCode: http.StatusFound,
			file:         "sample_test.pdf",
			invite: models.InviteSignatures{
				Email: []string{"rizqi@rizwijaya.com", "smartsign@rizwijaya.com"},
				Judul: "Test Judul Invite Signature 2",
				Note:  "Test Note Invite Signature 2",
			},
			pages: "/invite-signatures",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				docData := models.SignDocuments{
					Name:  "sample_test.pdf",
					Email: []string{"rizqi@rizwijaya.com 2", "smartsign@rizwijaya.com 2"},
					Judul: "Test Judul Invite Signature",
					Note:  "Test Note Invite Signature",
				}
				path := fmt.Sprintf("./public/temp/pdfsign/%s", docData.Name)
				serviceSignature.EXPECT().WaterMarking(path).Return("").Times(1)
			},
		},
		{
			name:         "Invite Signature Case 5: Failed Input Document to IPFS",
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
				serviceSignature.EXPECT().WaterMarking(path).Return(path).Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("", errors.New("Error Upload IPFS")).Times(1)
			},
		},
		{
			name:         "Invite Signature Case 6: Failed Add Document to Blockchain",
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
				serviceSignature.EXPECT().WaterMarking(path).Return(path).Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("hdsa734ndbams9032k2l22das", nil).Times(1)
				docData.IPFS = "hdsa734ndbams9032k2l22das"
				serviceUser.EXPECT().Encrypt([]byte("hdsa734ndbams9032k2l22das"), conf.App.Secret_key).Return([]byte("sr23sdajdadsasdasr546fgfdsfs")).Times(1)
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
			name:         "Invite Signature Case 7: Failed Add Documents to Database",
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
				serviceSignature.EXPECT().WaterMarking(path).Return(path).Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("hdsa734ndbams9032k2l22das", nil).Times(1)
				docData.IPFS = "hdsa734ndbams9032k2l22das"
				serviceUser.EXPECT().Encrypt([]byte("hdsa734ndbams9032k2l22das"), conf.App.Secret_key).Return([]byte("sr23sdajdadsasdasr546fgfdsfs")).Times(1)
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
			name:         "Invite Signature Case 8: Success Invite Signers",
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
				serviceSignature.EXPECT().WaterMarking(path).Return(path).Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docData.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docData.Hash = docData.Hash_original
				docData.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docData.Creator_id = "rizwijaya"
				serviceUser.EXPECT().UploadIPFS(path).Return("hdsa734ndbams9032k2l22das", nil).Times(1)
				docData.IPFS = "hdsa734ndbams9032k2l22das"
				serviceUser.EXPECT().Encrypt([]byte("hdsa734ndbams9032k2l22das"), conf.App.Secret_key).Return([]byte("sr23sdajdadsasdasr546fgfdsfs")).Times(1)
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
				var part7 io.Writer
				if tt.file[len(tt.file)-3:] == "pdf" {
					part7, errFile7 = CreateFilePDF(t, writer, path+tt.file)
				} else {
					part7, errFile7 = writer.CreateFormFile("file", filepath.Base(path+tt.file))
				}
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
	conf, _ := config.Init()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location).String()
	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="
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
			name:         "Document Sign Now Case 1: Invalid input",
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
			name:         "Document Sign Now Case 2: Failed Upload to IPFS",
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
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf")
				serviceUser.EXPECT().UploadIPFS(path+"signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Return("", errors.New("Failed Upload to IPFS")).Times(1)
			},
		},
		{
			name:         "Document Sign Now Case 3: Document Signed Success",
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
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Return("js63hd9asn32nasddy783en9djas933").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b.pdf").Return("2c3idnsymuia7n8sb7dl92j63onsfdf", nil).Times(1)
				docs.Hash = "js63hd9asn32nasddy783en9djas933"
				docs.IPFS = "2c3idnsymuia7n8sb7dl92j63onsfdf"
				serviceUser.EXPECT().Encrypt([]byte("2c3idnsymuia7n8sb7dl92j63onsfdf"), conf.App.Secret_key).Return([]byte("2dj3d1d6a34323ds4d4as43asda4sr5456fgfsdsfs")).Times(1)
				docs.IPFS = "2dj3d1d6a34323ds4d4as43asda4sr5456fgfsdsfs"
				signDocs := models.SignDocs{
					Hash_original: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Creator:       "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
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

func Test_signaturesHandler_Verification(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="

	tests := []struct {
		name         string
		responseCode int
		file         string
		pages        string
		serviceTest  func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService)
	}{
		{
			name:         "Verification Case 1: Input File Document Invalid",
			responseCode: http.StatusFound,
			file:         "",
			pages:        "/verification",
		},
		{
			name:         "Verification Case 2: Document Not PDF Format",
			responseCode: http.StatusFound,
			file:         "card_test.jpeg",
			pages:        "/verification",
		},
		{
			name:         "Verification Case 3: Document Not Signed Status Success",
			responseCode: http.StatusOK,
			file:         "sample_test.pdf",
			pages:        "/verification",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				path := "./public/temp/pdfverify/sample_test.pdf"
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceSignature.EXPECT().GetDocumentAllSign("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Return(models.DocumentAllSign{}, true).Times(1)
			},
		},
		{
			name:         "Verification Case 4: Document Signed Status Success",
			responseCode: http.StatusOK,
			file:         "sample_test.pdf",
			pages:        "/verification",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				path := "./public/temp/pdfverify/sample_test.pdf"
				serviceSignature.EXPECT().GenerateHashDocument(path).Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceSignature.EXPECT().GetDocumentAllSign("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Return(models.DocumentAllSign{}, false).Times(1)
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
			got := w.Verification
			router := NewRouter()
			router.POST("/verification", got)
			//Testing Handler Functions
			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			if tt.file != "" {
				path := "public/unit_testing/"
				file, errFile7 := os.Open(path + tt.file)
				assert.NoError(t, errFile7)
				defer file.Close()
				var part7 io.Writer
				if tt.file[len(tt.file)-3:] == "pdf" {
					part7, errFile7 = CreateFilePDF(t, writer, path+tt.file)
				} else {
					part7, errFile7 = writer.CreateFormFile("file", filepath.Base(path+tt.file))
				}
				assert.NoError(t, errFile7)
				_, errFile7 = io.Copy(part7, file)
				assert.NoError(t, errFile7)
				err := writer.Close()
				assert.Nil(t, err)
			}

			req, err := http.NewRequest("POST", "/verification", payload)
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
				assert.Contains(t, string(responseData), "Hasil Verifikasi - SmartSign")
			}
		})
	}
}

func Test_signaturesHandler_Download(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conf, _ := config.Init()

	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="

	tests := []struct {
		name         string
		responseCode int
		hash         string
		pages        string
		serviceTest  func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService)
	}{
		{
			name:         "Download Case 1: Download Documents Success",
			responseCode: http.StatusFound,
			hash:         "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
			pages:        "/download",
			serviceTest: func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService) {
				doc := models.DocumentBlockchain{
					// Document_id: "1",
					// Creator: common.HexToAddress("0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"),
					// Creator_string: ""
					Metadata: "sample_test.pdf",
					IPFS:     "2c3idnsymuia7n8sb7dl92j63onsfdf",
				}
				hash := "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				serviceSignature.EXPECT().GetDocumentNoSigners(hash).Return(doc).Times(1)
				serviceUser.EXPECT().Decrypt([]byte(doc.IPFS), conf.App.Secret_key).Return([]byte("2dj3d1d6a34323ds4d4as43asda4sr5456fgfsdsfs")).Times(1)
				doc.IPFS = "2dj3d1d6a34323ds4d4as43asda4sr5456fgfsdsfs"
				directory := "./public/temp/pdfdownload/"
				serviceUser.EXPECT().GetFileIPFS(doc.IPFS, doc.Metadata+".pdf", directory)
				serviceUser.EXPECT().Logging("Mengunduh dokumen "+doc.Metadata+".pdf", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
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
			got := w.Download
			router := NewRouter()
			router.GET("/download/:hash", got)
			req, err := http.NewRequest("GET", "/download/"+tt.hash, nil)
			assert.NoError(t, err)
			req.Header.Set("Cookie", cookies)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.responseCode, resp.Code)
			location, err := resp.Result().Location()
			assert.NoError(t, err)
			assert.Equal(t, tt.pages, location.Path)
		})
	}
}

func Test_signaturesHandler_Integrity(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	conf, _ := config.Init()

	type Meta struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Status  string `json:"status"`
	}
	type Response struct {
		Meta Meta `json:"meta"`
	}

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	times := time.Now().In(location).String()
	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="

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
		message      string
		serviceTest  func(serviceUser *m_serviceUser.MockService, serviceSignature *m_serviceSignature.MockService)
	}{
		{
			name:         "Integrity Analysis Case 1: Input Invalid",
			responseCode: http.StatusNonAuthoritativeInfo,
			file:         "",
			docs: models.SignDocuments{
				SignPage:   1.0,
				X_coord:    1,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: false,
			},
			message: "Data yang anda masukan salah!",
		},
		{
			name:         "Integrity Analysis Case 2: Not File Document Request",
			responseCode: http.StatusNonAuthoritativeInfo,
			file:         "",
			docs: models.SignDocuments{
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: false,
			},
			message: "Data yang anda masukan salah!",
		},
		{
			name:         "Integrity Analysis Case 3:  Not File format PDF in Request",
			responseCode: http.StatusNonAuthoritativeInfo,
			file:         "card_test.jpeg",
			docs: models.SignDocuments{
				SignPage:   1.0,
				X_coord:    1.3,
				Y_coord:    1.2,
				Height:     4.2,
				Width:      5.3,
				Invite_sts: false,
			},
			message: "Data yang anda masukan salah!",
		},
		{
			name:         "Integrity Analysis Case 4: Failed to Input IPFS",
			responseCode: http.StatusNonAuthoritativeInfo,
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
				Creator_id: "rizwijaya",
			},
			message: "Data yang anda masukan salah!",
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
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Hash = docs.Hash_original
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("", errors.New("Failed to Input IPFS")).Times(1)
			},
		},
		{
			name:         "Integrity Analysis Case 5: No Invite Signers and Failed Add To Blockchain",
			responseCode: http.StatusNonAuthoritativeInfo,
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
				Creator_id: "rizwijaya",
			},
			message: "Data yang anda masukan salah!",
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
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Hash = docs.Hash_original
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), conf.App.Secret_key).Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				docs.Address = append(docs.Address, common.HexToAddress("0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"))
				docs.IdSignature = append(docs.IdSignature, docs.Creator_id)
				docs.Mode = "3"
				serviceSignature.EXPECT().AddToBlockhain(docs).Return(errors.New("Failed to Add to Blockchain")).Times(1)
			},
		},
		{
			name:         "Integrity Analysis Case 6: Invite Signers and Failed to Add User Documents",
			responseCode: http.StatusNonAuthoritativeInfo,
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
				Creator_id: "rizwijaya",
			},
			message: "Data yang anda masukan salah!",
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
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Hash = docs.Hash_original
				signDocs := models.SignDocs{
					Hash: docs.Hash,
				}
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), conf.App.Secret_key).Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				serviceUser.EXPECT().GetPublicKey(docs.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docs.Address = append(docs.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"))
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
			name:         "Integrity Analysis Case 7: Success Signed Document",
			responseCode: http.StatusOK,
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
				Creator:    "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682",
				Creator_id: "rizwijaya",
			},
			message: "Berhasil melakukan tanda tangan dokumen!",
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
				docs.Creator = "0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"
				docs.Creator_id = "rizwijaya"
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6395e69cebc81f0ea1f6b75d", "Rizqi Wijaya").Return(mysignature).Times(1)
				serviceSignature.EXPECT().ResizeImages(mysignature, docs).Return("./public/temp/sizes-signature.png").Times(1)
				serviceSignature.EXPECT().SignDocuments("./public/temp/sizes-signature.png", docs).Return(path + "signed_sample_test.pdf").Times(1)
				serviceSignature.EXPECT().GenerateHashDocument(path + "signed_sample_test.pdf").Return("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b").Times(1)
				docs.Hash_original = "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b"
				docs.Hash = docs.Hash_original
				signDocs := models.SignDocs{
					Hash: docs.Hash,
				}
				serviceUser.EXPECT().UploadIPFS(path+"signed_sample_test.pdf").Return("j8329dnsay80e2asdas", nil).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("j8329dnsay80e2asdas"), conf.App.Secret_key).Return([]byte("jdadsasdasr546fgfdsfs")).Times(1)
				docs.IPFS = "jdadsasdasr546fgfdsfs"
				serviceUser.EXPECT().GetPublicKey(docs.Email).Return([]common.Address{common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"), common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a")}, []string{"signed_1", "signed2"}).Times(1)
				docs.Address = append(docs.Address, common.HexToAddress("0xAyysae6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xBha62e6513c99443cF32Ca8A449f5287aaD6f91a"))
				docs.Address = append(docs.Address, common.HexToAddress("0xC443cfFc795f21493c68911f8F7Dbe6Ccb8B3682"))
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
				serviceUser.EXPECT().Logging("Menandatangani dokumen "+docs.Name+" via API", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
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
			got := w.Integrity
			router := NewRouter()
			router.POST("/api/v1/analysis/integrity-document", got)
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
				var part7 io.Writer
				if tt.file[len(tt.file)-3:] == "pdf" {
					part7, errFile7 = CreateFilePDF(t, writer, path+tt.file)
				} else {
					part7, errFile7 = writer.CreateFormFile("file", filepath.Base(path+tt.file))
				}
				assert.NoError(t, errFile7)
				_, errFile7 = io.Copy(part7, file)
				assert.NoError(t, errFile7)
			}
			err := writer.Close()
			assert.Nil(t, err)

			req, err := http.NewRequest("POST", "/api/v1/analysis/integrity-document", payload)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			req.Header.Set("Cookie", cookies)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			var response Response
			err = json.Unmarshal(responseData, &response)
			assert.NoError(t, err)

			assert.Equal(t, tt.responseCode, resp.Code)
			assert.Equal(t, resp.Code, response.Meta.Code)
			assert.Equal(t, tt.message, response.Meta.Message)
		})
	}
}
