package api

import (
	"github.com/datvvan/affiliate/api/admin"
	"github.com/datvvan/affiliate/api/user"
	"github.com/datvvan/affiliate/api/webhook"
	"github.com/gin-gonic/gin"
)

func RegisterAPI(app *gin.Engine) {
	userAPIs(app)
	adminAPIs(app)
	webhookAPIs(app)

}

func userAPIs(app *gin.Engine) {
	controller := user.NewController()
	group := app.Group("user")
	group.POST("/ref-code", controller.AddRefCode)
	group.POST("/subscription", controller.Subscription)
	group.GET("/:id/referral-list", controller.GetReferralList)

}

func adminAPIs(app *gin.Engine) {
	controller := admin.NewController()
	group := app.Group("admin")
	group.GET("/affiliate-list", controller.GetAffiliateCommission)
	group.POST("/commission-review", controller.ReviewCommission)
}

func webhookAPIs(app *gin.Engine) {
	controller := webhook.NewController()
	group := app.Group("webhook")
	group.POST("/subscription", controller.Subscription)
	group.POST("/commission-payout", controller.PayoutWebhook)

}
