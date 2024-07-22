package router

import (
	"net/http"

	"github.com/islu/ASS0720/docs"
	"github.com/islu/ASS0720/internal/domain/common"
	"github.com/islu/ASS0720/internal/usecase"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func RegisterHandlers(app *usecase.Application) *gin.Engine {

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	if common.Local == app.Params.Environment {
		gin.SetMode(gin.DebugMode)
	}

	// Creates a router without any middleware by default
	r := gin.New()

	/*
		Global middleware
	*/
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/ping", "/healthz"},
	}))

	// Recovery middleware recovers from any panics and writes a 500 if there was one
	r.Use(gin.Recovery())

	// Add CORS middleware
	r.Use(CORSMiddleware())

	/*
		Handlers
	*/

	// Add swagger-ui
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	// Set docs info for swagger in local
	if common.Local == app.Params.Environment {
		// docs.SwaggerInfo.BasePath = "/api/v1"
		docs.SwaggerInfo.Host = "localhost:8080"
	}

	// Add ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Add health-check
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "is alive"})
	})

	// We mount all handlers under /api path
	api := r.Group("/api")
	v1 := api.Group("/v1")
	// Add user namespace
	userGroup := v1.Group("/user")
	{
		// Get user tasks status by address
		userGroup.GET("/tasks/:address", GetUserTaskStatus(app))
		// Get user points history for distributed tasks
		userGroup.GET("/tasks", GetUserPointsHistory(app))
	}

	// Other handlers

	// TODO: Require access control
	dashboardGroup := v1.Group("/dashboard")
	{
		dashboardGroup.POST("/uniswap-log", UpdateUniswapUSDCETHPairSwapLog(app))
	}

	return r
}
