package view

import (
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

func Test_userView_Index(t *testing.T) {
	t.Run("Test userView Index Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		//service := m_service.NewMockService(ctrl)
		serviceUser := m_serviceUser.NewMockService(ctrl)
		userView := NewUserView(serviceUser)
		router := NewRouter()
		router.GET("/", userView.Index)

		req, err := http.NewRequest("GET", "/", nil)
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		responseData, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, string(responseData), "SmartSign - Smart Digital Signatures")
	})
}

func Test_userView_Dashboard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		sign_id    string
		beforeTest func(userview *m_serviceUser.MockService)
	}{
		{
			name:    "Test userView Dashboard Success",
			sign_id: "rizwijaya",
			beforeTest: func(userview *m_serviceUser.MockService) {
				serviceUser.EXPECT().GetCardDashboard("rizwijaya").Times(1)
				serviceUser.EXPECT().Logging("Mengakses Halaman Dashboard", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &userview{
				userService: serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(serviceUser)
			}

			got := w.Dashboard

			router := NewRouter()
			router.GET("/dashboard", got)
			req, err := http.NewRequest("GET", "/dashboard", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Dashboard - SmartSign")
		})
	}
}

func Test_userview_Register(t *testing.T) {
	t.Run("Test userView Index Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		serviceUser := m_serviceUser.NewMockService(ctrl)
		userView := NewUserView(serviceUser)
		router := NewRouter()
		router.GET("/register", userView.Register)

		req, err := http.NewRequest("GET", "/register", nil)
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		responseData, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, string(responseData), "Pendaftaran - SmartSign")
	})
}

func Test_userview_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	test := []struct {
		name    string
		cookies string
	}{
		{
			name:    "Test userView Login Success",
			cookies: "kosong",
		},
		{
			name:    "Test userView Login with Register Notification Success",
			cookies: "registered=",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &userview{
				userService: serviceUser,
			}
			got := w.Login

			router := NewRouter()
			router.GET("/login", got)
			req, err := http.NewRequest("GET", "/login", nil)
			if tt.cookies != "kosong" {
				req.Header.Set("Cookie", tt.cookies)
			}
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Masuk - SmartSign")
		})
	}
}
