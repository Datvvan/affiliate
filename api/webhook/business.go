package webhook

import (
	"context"
	"errors"
	"time"

	"github.com/datvvan/affiliate/db"
	"github.com/datvvan/affiliate/model"
	"github.com/go-pg/pg/v10"
)

type Business interface {
	SubscriptionPaymentComplete(context.Context, string) error
	BatchPayoutComplete(context.Context, string) error
	BatchPayoutDenied(context.Context, string) error
}

type business struct{}

func NewBiz() Business {
	return &business{}
}

func (b *business) SubscriptionPaymentComplete(ctx context.Context, emailAddress string) error {
	user := &model.User{}

	user, err := db.UserFindOne("email = ?", emailAddress)
	if err != nil {
		return errors.New("User not found")
	}

	userSub := &model.UserSubscription{}
	userSub, err = db.GetUserSubscriptionByUserID(user.ID)

	err = subscribe(ctx, userSub)
	if err != nil {
		return err
	}

	return nil
}

func (b *business) BatchPayoutComplete(ctx context.Context, batchID string) error {
	data, err := db.GetAffiliateTransaction(batchID)
	if err != nil {
		return err
	}
	return db.GetInstance().DB.RunInTransaction(ctx, func(tx *pg.Tx) error {
		for _, v := range data {
			transaction := &model.UserTransaction{
				ID:       uint64(v.TransactionID),
				UserID:   v.Affiliate,
				Type:     v.TransactionType,
				Status:   model.TransComplete,
				Amount:   v.Amount,
				UpdateAt: time.Now(),
				CreateAt: v.CreateAt,
			}
			err := db.NewDBQuery(tx).TransactionUpdateByID(ctx, transaction)
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (b *business) BatchPayoutDenied(ctx context.Context, batchID string) error {
	data, err := db.GetAffiliateTransaction(batchID)
	if err != nil {
		return err
	}
	return db.GetInstance().DB.RunInTransaction(ctx, func(tx *pg.Tx) error {
		for _, v := range data {
			transaction := &model.UserTransaction{
				ID:       uint64(v.TransactionID),
				UserID:   v.Affiliate,
				Type:     v.TransactionType,
				Amount:   v.Amount,
				Status:   model.TransCancel,
				UpdateAt: time.Now(),
				CreateAt: v.CreateAt,
			}
			err := db.NewDBQuery(tx).TransactionUpdateByID(ctx, transaction)
			if err != nil {
				return err
			}

			affiliateReferrals := &model.AffiliateReferrals{
				ID:               v.ID,
				Affiliate:        v.Affiliate,
				Referral:         v.Referral,
				IsConversion:     v.IsConversion,
				IsCanceledSub:    v.IsCanceledSub,
				CommissionStatus: model.CommissionPending,
				TransactionID:    0,
				BatchID:          "",
				UpdateAt:         time.Now(),
				CreateAt:         v.AffiliateCreateAt,
			}

			err = db.NewDBQuery(tx).AffiliateReferralUpdateByID(ctx, affiliateReferrals)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
