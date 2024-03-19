package webhook

import (
	"context"
	"time"

	"github.com/datvvan/affiliate/db"
	"github.com/datvvan/affiliate/model"
	"github.com/datvvan/affiliate/util"
	"github.com/go-pg/pg/v10"
)

func subscribe(ctx context.Context, user *model.UserSubscription) error {
	instance := db.GetInstance()
	return instance.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {

		transaction := &model.UserTransaction{
			UserID: user.ID,
			Type:   model.SubscriptionTrans,
			Status: model.TransComplete,
		}
		if err := db.NewDBQuery(tx).TransactionInsertOne(ctx, transaction); err != nil {
			return err
		}

		subscription := &model.Subscription{
			ID:             user.SubscriptionID,
			UserID:         user.ID,
			MemberType:     user.MemberType,
			EndOfTrialTime: user.EndOfTrialTime,
			TransactionID:  transaction.ID,
			ExpiredTime:    user.ExpiredTime,
			UpdateAt:       user.SubUpdateAt,
		}

		if user.MemberType == model.TrialMember {
			subscription.ExpiredTime = newExpiredTimeTrialMemberAndPaidMember(subscription.ExpiredTime)
			if user.Intermediary != "" {
				affiliateReferrals, err := changeCommissionStatus(user.Intermediary, user.ID)
				if err != nil {
					return err
				}
				affiliateReferrals.IsConversion = true
				if err = db.NewDBQuery(tx).AffiliateReferralUpdateByID(ctx, affiliateReferrals); err != nil {
					return err
				}
			}

		} else if user.MemberType == model.PaidMember {
			subscription.ExpiredTime = newExpiredTimeTrialMemberAndPaidMember(subscription.ExpiredTime)

		} else {
			subscription.ExpiredTime = newExpiredTimeNonPaidMember(subscription.ExpiredTime)
			if user.Intermediary != "" {
				affiliateReferrals, err := changeCommissionStatus(user.Intermediary, user.ID)
				if err != nil {
					return err
				}
				affiliateReferrals.IsConversion = true
				subscription.EndOfTrialTime = time.Now()
				subscription.ExpiredTime = time.Now().Add(time.Hour * util.SubscriptionTime)
				if err = db.NewDBQuery(tx).AffiliateReferralUpdateByID(ctx, affiliateReferrals); err != nil {
					return err
				}

			}
		}

		subscription.MemberType = model.PaidMember
		if err := db.NewDBQuery(tx).SubscriptionUpdateByID(ctx, subscription); err != nil {
			return err
		}

		return nil
	})

}

func unsubscribe(ctx context.Context, user *model.UserSubscription) error {
	if user.Intermediary != "" {
		affiliateReferrals, err := db.AffiliateReferralFindOneWithRelationParams(user.Intermediary, user.ID)
		if err != nil {
			return err
		}

		if time.Now().Before(user.EndOfTrialTime.Add(time.Hour * util.ChallengeTime)) {
			affiliateReferrals.IsCanceledSub = true
			affiliateReferrals.CommissionStatus = model.CommissionReject
			affiliateReferrals.IsConversion = false
		}

		err = db.NewDBQuery(nil).AffiliateReferralUpdateByID(ctx, affiliateReferrals)
		if err != nil {
			return err
		}
		return nil
	}
	return nil
}

func changeCommissionStatus(affiliate string, referral string) (*model.AffiliateReferrals, error) {
	affiliateReferrals := &model.AffiliateReferrals{}
	affiliateReferrals, err := db.AffiliateReferralFindOneWithRelationParams(affiliate, referral)
	if err != nil {
		return nil, err
	}
	if !affiliateReferrals.IsCanceledSub && affiliateReferrals.CommissionStatus == "" {
		affiliateReferrals.CommissionStatus = model.CommissionPending
	}

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
