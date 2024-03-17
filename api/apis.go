package api

import (
	"github.com/datvvan/affiliate/api/user"
	"github.com/gin-gonic/gin"
)

func RegisterAPI(app *gin.Engine) {
	userAPIs(app)

}

func userAPIs(app *gin.Engine) {
	controller := user.NewController()
	group := app.Group("user")
	group.POST("/ref_code", controller.AddRefCode)
	group.POST("/subscription", controller.Subscription)
}
