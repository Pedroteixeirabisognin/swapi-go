package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pedroteixeirabisognin/swapi-go/controller"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		planets := main.Group("planets")
		{
			planets.GET("/:id", controller.ShowPlanet)
			planets.GET("/", controller.ShowPlanets)
			planets.POST("/", controller.CreatePlanet)
		}
	}
	return router
}
