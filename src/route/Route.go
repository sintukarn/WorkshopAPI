package route

import (
	"controller"
	"github.com/gin-gonic/gin"
	"time"
	"github.com/itsjamie/gin-cors"
)

func RunREST() {
	router := gin.Default()
	
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.POST("/createBU", func(c *gin.Context) {
		controller.CreateBU(c)
	})

	router.POST("/createSquad", func(c *gin.Context) {
		controller.CreateSquad(c)
	})

	router.POST("/createDev", func(c *gin.Context) {
		controller.CreateDev(c)
	})

	router.POST("/insertSquadToBU", func(c *gin.Context) {
		controller.InsertSquadToBU(c)
	})

	router.POST("/insertDevToSquad", func(c *gin.Context) {
		controller.InsertDevToSquad(c)
	})

	router.POST("/deactiveBU", func(c *gin.Context) {
		controller.DeactiveBU(c)
	})

	router.POST("/deactiveSquad", func(c *gin.Context) {
		controller.DeactiveSquad(c)
	})

	router.GET("/displayDev", func(c *gin.Context) {
		controller.DisplayDev(c)
	})

	router.GET("/displaySquad", func(c *gin.Context) {
		controller.DisplaySquad(c)
	})

	router.Run(":8080")
}
