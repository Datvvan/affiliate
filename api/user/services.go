package user

import (
	"context"
	"time"

	"github.com/datvvan/affiliate/db"
	"github.com/datvvan/affiliate/model"
	"github.com/datvvan/affiliate/util"
	"github.com/go-pg/pg/v10"
)

func addRefCode(ctx context.Context, user *model.User, affiliateReferrals *model.AffiliateReferrals) error {
	instance := db.GetInstance()
	return instance.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {
		userQuery := db.NewDBQuery(tx)
		if err := userQuery.UserUpdateByID(ctx, user); err != nil {
			return err
		}

		affiliateQuery := db.NewDBQuery(tx)
		if err := affiliateQuery.AffiliateReferralInsertOne(ctx, affiliateReferrals); err != nil {
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
