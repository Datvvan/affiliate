package webhook

import (
	"fmt"

	"github.com/datvvan/affiliate/util"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Controller struct {
	biz Business
}

func NewController() *Controller {
	return &Controller{
		biz: NewBiz(),
	}
}

func (controller *Controller) Subscription(c *gin.Context) {

	input := &BodyCheckoutWebhook{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		log.Error("Have Error When listen Checkout Event: ", err)
		return
	}
	log.Info(fmt.Sprintf("Listen checkout webhook Paypal with email: %v", input.Resource.Payer.EmailAddress))

	if input.EventType == util.CheckoutOrderComplete {
		err := controller.biz.SubscriptionPaymentComplete(c, input.Resource.Payer.EmailAddress)
		if err != nil {
			log.Error("Have Error When listen Checkout Event: ", err)
		}
	}
	return
}

func (controller *Controller) PayoutWebhook(c *gin.Context) {
	input := &BodyPayoutWebhook{}
	err := c.ShouldBindJSON(input)
	if err != nil {
		log.Error("Have Error When listen Commission Payout Event: ", err)
		return
	}
	log.Info(fmt.Sprintf("Listen payout webhook Paypal with batch: %v", input.Resource.BatchHeader.PayoutBatchID))

	if input.EventType == util.PayoutBatchSuccess {
		err := controller.biz.BatchPayoutComplete(c, input.Resource.BatchHeader.PayoutBatchID)
		if err != nil {
			log.Error("Have Error When listen Commission Payout Success Event: ", err)
		}
		return

	} else if input.EventType == util.PayoutBatchDenied {
		err := controller.biz.BatchPayoutDenied(c, input.Resource.BatchHeader.PayoutBatchID)
		if err != nil {
			log.Error("Have Error When listen Commission Payout Denied Event: ", err)
		}
		return
	}

	return
}
