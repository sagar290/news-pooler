package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {

	c.HTML(http.StatusOK, "pages/home", gin.H{
		"title": "home",
	})
}


type Result struct {
	Date string
}

func SingleLinkpage(c *gin.Context) {

	link_id := c.Param("link_id")

 	c.HTML(http.StatusOK, "pages/singleurl", gin.H{
		"title": "singleurl",
		"link_id": link_id,
	})
}
