package user

import (
	"context"
	"errors"
	"time"

	"github.com/datvvan/affiliate/db"
	"github.com/datvvan/affiliate/model"
)

type Business interface {
	AddRefCode(context.Context, InputAddRefCode) (interface{}, error)
	Subscribe(context.Context, InputSubscription) (interface{}, error)
	Unsubscribe(context.Context, InputSubscription) (interface{}, error)
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
		ID:               user.ID,
		Email:            user.Email,
		RefCode:          user.RefCode,
		Type:             user.Type,
		CommissionAmount: user.CommissionAmount,
		Intermediary:     intermediary.ID,
		UpdateAt:         time.Now(),
		CreateAt:         user.CreateAt,
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

func (b *business) Subscribe(ctx context.Context, input InputSubscription) (interface{}, error) {
	userSub := &model.UserSubscription{}

	userSub, err := db.GetUserCommissionByUserID(input.UserID)
	if err != nil {
		return nil, err
	}

	if userSub.Type == model.AdminType {
		return nil, errors.New("your role is admin")
	}

	err = subscribe(ctx, userSub)

	return userSub, nil
}

func (b *business) Unsubscribe(ctx context.Context, input InputSubscription) (interface{}, error) {
	userSub := &model.UserSubscription{}

	userSub, err := db.GetUserCommissionByUserID(input.UserID)
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
