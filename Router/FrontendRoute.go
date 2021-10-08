package Router

import (
	"news-poolerr/Controller"

	"github.com/gin-gonic/gin"
)

func RegisterFrontendRoute(r *gin.Engine) {
	r.GET("/", Controller.HomePage)
	r.GET("/links/:link_id", Controller.SingleLinkpage)
}
