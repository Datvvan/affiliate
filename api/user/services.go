package user

import (
	"context"
	"log"
	"time"

	"github.com/datvvan/affiliate/db"
	"github.com/datvvan/affiliate/model"
	"github.com/datvvan/affiliate/util"
	"github.com/go-pg/pg/v10"
)

func addRefCode(ctx context.Context, user *model.User, affiliateReferrals *model.AffiliateReferrals) error {
	instance := db.GetInstance()
	return instance.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {
		userQuery := db.NewUserQuery(tx)
		if err := userQuery.UserUpdateByID(ctx, user); err != nil {
			return err
		}

		affiliateQuery := db.NewAffiliateQuery(tx)
		if err := affiliateQuery.AffiliateReferralInsertOne(ctx, affiliateReferrals); err != nil {
			return err
		}
		return nil
	})
}

func subscribe(ctx context.Context, user *model.UserSubscription) error {
	instance := db.GetInstance()
	return instance.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {
		subscription := &model.Subscription{
			ID:           user.SubscriptionID,
			UserID:       user.ID,
			MemberType:   user.MemberType,
			LastPaidDate: user.LastPaidDate,
			ExpiredTime:  user.ExpiredTime,
			UpdateAt:     user.SubUpdateAt,
		}

		log.Println("===================")

		if user.MemberType == model.TrialMember {
			subscription.ExpiredTime = newExpiredTimeTrialMemberAndPaidMember(subscription.ExpiredTime)
			if user.Intermediary != "" {
				affiliateReferrals, err := changeCommissionStatus(user.Intermediary, user.ID)
				if err != nil {
					return err
				}
				affiliateReferrals.CommissionStatus = model.CommissionPending
				if err = db.NewAffiliateQuery(tx).AffiliateReferralUpdateByID(ctx, affiliateReferrals); err != nil {
					return err
				}

			}
		} else if user.MemberType == model.PaidMember {
			subscription.ExpiredTime = newExpiredTimeTrialMemberAndPaidMember(subscription.ExpiredTime)
		} else {
			subscription.ExpiredTime = newExpiredTimeNonPaidMember(subscription.ExpiredTime)
		}

		subscription.MemberType = model.PaidMember
		subscription.LastPaidDate = time.Now()
		if err := db.NewSubscriptionQuery(tx).SubscriptionUpdateByID(ctx, subscription); err != nil {
			return err
		}

		return nil
	})

}

func changeCommissionStatus(affiliate string, referral string) (*model.AffiliateReferrals, error) {
	affiliateReferrals := &model.AffiliateReferrals{}
	affiliateReferrals, err := db.AffiliateReferralFindOneWithRelationParams(affiliate, referral)
	if err != nil {
		return nil, err
	}
	affiliateReferrals.CommissionStatus = model.CommissionPending

	return affiliateReferrals, nil

}

func newExpiredTimeTrialMemberAndPaidMember(currentExpiredTime time.Time) time.Time {
	var newExpiredTime time.Time
	today := time.Now()
	if currentExpiredTime.Before(today) {
		newExpiredTime = today.Add(time.Hour * util.SubscriptionTime)
	} else {
		newExpiredTime = currentExpiredTime.Add(time.Hour * util.SubscriptionTime)
	}
	return newExpiredTime
}

func newExpiredTimeNonPaidMember(currentExpiredTime time.Time) time.Time {
	today := time.Now()
	return today.Add(time.Hour * util.SubscriptionTime)
}
