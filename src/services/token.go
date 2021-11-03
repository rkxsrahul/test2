package services

import (
	"git.xenonstack.com/util/test-portal/config"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// ChangeMail is an api handler to toggle mail service
func ChangeMail(c *gin.Context) {
	config.DisableMail = (c.Param("value"))
}

// CheckAdmin is a middleware for checking user is admin or not
func CheckAdmin(c *gin.Context) {
	// extracting jwt claims
	claims := jwt.ExtractClaims(c)
	// checking sys role
	if claims["sys_role"].(string) != "admin" {
		c.Abort()
		c.JSON(401, gin.H{
			"error":   true,
			"message": "You are not authorized",
		})
		return
	}
	c.Next()
}
