package route

import (
	"github.com/gin-gonic/gin"
	"controller"
)

func RunREST()  {
	router := gin.Default()

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