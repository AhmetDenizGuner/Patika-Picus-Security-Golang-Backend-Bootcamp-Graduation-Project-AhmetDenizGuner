package middleware

import (
	jwtHelper "github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TODO check redis
func UserAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				for _, role := range decodedClaims.Roles {
					if role.Name == "user" {
						c.Next()
						c.Abort()
						return
					}
				}
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}

func AdminAuthMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwtHelper.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				for _, role := range decodedClaims.Roles {
					if role.Name == "admin" {
						c.Next()
						c.Abort()
						return
					}
				}
			}
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized!"})
		}
		c.Abort()
		return
	}
}
