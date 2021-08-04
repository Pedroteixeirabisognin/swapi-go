package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pedroteixeirabisognin/swapi-go/database"
	"github.com/pedroteixeirabisognin/swapi-go/model"
)

func ShowPlanet(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var planet model.Planet
	err = db.First(&planet, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot find book: " + err.Error(),
		})

		return
	}

	c.JSON(200, planet)
}

func CreatePlanet(c *gin.Context) {
	db := database.GetDatabase()

	var planet model.Planet

	err := c.ShouldBindJSON(&planet)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannor bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&planet).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot create book:" + err.Error(),
		})
		return
	}

	c.JSON(201, planet)

}

func ShowPlanets(c *gin.Context) {

	db := database.GetDatabase()

	var planets []model.Planet
	err := db.Find(&planets).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot list planet: " + err.Error(),
		})

		return
	}

	c.JSON(200, planets)
}
