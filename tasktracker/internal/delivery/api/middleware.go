package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (a *Api) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if headerParts[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//TODO: В результате парсинга токена из payload получаем PublicId
		//			Возможно, его нужно передавать в основной процесс обработки как автора запроса
		_, err := a.auth.ParseToken(headerParts[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
