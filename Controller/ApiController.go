package Controller

import (
	models "news-poolerr/Model"
	services "news-poolerr/Services"
	structs "news-poolerr/Structs"
	"strconv"

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

func GetLink(c *gin.Context) {
	link_id := c.Param("feed_id")
	var url models.Link

	// convert link_id to int
	id, err := strconv.Atoi(link_id)
	if err != nil {

		c.JSON(400, gin.H{
			"message": "url_id type is not int",
			"data":    nil,
		})

		return
	}

	url = services.GetLink(id)

	c.JSON(200, gin.H{
		"message": "",
		"data":    url,
	})
}

func DeleteLink(c *gin.Context) {

	feed_id := c.Param("feed_id")

	// convert feed_id to int
	id, err := strconv.Atoi(feed_id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "feed_id type is not int",
			"data":    nil,
		})
		return
	}

	url, message := services.DeleteLink(id)

	if !url {
		c.JSON(400, gin.H{
			"message": message,
			"data":    nil,
		})
	}

	c.JSON(200, gin.H{
		"message": message,
		"data":    nil,
	})
}
