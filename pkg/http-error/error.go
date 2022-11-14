package error

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Form struct {
	Field   string
	Message string
}

func PageNotFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(404, "error_404.html", nil)
	}
}

func NoMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(404, gin.H{"status": "404", "message": "Method Not Found"})
	}
}

func FormValidationError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " wajib diisi!"
	case "email":
		return fe.Field() + " harus berupa email!"
	case "min":
		return fe.Field() + " minimal " + fe.Param() + " karakter!"
	case "max":
		return fe.Field() + " maksimal " + fe.Param() + " karakter!"
	case "alphanum":
		return fe.Field() + " hanya boleh berisi huruf dan angka!"
	case "numeric":
		return fe.Field() + " hanya boleh berisi angka!"
	case "eqfield":
		return fe.Field() + " harus sama dengan " + fe.Param() + "!"
	default:
		return fe.Field() + " tidak valid!"
	}
}
