package view

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_userView_Index(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *userView
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.Index(tt.args.c)
		})
	}
}
