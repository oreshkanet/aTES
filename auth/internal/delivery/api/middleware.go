package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func getToken(c *gin.Context) (string, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("token is undefined")
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 {
		return "", fmt.Errorf("invalid token")
	}

	if headerParts[0] != "Bearer" {
		return "", fmt.Errorf("invalid token")
	}

	return headerParts[1], nil
}

func (a *Api) UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//TODO: В результате парсинга токена из payload получаем PublicId
		//			Возможно, его нужно передавать в основной процесс обработки как автора запроса
		claims := new(*jwt.Claims)
		err = a.auth.ParseToken(token, claims)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// TODO: В claims лежат PublicId и Role. Надо проверить доступно ли для Role
		// 			Параметры доступа засунуть в миддлвар или в API
		// if claims.Role ...

		c.Next()
	}
}
