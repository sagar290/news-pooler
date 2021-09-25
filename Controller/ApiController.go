package Controller

import (
	models "news-poolerr/Model"
	services "news-poolerr/Services"
	structs "news-poolerr/Structs"

	"github.com/gin-gonic/gin"
)

func AddLinks(c *gin.Context) {

	var linkBody structs.LinkBody

	if err := c.ShouldBindJSON(&linkBody); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
		return
	}

	results, message := services.StoreUrl(linkBody)

	if !results {
		c.JSON(400, gin.H{
			"message": message,
			"data":    []gin.H{},
		})
	}

	c.JSON(200, gin.H{
		"message": message,
		"data":    []gin.H{},
	})
}

func GetLinks(c *gin.Context) {

	var links []models.Link

	links = services.GetLinks()

	c.JSON(200, gin.H{
		"message": "",
		"data":    links,
	})
}
