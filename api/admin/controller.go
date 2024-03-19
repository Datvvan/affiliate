package admin

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

func (controller *Controller) GetAffiliateCommission(c *gin.Context) {
	input := c.Request.URL.Query()
	page, _ := strconv.Atoi(input["page"][0])
	limit, _ := strconv.Atoi(input["limit"][0])
	email := input["email"][0]
	isConversion, _ := strconv.ParseBool(input["isConversion"][0])
	isReachCommission, _ := strconv.ParseBool(input["isReachCommission"][0])
	offset, limit := util.GetPagination(page, limit)

	resp, err := controller.biz.ListConversion(c, offset, limit, email, isConversion, isReachCommission)
	if err != nil {
		c.JSON(400, util.ErrorResponse(400, err.Error()))
		return
	}
	c.JSON(200, util.SuccessResponse(resp))
	return
}

func (controller *Controller) ReviewCommission(c *gin.Context) {
	input := InputReviewCommission{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, util.ErrorResponse(400, err.Error()))
		return
	}
	if input.IsApprove {
		err := controller.biz.ApproveCommission(c, input)
		if err != nil {
			c.JSON(400, util.ErrorResponse(400, err.Error()))
			return
		}
		c.JSON(200, util.SuccessResponse("Commission Approval"))
		return
	} else {
		err := controller.biz.RejectCommission(c, input)
		if err != nil {
			c.JSON(400, util.ErrorResponse(400, err.Error()))
			return
		}
		c.JSON(200, util.SuccessResponse("Commission Rejected"))
		return
	}
}

//TODO: Api approve -> save batch_id to affiliate
//change status to approve
//transaction status -> PROCESSING

//TODO: Handle webhook commission -> check batch_ID -> change transfer complete
