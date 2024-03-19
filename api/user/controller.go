package user

import (
	"strconv"

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
		resp, err := controller.biz.Subscribe(c)
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

func (controller *Controller) GetReferralList(c *gin.Context) {
	id := c.Param("id")
	input := c.Request.URL.Query()
	page, _ := strconv.Atoi(input["page"][0])
	limit, _ := strconv.Atoi(input["limit"][0])
	email := input["email"][0]
	isConversion, _ := strconv.ParseBool(input["isConversion"][0])
	offset, limit := util.GetPagination(page, limit)

	resp, err := controller.biz.ReferralList(c, id, offset, limit, email, isConversion)
	if err != nil {
		c.JSON(400, util.ErrorResponse(400, err.Error()))
		return
	}
	c.JSON(200, util.SuccessResponse(resp))
	return
}
