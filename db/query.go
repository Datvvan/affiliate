package db

import (
	"context"

	"github.com/datvvan/affiliate/model"
	"github.com/go-pg/pg/v10"
)

type DbQuery interface {
	UserUpdateByID(ctx context.Context, data *model.User) error
	SubscriptionUpdateByID(ctx context.Context, data *model.Subscription) error
	AffiliateReferralUpdateByID(ctx context.Context, data *model.AffiliateReferrals) error
	AffiliateReferralInsertOne(ctx context.Context, data *model.AffiliateReferrals) error
	AffiliateReferralUpdateOneField(ctx context.Context, field string, params interface{}, where string, whereParams interface{}) error
	TransactionInsertOne(ctx context.Context, data *model.UserTransaction) error
	TransactionUpdateByID(ctx context.Context, data *model.UserTransaction) error
}

type dbQuery struct{}

type dbTxQuery struct {
	tx *pg.Tx
}

func NewDBQuery(tx *pg.Tx) DbQuery {
	if tx == nil {
		return &dbQuery{}
	} else {
		return &dbTxQuery{
			tx: tx,
		}
	}
}
