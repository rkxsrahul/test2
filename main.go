package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"git.xenonstack.com/util/test-portal/config"
	"git.xenonstack.com/util/test-portal/health"
	"git.xenonstack.com/util/test-portal/src/logger"
	"git.xenonstack.com/util/test-portal/src/routes"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go.uber.org/zap"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// setup zap logger
	level, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		level = -1
	}
	err = logger.Init(level, os.Getenv("LOG_TYPE"), os.Getenv("LOG_ENVIRONMENT"))
	log.Println("zap logger", err)
	if err != nil {
		zap.S().Error(err)
		return
	}
	var ideal int
	idealStr := os.Getenv("IDEAL_CONNECTIONS")
	if idealStr == "" {
		ideal = 100
	} else {
		ideal, _ = strconv.Atoi(idealStr)
	}
	db, err := gorm.Open("postgres", config.DBConfig())
	if err != nil {
		zap.S().Error(err)
		return
	}
	// close db instance whenever whole work completed
	defer db.Close()
	db.DB().SetMaxIdleConns(ideal)

	db.DB().SetConnMaxLifetime(1 * time.Nanosecond)

	config.DB = db

	if os.Getenv("TEST_PORTAL_ENVIRONMENT") != "production" {
		// removing info file if any
		_ = os.Remove("info.txt")

		// creating and opening info.txt file for writting logs
		file, err := os.OpenFile("info.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		// changing default writer of gin to file and std output
		gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

		// setting output for logs this will writes all logs to file
		log.SetOutput(gin.DefaultWriter)
		// writing log to check all in working
		log.Print("Logging to a file in Go!")
	}

	router := gin.Default()
	//set zap logger as std logger
	router.Use(ginzap.Ginzap(logger.Log, time.RFC3339, true))

	//allowing CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	corsConfig.AddAllowMethods("DELETE")
	router.Use(cors.New(corsConfig))

	router.GET("/", func(c *gin.Context) {
		type finalPath struct {
			Method string
			Path   string
		}

		data := router.Routes()
		finalPaths := make([]finalPath, 0)

		for i := 0; i < len(data); i++ {
			finalPaths = append(finalPaths, finalPath{
				Path:   data[i].Path,
				Method: data[i].Method,
			})
		}
		c.JSON(200, gin.H{
			"routes": finalPaths,
		})
	})

	// creating healthz end point, logs end point and end end point
	router.GET("/healthz", apiHealthz)
	if os.Getenv("ENVIRONMENT") != "production" {
		router.GET("/logs", checkToken, readLogs)
		router.GET("/end", checkToken, readEnv)
	}
	routes.V1Routes(router)
	// run gin router on specific port
	router.Run(":" + config.Port)
}

// serving info file at browser
func readLogs(c *gin.Context) {
	http.ServeFile(c.Writer, c.Request, "info.txt")
}

//checking health
func apiHealthz(c *gin.Context) {
	now := time.Now().Unix()
	err := health.Healthz()
	if err != nil {
		c.JSON(500, gin.H{
			"error":         err,
			"response_time": time.Now().Unix() - now,
		})
		return
	}
	c.JSON(200, gin.H{
		"error":         false,
		"message":       "ok",
		"response_time": time.Now().Unix() - now,
	})
}

// read all environment variables set and pass to browser
func readEnv(c *gin.Context) {
	env := make([]string, 0)
	for _, pair := range os.Environ() {
		env = append(env, pair)
	}
	c.JSON(200, gin.H{
		"environments": env,
	})
}

// check header is set or not for secured api
func checkToken(c *gin.Context) {
	xt := c.Request.Header.Get("X-TOKEN")
	if xt != "xyz" {
		c.Abort()
		c.JSON(401, gin.H{"message": "You are not authorised."})
		return
	}
	c.Next()
}
