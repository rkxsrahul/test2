package routes

import (
	"os"

	"git.xenonstack.com/util/test-portal/src/jwt"
	"git.xenonstack.com/util/test-portal/src/services"
	"github.com/gin-gonic/gin"
)

func V1Routes(router *gin.Engine) {
	v1 := router.Group("/v1")
	//setting up middleware for protected apis
	authMiddleware := jwt.MwInitializer()
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/submittest", services.SubmitTest)
		v1.POST("/markandfetch", services.MarkAndFetch)
		v1.PUT("/increment-switch", services.IncrementSwitch)
		// admin apis
		admin := v1.Group("/admin")
		admin.Use(services.CheckAdmin)
		{
			if os.Getenv("TEST_PORTAL_ENVIRONMENT") != "production" {
				// toggle mail service
				admin.PUT("/mail/:value", services.ChangeMail)
			}
		}
	}
}
