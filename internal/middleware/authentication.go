package middleware

import (
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/jwt"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/pkg/redis"
	"github.com/AhmetDenizGuner/Patika-Picus-Security-Golang-Backend-Bootcamp-Graduation-Project-AhmetDenizGuner/shared"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

// TODO check redis
func UserAuthMiddleware(secretKey string, redisClient *redis.RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwt.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				log.Println(decodedClaims)
				if !checkTokenValidInRedis(c.GetHeader("Authorization"), decodedClaims, redisClient) {
					c.JSON(http.StatusForbidden, shared.ApiErrorResponse{
						IsSuccess:    false,
						ErrorMessage: "You are not authorized!"})
					c.Abort()
					return
				}
				if decodedClaims.Role.Name == "USER" || decodedClaims.Role.Name == "ADMIN" {
					c.Next()
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusForbidden, shared.ApiErrorResponse{
				IsSuccess:    false,
				ErrorMessage: "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, shared.ApiErrorResponse{
				IsSuccess:    false,
				ErrorMessage: "You are not authorized!"})
		}
		c.Abort()
		return
	}
}

func AdminAuthMiddleware(secretKey string, redisClient *redis.RedisClient) gin.HandlerFunc {
	return func(c *gin.Context) {

		if c.GetHeader("Authorization") != "" {
			decodedClaims := jwt.VerifyToken(c.GetHeader("Authorization"), secretKey)
			if decodedClaims != nil {
				log.Println(decodedClaims)
				if !checkTokenValidInRedis(c.GetHeader("Authorization"), decodedClaims, redisClient) {
					c.JSON(http.StatusForbidden, shared.ApiErrorResponse{
						IsSuccess:    false,
						ErrorMessage: "You are not authorized!"})
					c.Abort()
					return
				}
				if decodedClaims.Role.Name == "ADMIN" {
					c.Next()
					c.Abort()
					return
				}
			}
			c.JSON(http.StatusForbidden, shared.ApiErrorResponse{
				IsSuccess:    false,
				ErrorMessage: "You are not allowed to use this endpoint!"})
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, shared.ApiErrorResponse{
				IsSuccess:    false,
				ErrorMessage: "You are not authorized!"})
		}
		c.Abort()
		return
	}
}

func checkTokenValidInRedis(token string, decodedClaims *jwt.DecodedToken, redisClient *redis.RedisClient) bool {
	var cachedToken string
	err := redisClient.GetKey(decodedClaims.Email, &cachedToken)
	if err != nil {
		return false
	}
	if strings.Compare(cachedToken, token) != 0 {
		return false
	}
	return true
}
