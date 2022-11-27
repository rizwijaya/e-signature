package view

import (
	"e-signature/modules/v1/utilities/signatures/models"
	m_serviceSignature "e-signature/modules/v1/utilities/signatures/service/mock"
	m_serviceUser "e-signature/modules/v1/utilities/user/service/mock"
	"e-signature/pkg/html"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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

func TestView(t *testing.T) {
	signatures := View(nil, nil, nil)
	assert.NotNil(t, signatures)
}

func Test_signaturesView_MySignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	serviceSignature := m_serviceSignature.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		beforeTest func()
	}{
		{
			name: "Test signaturesView MySignature Success",
			beforeTest: func() {
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Times(1)
				serviceUser.EXPECT().Logging("Mengakses tanda tangan saya", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesview{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got := w.MySignatures

			router := NewRouter()
			router.GET("/my-signatures", got)
			req, err := http.NewRequest("GET", "/my-signatures", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Tanda Tangan Saya - SmartSign")
		})
	}
}

func Test_signaturesView_SignDocuments(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	serviceSignature := m_serviceSignature.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		beforeTest func()
	}{
		{
			name: "Test signaturesView Sign Documents Success",
			beforeTest: func() {
				serviceSignature.EXPECT().GetMySignature("rizwijaya", "6380b5cbdc938c5fdf8e6bfe", "Rizqi Wijaya").Times(1)
				serviceUser.EXPECT().Logging("Mengakses tanda tangan dan minta tanda tangan", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesview{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got := w.SignDocuments

			router := NewRouter()
			router.GET("/sign-documents", got)
			req, err := http.NewRequest("GET", "/sign-documents", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Tanda Tangan Dokumen - SmartSign")
		})
	}
}

func Test_signaturesview_InviteSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	serviceSignature := m_serviceSignature.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		beforeTest func()
	}{
		{
			name: "Test signaturesView Invite Signatures Success",
			beforeTest: func() {
				serviceUser.EXPECT().Logging("Mengakses undang orang lain untuk tanda tangan", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesview{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got := w.InviteSignatures

			router := NewRouter()
			router.GET("/invite-signatures", got)
			req, err := http.NewRequest("GET", "/invite-signatures", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Undang untuk Tanda tangan - SmartSign")
		})
	}
}

func Test_signaturesview_RequestSignatures(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	serviceSignature := m_serviceSignature.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		beforeTest func()
	}{
		{
			name: "Test signaturesView Request Signatures Success",
			beforeTest: func() {
				serviceSignature.EXPECT().GetListDocument(gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengakses halaman permintaan tanda tangan", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesview{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got := w.RequestSignatures

			router := NewRouter()
			router.GET("/request-signatures", got)
			req, err := http.NewRequest("GET", "/request-signatures", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Daftar Permintaan Tanda Tangan - SmartSign")
		})
	}
}

func Test_signaturesview_Document(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	serviceSignature := m_serviceSignature.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		hash       string
		beforeTest func()
	}{
		{
			name: "Test signaturesView Request Signatures Allow Permission",
			hash: gomock.Any().String(),
			beforeTest: func() {
				serviceSignature.EXPECT().CheckSignature(gomock.Any(), gomock.Any()).Times(1)
				serviceSignature.EXPECT().GetDocument(gomock.Any(), gomock.Any()).Times(1)
				serviceSignature.EXPECT().GetMySignature(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().GetFileIPFS(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengakses dokumen untuk ditanda tangani", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
		{
			name: "Test signaturesView Request Signatures Not Allow Permission For Creator",
			hash: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
			beforeTest: func() {
				serviceSignature.EXPECT().CheckSignature(gomock.Any(), gomock.Any()).Times(1)

				serviceSignature.EXPECT().GetDocument("84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b", "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a").Return(models.DocumentBlockchain{
					Document_id:    "63820321334838fe2021b0fe",
					Creator:        common.HexToAddress("0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a"),
					Creator_string: "0xDBE4146513c99443cF32Ca8A449f5287aaD6f91a",
					Metadata:       "File-Testing.pdf",
					Hash_ori:       "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					Hash:           "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
					IPFS:           "6a4c244e82419f9d725e15d9cc4b",
					State:          "1",
					Mode:           "1",
					Createdtime:    "27112022",
					Completedtime:  "28112022",
					Exist:          true,
				}).Times(1)

				serviceSignature.EXPECT().GetMySignature(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().GetFileIPFS(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengakses dokumen untuk ditanda tangani", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
		{
			name: "Test signaturesView Request Signatures Failed to Get File From IPFS",
			hash: "84637c537106cb54272b66cda69f1bf51bd36a4c244e82419f9d725e15d9cc4b",
			beforeTest: func() {
				serviceSignature.EXPECT().CheckSignature(gomock.Any(), gomock.Any()).Times(1)
				serviceSignature.EXPECT().GetDocument(gomock.Any(), gomock.Any()).Times(1)
				serviceSignature.EXPECT().GetMySignature(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Times(1)
				serviceUser.EXPECT().GetFileIPFS("", ".pdf", "./public/temp/pdfsign/").Return("", errors.New("Failed to Get File From IPFS")).Times(1)
				serviceUser.EXPECT().Logging("Mengakses dokumen untuk ditanda tangani", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesview{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got := w.Document

			router := NewRouter()
			router.GET("/document/:hash", got)
			req, err := http.NewRequest("GET", "/document/"+tt.hash, nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, 301, resp.Code)
			assert.Contains(t, string(responseData), "Tanda Tangan Dokumen - SmartSign")
		})
	}
}

func Test_signaturesview_Verification(t *testing.T) {
	t.Run("Test signaturesView Verification Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		serviceUser := m_serviceUser.NewMockService(ctrl)
		serviceSignature := m_serviceSignature.NewMockService(ctrl)

		w := &signaturesview{
			serviceSignature: serviceSignature,
			serviceUser:      serviceUser,
		}

		router := NewRouter()
		router.GET("/verification", w.Verification)

		req, err := http.NewRequest("GET", "/verification", nil)
		assert.NoError(t, err)
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		responseData, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.Code)
		assert.Contains(t, string(responseData), "Verifikasi - SmartSign")
	})
}

func Test_signaturesview_History(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	serviceSignature := m_serviceSignature.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		beforeTest func()
	}{
		{
			name: "Test signaturesView History Success",
			beforeTest: func() {
				serviceSignature.EXPECT().GetListDocument(gomock.Any()).Times(1)
				serviceUser.EXPECT().Logging("Mengakses halaman riwayat tanda tangan", "rizwijaya", gomock.Any(), gomock.Any()).Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesview{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got := w.History

			router := NewRouter()
			router.GET("/history", got)
			req, err := http.NewRequest("GET", "/history", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Riwayat Tanda Tangan - SmartSign")
		})
	}
}

func Test_signaturesview_Transactions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	serviceUser := m_serviceUser.NewMockService(ctrl)
	serviceSignature := m_serviceSignature.NewMockService(ctrl)

	cookies := "smartsign=MTY2OTQ3NDEyOHxEdi1CQkFFQ180SUFBUkFCRUFBQV9nRXdfNElBQmdaemRISnBibWNNQkFBQ2FXUUdjM1J5YVc1bkRCb0FHRFl6T0RCaU5XTmlaR001TXpoak5XWmtaamhsTm1KbVpRWnpkSEpwYm1jTUJnQUVjMmxuYmdaemRISnBibWNNQ3dBSmNtbDZkMmxxWVhsaEJuTjBjbWx1Wnd3R0FBUnVZVzFsQm5OMGNtbHVad3dPQUF4U2FYcHhhU0JYYVdwaGVXRUdjM1J5YVc1bkRBd0FDbkIxWW14cFkxOXJaWGtHYzNSeWFXNW5EQ3dBS2pCNFJFSkZOREUwTmpVeE0yTTVPVFEwTTJOR016SkRZVGhCTkRRNVpqVXlPRGRoWVVRMlpqa3hZUVp6ZEhKcGJtY01CZ0FFY205c1pRTnBiblFFQWdBRUJuTjBjbWx1Wnd3SUFBWndZWE56Y0dnR2MzUnlhVzVuRERnQU5rWkNTQ3RMYkZwd1dHOHhlVTFSUTNnMU9VVTBNRnAxYlROWVVHa3dSbmxWT1c1TFVsTkRNbWR4UkhVNGJteFNSMHM0TTJkRlp3PT189RnNnJPqyThKonDOKwf4QeHI-7SwOwzto9OciAktNLw="
	test := []struct {
		name       string
		beforeTest func()
	}{
		{
			name: "Test signaturesView History Success",
			beforeTest: func() {
				serviceSignature.EXPECT().GetTransactions().Times(1)
			},
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			w := &signaturesview{
				serviceSignature: serviceSignature,
				serviceUser:      serviceUser,
			}
			if tt.beforeTest != nil {
				tt.beforeTest()
			}

			got := w.Transactions

			router := NewRouter()
			router.GET("/transactions", got)
			req, err := http.NewRequest("GET", "/transactions", nil)
			req.Header.Set("Cookie", cookies)
			assert.NoError(t, err)
			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)
			responseData, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			assert.Equal(t, http.StatusOK, resp.Code)
			assert.Contains(t, string(responseData), "Transaksi - SmartSign")
		})
	}
}
