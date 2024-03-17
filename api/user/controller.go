package user

import (
	"github.com/datvvan/affiliate/util"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	biz Business
}

func NewController() *Controller {
	return &Controller{
		biz: NewBiz(),
	}
}

func (controller *Controller) AddRefCode(c *gin.Context) {
	input := InputAddRefCode{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, util.ErrorResponse(400, err.Error()))
		return
	}
	resp, err := controller.biz.AddRefCode(c, input)
	if err != nil {
		c.JSON(400, util.ErrorResponse(400, err.Error()))
		return
	}
	c.JSON(200, util.SuccessResponse(resp))
}

func (controller *Controller) Subscription(c *gin.Context) {
	input := InputSubscription{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, util.ErrorResponse(400, err.Error()))
		return
	}
	if input.IsSubscribe {
		resp, err := controller.biz.Subscribe(c, input)
		if err != nil {
			c.JSON(400, util.ErrorResponse(400, err.Error()))
			return
		}
		c.JSON(200, util.SuccessResponse(resp))
		return
	} else {
		resp, err := controller.biz.Unsubscribe(c, input)
		if err != nil {
			c.JSON(400, util.ErrorResponse(400, err.Error()))
			return
		}
		c.JSON(200, util.SuccessResponse(resp))
		return
	}
}
