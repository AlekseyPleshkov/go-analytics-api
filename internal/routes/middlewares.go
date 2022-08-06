package routes

import (
	"go-analytics/internal/routes/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	headerTokenKey = "token"
)

func tokenMiddlewere(token string) func(c *gin.Context) {
	return func(c *gin.Context) {
		headerToken := c.GetHeader(headerTokenKey)

		if headerToken != token {
			c.JSON(http.StatusUnauthorized, errors.NewError(errors.NotValidToken, nil))
			c.Abort()
			return
		}

		c.Next()
	}
}
