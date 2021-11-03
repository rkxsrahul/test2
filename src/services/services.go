package services

import (
	"log"
	"net/http"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"git.xenonstack.com/util/test-portal/src/bodytypes"
	"git.xenonstack.com/util/test-portal/src/test"
)

func SubmitTest(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	mapd := claims["claim"].(map[string]interface{})
	newOverview, err := test.SubmitTest(mapd)
	if err != nil {
		zap.S().Error(err)
		c.JSON(500, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": newOverview,
	})
}

func MarkAndFetch(ctx *gin.Context) {

	var data bodytypes.New
	if err := ctx.BindJSON(&data); err != nil {
		// if there is some error passing bad status code
		ctx.JSON(400, gin.H{"error": true, "message": "Please pass the required fields."})
		return
	}

	//extracting jwt claims
	claims := jwt.ExtractClaims(ctx)
	mapd := claims["claim"].(map[string]interface{})

	ch := make(chan test.Response)

	go test.Markandfetch(mapd, data, ch)

	var resp test.Response
	for i := 0; i < 1; i++ {
		resp = <-ch
		if resp.Error != nil {
			zap.S().Error(resp.Error)
			log.Println(resp.Error.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error":   "true",
				"message": resp.Error.Error(),
			})
			return
		}
		if resp.Question.Title != "" {
			ctx.JSON(200, gin.H{
				"error":    "false",
				"question": resp.Question,
			})
		} else {
			ctx.JSON(200, gin.H{
				"error":    "false",
				"message":  "Test submitted successfully",
				"question": resp.Question,
			})
		}
	}
}

func IncrementSwitch(ctx *gin.Context) {
	//extracting jwt claims
	claims := jwt.ExtractClaims(ctx)
	mapd := claims["claim"].(map[string]interface{})
	code, err := test.BrowserSwitch(mapd)
	if err != nil {
		zap.S().Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   false,
			"message": err,
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"error":   err == nil,
			"message": code,
		})
		return
	}
}
