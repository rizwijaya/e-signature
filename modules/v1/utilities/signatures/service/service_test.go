package service

import (
	m_repo "e-signature/modules/v1/utilities/signatures/repository/mock"
	m_docs "e-signature/pkg/document/mock"
	m_images "e-signature/pkg/images/mock"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/tkuchiki/faketime"
)

func TestInit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	os.Chdir("../../../../../")
}

func Test_service_NewService(t *testing.T) {
	t.Run("Case 1: Success New Service", func(t *testing.T) {
		serv := NewService(nil, nil, nil)
		assert.NotNil(t, serv)
	})
}

func Test_service_TimeFormating(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//var err error
	f := faketime.NewFaketime(2022, time.November, 27, 11, 30, 01, 0, time.UTC)
	defer f.Undo()
	f.Do()
	// location, err := time.LoadLocation("Asia/Jakarta")
	// assert.NoError(t, err)
	// times := time.Now().In(location)

	test := []struct {
		name       string
		time       string
		timeFormat string
	}{
		{
			name:       "Time Formatting Case 1: Success Convert Format Time 14 Digit",
			time:       "15040528022022",
			timeFormat: "Senin, 28 Feb 2022 | 15:04 WIB",
		},
		{
			name:       "Time Formatting Case 2: Success Convert Format Time 13 Digit",
			time:       "7040526022022",
			timeFormat: "Sabtu, 26 Feb 2022 | 07:04 WIB",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			repo := m_repo.NewMockRepository(ctrl)
			images := m_images.NewMockImages(ctrl)
			docs := m_docs.NewMockDocuments(ctrl)
			s := NewService(repo, images, docs)
			timeformat := s.TimeFormating(tt.time)
			assert.Equal(t, tt.timeFormat, timeformat)
		})
	}
}
