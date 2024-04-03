package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/niazlv/subscribe-middleware-go/internal/api/managment"
	subscribeproxy "github.com/niazlv/subscribe-middleware-go/internal/api/subscribe_proxy"
)

func Setup(app *gin.Engine) {
	api := app.Group("")

	subscribeproxy.Setup(api)
	managment.Setup(api)

}
