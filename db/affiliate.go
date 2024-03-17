package db

import (
	"context"

	"github.com/datvvan/affiliate/model"
	"github.com/go-pg/pg/v10"
)

type AffiliateQuery interface {
	AffiliateReferralUpdateByID(ctx context.Context, data *model.AffiliateReferrals) error
	AffiliateReferralInsertOne(ctx context.Context, data *model.AffiliateReferrals) error
	AffiliateReferralUpdateOneField(ctx context.Context, field string, params interface{}, where string, whereParams interface{}) error
}

type affiliateQuery struct{}
type affiliateTxQuery struct {
	tx *pg.Tx
}

func NewAffiliateQuery(tx *pg.Tx) AffiliateQuery {
	if tx == nil {
		return &affiliateQuery{}
	} else {
		return &affiliateTxQuery{
			tx: tx,
		}
	}
}

func AffiliateReferralFindOne(where string, params interface{}) (*model.AffiliateReferrals, error) {
	instance := GetInstance()
	data := &model.AffiliateReferrals{}
	if err := instance.DB.Model(data).Where(where, params).Select(); err != nil {
		return nil, err
	}
	return data, nil
}

func AffiliateReferralFindOneWithRelationParams(affiliate string, referral string) (*model.AffiliateReferrals, error) {
	instance := GetInstance()
	data := &model.AffiliateReferrals{}
	if err := instance.DB.Model(data).Where("affiliate = ?", affiliate).Where("referral = ?", referral).First(); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *affiliateQuery) AffiliateReferralInsertOne(ctx context.Context, data *model.AffiliateReferrals) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).Returning("*").Insert(); err != nil {
		return err
	}
	return nil
}

func (a *affiliateQuery) AffiliateReferralUpdateByID(ctx context.Context, data *model.AffiliateReferrals) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (a *affiliateQuery) AffiliateReferralUpdateOneField(ctx context.Context, field string, params interface{}, where string, whereParams interface{}) error {
	instance := GetInstance()
	data := &model.AffiliateReferrals{}
	if _, err := instance.DB.Model(data).Set(field, params).Where(where, whereParams).Update(); err != nil {
		return err
	}
	return nil
}

func (a *affiliateTxQuery) AffiliateReferralUpdateOneField(ctx context.Context, field string, params interface{}, where string, whereParams interface{}) error {
	data := &model.AffiliateReferrals{}
	if _, err := a.tx.Model(data).Set(field, params).Where(where, whereParams).Update(); err != nil {
		return err
	}
	return nil
}

func (a *affiliateTxQuery) AffiliateReferralInsertOne(ctx context.Context, data *model.AffiliateReferrals) error {
	if _, err := a.tx.Model(data).Returning("*").Insert(); err != nil {
		return err
	}
	return nil
}

func (a *affiliateTxQuery) AffiliateReferralUpdateByID(ctx context.Context, data *model.AffiliateReferrals) error {
	if _, err := a.tx.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}
