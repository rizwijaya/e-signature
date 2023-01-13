package view

import (
	"e-signature/app/config"
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

func Test_userView_Index(t *testing.T) {
	t.Run("Landing Page Case 1: Success View Landing Page", func(t *testing.T) {
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

	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="
	test := []struct {
		name       string
		sign_id    string
		beforeTest func(serviceUser *m_serviceUser.MockService)
	}{
		{
			name:    "Dashboard Case 1: Success View Dashboard",
			sign_id: "rizwijaya",
			beforeTest: func(serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().GetCardDashboard("rizwijaya").Times(1)
				serviceUser.EXPECT().Logging("Mengakses Halaman Dashboard", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serviceUser := m_serviceUser.NewMockService(ctrl)
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
	t.Run("Register Case 1: Success View Register", func(t *testing.T) {
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

	test := []struct {
		name    string
		cookies string
	}{
		{
			name:    "Login Case 1: Success View Login",
			cookies: "kosong",
		},
		{
			name:    "Login Case 2: Success View Login with Register Notification Success",
			cookies: "registered=",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serviceUser := m_serviceUser.NewMockService(ctrl)
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

func Test_userview_Logg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cookies := "smartsign=MTY3MjExNzQ3OHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQ0FBR2NHRnpjM0JvQm5OMGNtbHVad3c0QURZMU9FVmphMUZ4YWk5RVduTjRPVTkwZEVFNWVWRkdMMWs0TlV4UFJscEJiRlp3YUdOUFluTjRMMGhxYTA5U1MyTlRjR1ZTWjFFR2MzUnlhVzVuREFRQUFtbGtCbk4wY21sdVp3d2FBQmcyTXprMVpUWTVZMlZpWXpneFpqQmxZVEZtTm1JM05XUUdjM1J5YVc1bkRBWUFCSE5wWjI0R2MzUnlhVzVuREFzQUNYSnBlbmRwYW1GNVlRWnpkSEpwYm1jTUJnQUVibUZ0WlFaemRISnBibWNNRGdBTVVtbDZjV2tnVjJscVlYbGhCbk4wY21sdVp3d01BQXB3ZFdKc2FXTmZhMlY1Qm5OMGNtbHVad3dzQUNvd2VFTTBORE5qWmtaak56azFaakl4TkRrell6WTRPVEV4WmpoR04wUmlaVFpEWTJJNFFqTTJPRElHYzNSeWFXNW5EQVlBQkhKdmJHVURhVzUwQkFJQUJBPT18REvJoQ2Um3XpCGKLgxqGGWqS1ozjSafhkr2Svy2HHaM="
	test := []struct {
		name       string
		sign_id    string
		beforeTest func(serviceUser *m_serviceUser.MockService)
	}{
		{
			name:    "Log User Case 1: Success View Log User",
			sign_id: "rizwijaya",
			beforeTest: func(serviceUser *m_serviceUser.MockService) {
				serviceUser.EXPECT().GetLogUser("rizwijaya")
				serviceUser.EXPECT().Logging("Mengakses halaman log akses user", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			serviceUser := m_serviceUser.NewMockService(ctrl)
			w := &userview{
				userService: serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest(serviceUser)
			}

			got := w.Logg

			router := NewRouter()
			router.GET("/log-user", got)
			req, err := http.NewRequest("GET", "/log-user", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Log Akses User - SmartSign")
		})
	}
}

func TestView(t *testing.T) {
	user := View(nil, nil, nil)
	assert.NotNil(t, user)
}
