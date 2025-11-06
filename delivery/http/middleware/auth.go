package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"inventory-ticketing-system/infrastructure/jwt"
	"inventory-ticketing-system/pkg/common"
)

func AuthMiddleware(jwtManager *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.SendError(c, 401, "UNAUTHORIZED", "Authorization header required", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			common.SendError(c, 401, "UNAUTHORIZED", "Bearer token required", nil)
			c.Abort()
			return
		}

		userID, role, err := jwtManager.ValidateToken(tokenString)
		if err != nil {
			common.SendError(c, 401, "UNAUTHORIZED", "Invalid token", nil)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("user_role", role)
		c.Next()
	}
}

func RoleMiddleware(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("user_role")
		if !exists {
			common.SendError(c, 401, "UNAUTHORIZED", "User role not found", nil)
			c.Abort()
			return
		}

		role, ok := userRole.(string)
		if !ok {
			common.SendError(c, 500, "INTERNAL_ERROR", "Invalid user role format", nil)
			c.Abort()
			return
		}

		authorized := false
		for _, r := range roles {
			if r == role {
				authorized = true
				break
			}
		}

		if !authorized {
			common.SendError(c, 403, "FORBIDDEN", "Insufficient permissions", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}

func GetUserID(c *gin.Context) (uuid.UUID, error) {
	userID, exists := c.Get("user_id")
	if !exists {
		return uuid.Nil, gin.Error{}
	}

	return userID.(uuid.UUID), nil
}

func GetUserRole(c *gin.Context) (string, error) {
	userRole, exists := c.Get("user_role")
	if !exists {
		return "", gin.Error{}
	}

	return userRole.(string), nil
}