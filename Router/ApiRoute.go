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
		url.GET("links/:link_id", Controller.GetLink)
		url.PATCH("links/:link_id", Controller.UpdateLink)
		url.DELETE("links/:link_id", Controller.DeleteLink)

	}

}
