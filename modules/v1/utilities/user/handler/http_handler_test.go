package user

import (
	"bytes"
	m_serviceSignature "e-signature/modules/v1/utilities/signatures/service/mock"
	"e-signature/modules/v1/utilities/user/models"
	m_serviceUser "e-signature/modules/v1/utilities/user/service/mock"
	"e-signature/pkg/html"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tkuchiki/faketime"

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

func Test_userHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	test := []struct {
		name         string
		idsignature  string
		password     string
		formValidate bool
		ResponseCode int
		pages        string
		beforeTest   func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService)
	}{
		{
			name:         "Test userHandler Login Success",
			idsignature:  "admin",
			password:     "admin",
			formValidate: true,
			ResponseCode: http.StatusFound,
			pages:        "/dashboard",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().Login(gomock.Any()).Times(1)
				serviceUser.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Masuk ke SmartSign", "", gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().Encrypt([]byte("admin"), gomock.Any()).Times(1)
			},
		},
		{
			name:         "Test userHandler Login Success",
			idsignature:  "admin",
			password:     "admin",
			formValidate: false,
			ResponseCode: http.StatusOK,
			pages:        "/login",
		},
		{
			name:         "Test userHandler Login Failed (Idsignature/Password Wrong)",
			idsignature:  "admin",
			password:     "admin",
			formValidate: true,
			ResponseCode: http.StatusOK,
			pages:        "/login",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().Login(gomock.Any()).Return(models.ProfileDB{}, errors.New("ID Signature/Password salah!")).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serviceUser := m_serviceUser.NewMockService(ctrl)
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			w := &userHandler{
				userService:      serviceUser,
				signatureService: serviceSignature,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(serviceSignature, serviceUser)
			}

			got := w.Login

			router := NewRouter()
			router.POST("/login", got)
			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			_ = writer.WriteField("idsignature", tt.idsignature)
			_ = writer.WriteField("password", tt.password)
			err := writer.Close()
			assert.Nil(t, err)

			req, err := http.NewRequest("POST", "/login", payload)
			assert.NoError(t, err)
			if tt.formValidate {
				req.Header.Set("Content-Type", writer.FormDataContentType())
			}
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.ResponseCode, resp.Code)
			if tt.ResponseCode == http.StatusFound {
				location, err := resp.Result().Location()
				assert.NoError(t, err)
				assert.Equal(t, tt.pages, location.Path)
			} else {
				responseData, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Contains(t, string(responseData), "Login - SmartSign")
			}
		})
	}
}

func Test_userHandler_Logout(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="

	test := []struct {
		name         string
		ResponseCode int
		pages        string
		beforeTest   func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService)
	}{
		{
			name:         "Test userHandler Logout Success",
			ResponseCode: http.StatusFound,
			pages:        "/",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().Logging("Keluar dari SmartSign", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serviceUser := m_serviceUser.NewMockService(ctrl)
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			w := &userHandler{
				userService:      serviceUser,
				signatureService: serviceSignature,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(serviceSignature, serviceUser)
			}

			got := w.Logout

			router := NewRouter()
			router.GET("/logout", got)

			req, err := http.NewRequest("GET", "/logout", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.ResponseCode, resp.Code)
			location, err := resp.Result().Location()
			assert.NoError(t, err)
			assert.Equal(t, tt.pages, location.Path)
		})
	}
}

// func () Matches(x interface{}) bool {
// 	return true
// }

