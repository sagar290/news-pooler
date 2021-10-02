package Router

import (
	"news-poolerr/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterApisRoute(r *gin.Engine) {
	url := r.Group("/api")

	{
		url.POST("links", Controller.AddLinks)
		url.GET("links", Controller.GetLinks)
		url.PATCH("links/:feed_id", Controller.UpdateLink)
		url.DELETE("links/:feed_id", Controller.DeleteLink)

	}

}
