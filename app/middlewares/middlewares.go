package middlewares

import (
	"net/http"
	"strings"

	"e-signature/app/config"
	"e-signature/modules/v1/utilities/user/repository"
	"e-signature/modules/v1/utilities/user/service"
	respon "e-signature/pkg/api_response"
	crypto "e-signature/pkg/crypto"
	token "e-signature/pkg/jwt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		userIDSession := session.Get("id")

		if userIDSession == nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
	}
}

func AuthAPI(db *mongo.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		conf, _ := config.Init()
		session := sessions.Default(c)
		Repository := repository.NewRepository(db, nil)
		crypto := crypto.NewCrypto()
		serviceUser := service.NewService(Repository, crypto)

		authHeader := c.Request.Header.Get("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := respon.APIRespon("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := token.ValidateToken(tokenString)
		if err != nil {
			response := respon.APIRespon("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := respon.APIRespon("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		email := claim["email"].(string)
		pw := claim["pw"].(string)
		user, _ := serviceUser.GetUserByEmail(email)
		if len(user.Email) < 1 {
			response := respon.APIRespon("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		user.Publickey = string(crypto.Decrypt([]byte(user.Publickey), conf.App.Secret_key))
		session.Set("id", user.Id.Hex())
		session.Set("sign", user.Idsignature)
		session.Set("name", user.Name)
		session.Set("public_key", user.Publickey)
		session.Set("role", user.Role)
		session.Set("passph", string(crypto.Encrypt([]byte(pw), user.Publickey)))
		session.Save()
	}
}