func Test_userHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	location, err := time.LoadLocation("Asia/Jakarta")
	assert.NoError(t, err)
	user := models.User{
		Idsignature:    "adminsmartsign",
		Name:           "Administrator",
		Password:       "admin12345",
		Role:           2,
		Email:          "admin@smartsign.com",
		Phone:          "081234567890",
		Dateregistered: time.Now().In(location).String(),
		Identity_card:  "card-adminsmartsign.peg",
	}

	test := []struct {
		nameTest     string
		idsignature  string
		name         string
		email        string
		phone        string
		password     string
		cpassword    string
		file         string
		formValidate bool
		ResponseCode int
		pages        string
		beforeTest   func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService)
	}{
		{
			nameTest:     "Test userHandler Register Success",
			idsignature:  "adminsmartsign",
			name:         "Administrator",
			email:        "admin@smartsign.com",
			phone:        "081234567890",
			password:     "admin12345",
			cpassword:    "admin12345",
			file:         "card_test.jpeg",
			formValidate: true,
			ResponseCode: http.StatusFound,
			pages:        "/login",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().CheckUserExist("adminsmartsign").Times(1)
				serviceUser.EXPECT().CheckEmailExist("admin@smartsign.com").Times(1)
				serviceUser.EXPECT().EncryptFile("./public/images/identity_card/card-adminsmartsign.peg", "admin12345").Times(1)
				serviceUser.EXPECT().CreateAccount(user)
				serviceSignature.EXPECT().CreateLatinSignatures(user, "").Times(1)
				serviceSignature.EXPECT().CreateLatinSignaturesData(user, "", "").Times(1)
				serviceSignature.EXPECT().DefaultSignatures(user, "").Times(1)
			},
		},
		{
			nameTest:     "Test Register Failed Input Not Valid",
			idsignature:  "admin",
			name:         "Admin",
			email:        "admin@smartsign.com",
			phone:        "081234567890",
			password:     "admin",
			cpassword:    "admin",
			file:         "card_test.jpeg",
			formValidate: true,
			ResponseCode: http.StatusOK,
			pages:        "/register",
		},
		{
			nameTest:     "Test Registed Failed idsignature Exist",
			idsignature:  "adminsmartsign",
			name:         "Administrator",
			email:        "admin@smartsign.com",
			phone:        "081234567890",
			password:     "admin12345",
			cpassword:    "admin12345",
			file:         "card_test.jpeg",
			formValidate: true,
			ResponseCode: http.StatusOK,
			pages:        "/register",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().CheckUserExist("adminsmartsign").Return("exist", errors.New("Id Signature Exist")).Times(1)
			},
		},
		{
			nameTest:     "Test Register Failed Email Exist",
			idsignature:  "adminsmartsign",
			name:         "Administrator",
			email:        "admin@smartsign.com",
			phone:        "081234567890",
			password:     "admin12345",
			cpassword:    "admin12345",
			file:         "card_test.jpeg",
			formValidate: true,
			ResponseCode: http.StatusOK,
			pages:        "/register",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().CheckUserExist("adminsmartsign").Times(1)
				serviceUser.EXPECT().CheckEmailExist("admin@smartsign.com").Return("exist", errors.New("Email exist")).Times(1)
			},
		},
		{
			nameTest:     "Test Register Failed File Not Exist",
			idsignature:  "adminsmartsign",
			name:         "Administrator",
			email:        "admin@smartsign.com",
			phone:        "081234567890",
			password:     "admin12345",
			cpassword:    "admin12345",
			file:         "",
			formValidate: true,
			ResponseCode: http.StatusOK,
			pages:        "/register",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().CheckUserExist("adminsmartsign").Times(1)
				serviceUser.EXPECT().CheckEmailExist("admin@smartsign.com").Times(1)
			},
		},
		{
			nameTest:     "Test Register Failed Create Account",
			idsignature:  "adminsmartsign",
			name:         "Administrator",
			email:        "admin@smartsign.com",
			phone:        "081234567890",
			password:     "admin12345",
			cpassword:    "admin12345",
			file:         "card_test.jpeg",
			formValidate: true,
			ResponseCode: http.StatusOK,
			pages:        "/register",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().CheckUserExist("adminsmartsign").Times(1)
				serviceUser.EXPECT().CheckEmailExist("admin@smartsign.com").Times(1)
				serviceUser.EXPECT().EncryptFile("./public/images/identity_card/card-adminsmartsign.peg", "admin12345").Times(1)
				serviceUser.EXPECT().CreateAccount(user).Return("", errors.New("Create Account Failed")).Times(1)
			},
		},
		{
			nameTest:     "Test Register Failed to Save Default Signatures",
			idsignature:  "adminsmartsign",
			name:         "Administrator",
			email:        "admin@smartsign.com",
			phone:        "081234567890",
			password:     "admin12345",
			cpassword:    "admin12345",
			file:         "card_test.jpeg",
			formValidate: true,
			ResponseCode: http.StatusOK,
			pages:        "/register",
			beforeTest: func(serviceSignature *m_serviceSignature.MockService, serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().CheckUserExist("adminsmartsign").Times(1)
				serviceUser.EXPECT().CheckEmailExist("admin@smartsign.com").Times(1)
				serviceUser.EXPECT().EncryptFile("./public/images/identity_card/card-adminsmartsign.peg", "admin12345").Times(1)
				serviceUser.EXPECT().CreateAccount(user).Times(1)
				serviceSignature.EXPECT().CreateLatinSignatures(user, "").Times(1)
				serviceSignature.EXPECT().CreateLatinSignaturesData(user, "", "").Times(1)
				serviceSignature.EXPECT().DefaultSignatures(user, "").Return(errors.New("Failed to save default signatures")).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.nameTest, func(t *testing.T) {
			serviceSignature := m_serviceSignature.NewMockService(ctrl)
			serviceUser := m_serviceUser.NewMockService(ctrl)

			w := &userHandler{
				userService:      serviceUser,
				signatureService: serviceSignature,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(serviceSignature, serviceUser)
			}

			got := w.Register

			router := NewRouter()
			router.POST("/register", got)
			//Payload POST Request
			payload := &bytes.Buffer{}
			writer := multipart.NewWriter(payload)
			err = writer.WriteField("idsignature", tt.idsignature)
			assert.NoError(t, err)
			err = writer.WriteField("name", tt.name)
			assert.NoError(t, err)
			err = writer.WriteField("email", tt.email)
			assert.NoError(t, err)
			err = writer.WriteField("phone", tt.phone)
			assert.NoError(t, err)
			err = writer.WriteField("password", tt.password)
			assert.NoError(t, err)
			err = writer.WriteField("cpassword", tt.cpassword)
			assert.NoError(t, err)
			if tt.file != "" {
				path := "public/unit_testing/"
				file, errFile7 := os.Open(path + tt.file)
				assert.NoError(t, errFile7)
				defer file.Close()
				part7, errFile7 := writer.CreateFormFile("file", filepath.Base(path+tt.file))
				assert.NoError(t, errFile7)
				_, errFile7 = io.Copy(part7, file)
				assert.NoError(t, errFile7)
			}
			err := writer.Close()
			assert.NoError(t, err)
			//Request to URL Register with Method POST
			req, err := http.NewRequest("POST", "/register", payload)
			assert.NoError(t, err)
			if tt.formValidate {
				req.Header.Set("Content-Type", writer.FormDataContentType())
			}
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.ResponseCode, resp.Code)
			if tt.ResponseCode == http.StatusFound {
				location, err := resp.Result().Location()
				assert.NoError(t, err)
				assert.Equal(t, tt.pages, location.Path)
			} else {
				responseData, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)
				assert.Contains(t, string(responseData), "Pendaftaran - SmartSign")
			}
		})
	}
}
