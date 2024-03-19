package db

import (
	"context"

	"github.com/datvvan/affiliate/model"
	"github.com/go-pg/pg/v10/orm"
)

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

func CountPendingCommissionEachAffiliate(affiliate string) (int, error) {
	data := &model.AffiliateReferrals{}
	count, err := GetInstance().DB.Model(data).Where("affiliate=?", affiliate).Where("is_conversion=?", true).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetAffiliateReferralList(affiliate string, limit int, offset int, emailQuery string, isConversion bool) ([]model.ReferralList, error) {
	data := model.AffiliateReferrals{}
	result := []model.ReferralList{}
	if isConversion == true {
		err := GetInstance().DB.Model(&data).
			ColumnExpr(`"affiliate_referrals"."id"`).
			ColumnExpr(`"affiliate_referrals"."affiliate"`).
			ColumnExpr(`"affiliate_referrals"."referral"`).
			ColumnExpr(`"affiliate_referrals"."is_canceled_sub"`).
			ColumnExpr(`"affiliate_referrals"."commission_status"`).
			ColumnExpr(`"u"."email"`).
			Join(`JOIN users as u ON "u"."id"="affiliate_referrals"."referral"`).
			Where("affiliate = ?", affiliate).Where("is_conversion", true).WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			q = q.Where("email LIKE ?", "%"+emailQuery+"%")
			return q, nil
		}).Limit(limit).Offset(offset).Select(&result)
		if err != nil {
			return nil, err
		}
	} else {
		err := GetInstance().DB.Model(&data).
			ColumnExpr(`"affiliate_referrals"."id"`).
			ColumnExpr(`"affiliate_referrals"."affiliate"`).
			ColumnExpr(`"affiliate_referrals"."referral"`).
			ColumnExpr(`"affiliate_referrals"."is_canceled_sub"`).
			ColumnExpr(`"affiliate_referrals"."commission_status"`).
			ColumnExpr(`"u"."email"`).
			Join(`JOIN users as u ON "u"."id"="affiliate_referrals"."referral"`).
			Where("affiliate = ?", affiliate).WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			q = q.Where("email LIKE ?", "%"+emailQuery+"%")
			return q, nil
		}).Limit(limit).Offset(offset).Select(&result)
		if err != nil {
			return nil, err
		}
	}
	return result, nil
}

func GetAffiliateTransaction(batchID string) ([]model.AffiliateTransaction, error) {
	data := model.AffiliateReferrals{}
	result := []model.AffiliateTransaction{}
	err := GetInstance().DB.Model(&data).
		ColumnExpr(`"affiliate_referrals"."id"`).
		ColumnExpr(`"affiliate_referrals"."affiliate"`).
		ColumnExpr(`"affiliate_referrals"."referral"`).
		ColumnExpr(`"affiliate_referrals"."is_canceled_sub"`).
		ColumnExpr(`"affiliate_referrals"."commission_status"`).
		ColumnExpr(`"affiliate_referrals"."is_conversion"`).
		ColumnExpr(`"affiliate_referrals"."transaction_id"`).
		ColumnExpr(`"affiliate_referrals"."batch_id" `).
		ColumnExpr(`"t"."type" as transaction_type`).
		ColumnExpr(`"t"."status" as transaction_status`).
		Join(`JOIN user_transaction as t ON "t"."id"="affiliate_referrals"."transaction_id"`).
		Where("batch_id = ?", batchID).Select(&result)

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (d *dbQuery) AffiliateReferralInsertOne(ctx context.Context, data *model.AffiliateReferrals) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).Returning("*").Insert(); err != nil {
		return err
	}
	return nil
}

func (d *dbQuery) AffiliateReferralUpdateByID(ctx context.Context, data *model.AffiliateReferrals) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (d *dbQuery) AffiliateReferralUpdateOneField(ctx context.Context, field string, params interface{}, where string, whereParams interface{}) error {
	instance := GetInstance()
	data := &model.AffiliateReferrals{}
	if _, err := instance.DB.Model(data).Set(field, params).Where(where, whereParams).Update(); err != nil {
		return err
	}
	return nil
}

func (d *dbTxQuery) AffiliateReferralUpdateOneField(ctx context.Context, field string, params interface{}, where string, whereParams interface{}) error {
	data := &model.AffiliateReferrals{}
	if _, err := d.tx.Model(data).Set(field, params).Where(where, whereParams).Update(); err != nil {
		return err
	}
	return nil
}

func (d *dbTxQuery) AffiliateReferralInsertOne(ctx context.Context, data *model.AffiliateReferrals) error {
	if _, err := d.tx.Model(data).Returning("*").Insert(); err != nil {
		return err
	}
	return nil
}

func (d *dbTxQuery) AffiliateReferralUpdateByID(ctx context.Context, data *model.AffiliateReferrals) error {
	if _, err := d.tx.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}
