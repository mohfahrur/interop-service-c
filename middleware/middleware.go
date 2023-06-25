package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mohfahrur/interop-service-c/entity"
	userUC "github.com/mohfahrur/interop-service-c/usecase/user"

	"github.com/microcosm-cc/bluemonday"
)

func AuthAndAuthorize(userUC userUC.UserUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		if os.Getenv("token") != c.GetHeader("token") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}
		var req entity.GetUserRequest
		err := c.BindJSON(&req)

		id := sanitizeInput(req.ID)

		user, err := userUC.GetUser(id)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to retrieve user information"})
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
