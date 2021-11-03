package jwt

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"

	"git.xenonstack.com/util/test-portal/config"
)

func MwInitializer() *jwt.GinJWTMiddleware {

	//===================================================================================

	// intializing the gin jwt middleware by setting new configuration values
	authMiddleware := &jwt.GinJWTMiddleware{
		// name to display to the user
		Realm: "test zone",
		// passing key string by converting into bytes
		Key: []byte(config.TestPortalKey),
		// Duration that a jwt token is valid
		Timeout: config.JWTExpireTime,
		// this field allows clients to refresh their token until MaxRefresh has passed
		MaxRefresh: time.Hour * 720,
		// own Unauthorized func.
		Unauthorized: unauthorizedFunc,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		TokenLookup: "header:Authorization",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc:         time.Now,
		SigningAlgorithm: "HS512",
	}
	// Initial middleware default setting.
	authMiddleware.MiddlewareInit()
	//===================================================================================

	return authMiddleware
}

func unauthorizedFunc(c *gin.Context, code int, msg string) {

	c.JSON(code, gin.H{"error": true, "code": code, "message": msg + " " + "Please Login Again"})

}
