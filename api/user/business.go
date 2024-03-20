package user

import (
	"context"
	"errors"
	"time"

	"github.com/datvvan/affiliate/db"
	"github.com/datvvan/affiliate/model"
	"github.com/datvvan/affiliate/util"
)

type Business interface {
	AddRefCode(context.Context, InputAddRefCode) (interface{}, error)
	Subscribe(context.Context) (interface{}, error)
	Unsubscribe(context.Context, InputSubscription) (interface{}, error)
	ReferralList(context.Context, string, int, int, string, bool) (interface{}, error)
}

type business struct{}

func NewBiz() Business {
	return &business{}
}

func (b *business) AddRefCode(ctx context.Context, input InputAddRefCode) (interface{}, error) {
	user := &model.User{}
	intermediary := &model.User{}

	user, err := db.UserFindOne("id =?", input.UserId)
	if err != nil {
		return nil, err
	}

	if user.Type == model.AdminType {
		return nil, errors.New("your role is admin")
	}

	if user.Intermediary != "" || user.RefCode == input.RefCode {
		return nil, errors.New("Can not add ref code")
	}

	intermediary, err = db.UserFindOne("ref_code = ?", input.RefCode)
	if err != nil {
		return nil, err
	}

	updateUser := &model.User{
		ID:           user.ID,
		Email:        user.Email,
		RefCode:      user.RefCode,
		Type:         user.Type,
		Intermediary: intermediary.ID,
		UpdateAt:     time.Now(),
		CreateAt:     user.CreateAt,
	}

	affiliateReferrals := &model.AffiliateReferrals{
		Affiliate: intermediary.ID,
		Referral:  user.ID,
	}

	if err := addRefCode(ctx, updateUser, affiliateReferrals); err != nil {
		return nil, err
	}

	return updateUser, nil
}

func (b *business) Subscribe(ctx context.Context) (interface{}, error) {

	//TODO: Handle payment process

	return "subscribe step complete, we are checking the payment", nil
}

func (b *business) Unsubscribe(ctx context.Context, input InputSubscription) (interface{}, error) {
	userSub := &model.UserSubscription{}

	userSub, err := db.GetUserSubscriptionByUserID(input.UserID)
	if err != nil {
		return nil, err
	}

	if userSub.Type == model.AdminType {
		return nil, errors.New("your role is admin")
	}

	if userSub.MemberType != model.PaidMember {
		return nil, errors.New("you are not paid member")
	}

	err = unsubscribe(ctx, userSub)
	return userSub, nil
}

func (b *business) ReferralList(ctx context.Context, userID string, page int, limit int, emailQuery string, isConversion bool) (interface{}, error) {
	pendingCommission, err := db.CountPendingCommissionEachAffiliate(userID)
	if err != nil {
		return nil, err
	}

	sumData, err := db.TotalPaidCommissionCompleteTransactionByUserId(userID)
	if err != nil {
		return nil, err
	}

	affiliateReferralsList, err := db.GetAffiliateReferralList(userID, limit, page, emailQuery, isConversion)
	if err != nil {
		return nil, err
	}

	response := &OutputReferralList{
		TotalPendingAmount: float32(pendingCommission * util.CommissionAmount),
		TotalPaidAmount:    sumData.Sum,
		ReferralList:       affiliateReferralsList,
	}
	return response, nil
}
