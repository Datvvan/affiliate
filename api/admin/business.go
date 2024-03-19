package admin

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/datvvan/affiliate/db"
	"github.com/datvvan/affiliate/model"
	"github.com/datvvan/affiliate/util"
	"github.com/go-pg/pg/v10"
)

type Business interface {
	ListConversion(context.Context, int, int, string, bool, bool) (interface{}, error)
	ApproveCommission(context.Context, InputReviewCommission) error
	RejectCommission(context.Context, InputReviewCommission) error
}

type business struct{}

func NewBiz() *business {
	return &business{}
}

func (b *business) ListConversion(ctx context.Context, offset int, limit int, emailQuery string, isConversion bool, isReachCommission bool) (interface{}, error) {
	data, err := db.GetAllAffiliateCommissionList(limit, offset, emailQuery, isConversion, isReachCommission)
	if err != nil {
		return nil, err
	}
	log.Println(data)
	return data, nil
}

func (b *business) ApproveCommission(ctx context.Context, input InputReviewCommission) error {

	//TODO: send batch payout -> received batchID
	//assume: batchID in input
	instance := db.GetInstance()
	return instance.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {

		for _, v := range input.AffiliateReferralsID {
			affiliateReferrals, err := db.AffiliateReferralFindOne("id = ?", v)
			if err != nil {
				return err
			}

			subscription, err := db.SubscriptionFindOne("user_id = ?", affiliateReferrals.Referral)
			if err != nil {
				return err
			}
			if subscription.EndOfTrialTime.Before(time.Now().Add(-time.Hour*util.ChallengeTime)) == false ||
				affiliateReferrals.CommissionStatus == model.CommissionApproved ||
				affiliateReferrals.IsCanceledSub == true ||
				affiliateReferrals.IsConversion == false {
				return errors.New("Have commission does not qualify")
			}

			transaction := &model.UserTransaction{
				UserID: affiliateReferrals.Affiliate,
				Type:   model.CommissionTrans,
				Status: model.TransProcessing,
				Amount: util.CommissionAmount,
			}

			err = db.NewDBQuery(tx).TransactionInsertOne(ctx, transaction)
			if err != nil {
				return err
			}

			affiliateReferrals.BatchID = input.BatchID
			affiliateReferrals.CommissionStatus = model.CommissionApproved
			affiliateReferrals.TransactionID = int(transaction.ID)

			err = db.NewDBQuery(tx).AffiliateReferralUpdateByID(ctx, affiliateReferrals)
			if err != nil {
				return err
			}
		}
		return nil
	})

}

func (b *business) RejectCommission(ctx context.Context, input InputReviewCommission) error {
	instance := db.GetInstance()
	return instance.DB.RunInTransaction(ctx, func(tx *pg.Tx) error {

		for _, v := range input.AffiliateReferralsID {
			affiliateReferrals, err := db.AffiliateReferralFindOne("id = ?", v)
			if err != nil {
				return err
			}

			affiliateReferrals.CommissionStatus = model.CommissionReject
			affiliateReferrals.IsConversion = false

			err = db.NewDBQuery(tx).AffiliateReferralUpdateByID(ctx, affiliateReferrals)
			if err != nil {
				return err
			}
		}
		return nil
	})
}
