package authMiddleware

import (
	"net/http"

	"golang-gin-boilerplate/services/user"

	"golang-gin-boilerplate/helper"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type AuthMiddlerware interface {
	GetAuthMiddleware() gin.HandlerFunc
}

type middleware struct {
	service     Service
	userService user.Service
}

func CreateAuthMiddleware(service Service, userService user.Service) *middleware {
	return &middleware{service, userService}
}

func (m *middleware) GetAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		// check if ther's no prefix "Bearer"
		if !strings.Contains(authHeader, "Bearer") {
			helper.SendErrorResponse(c, "Unauthorized", http.StatusUnauthorized, "error", nil, nil)
			return
		}
		// split with space "Bearer tokentokentoken"
		arrayString := strings.Split(authHeader, " ")
		// get token string from array
		tokenString := arrayString[1]

		token, err := m.service.ValidateToken(tokenString)

		if err != nil {
			helper.SendErrorResponse(c, "Unauthorized", http.StatusUnauthorized, "error", nil, nil)
			return
		}
		// get data
		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			helper.SendErrorResponse(c, "Unauthorized", http.StatusUnauthorized, "error", nil, nil)
			return
		}
		// get id from claim (by default is float64) then convert it to int
		userId := int(claim["user_id"].(float64))
		// get user by id
		user, err := m.userService.GetUserById(userId)

		if err != nil {
			helper.SendErrorResponse(c, "Unauthorized", http.StatusUnauthorized, "error", nil, nil)
			return
		}
		// set currentUser on gin context
		c.Set("loggedInUser", user)
		c.Next()
	}
}
