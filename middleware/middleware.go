package middleware

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	userUC "github.com/mohfahrur/interop-service-c/usecase/user"

	"github.com/microcosm-cc/bluemonday"
)

func AuthAndAuthorize(userUC userUC.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("token") == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		token := sanitizeInput(c.GetHeader("token"))

		h := sha256.New()
		h.Write([]byte(token))
		bs := h.Sum(nil)
		sEnc := base64.StdEncoding.EncodeToString([]byte(bs))

		user, err := userUC.UserDomain.GetUserAuth(sEnc)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to retrieve user information"})
			return
		}
		if user == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		c.Set("userID", user.ID)
		c.Set("userRole", user.Role)

		isAuthorized := user.Role == "admin"
		if !isAuthorized {
			c.AbortWithStatusJSON(403, gin.H{"error": "Forbidden"})
			return
		}

		c.Next()
	}
}

func sanitizeInput(input string) string {
	p := bluemonday.StrictPolicy()
	sanitized := p.Sanitize(input)
	return sanitized
}
