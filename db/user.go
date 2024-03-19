package db

import (
	"context"
	"time"

	"github.com/datvvan/affiliate/model"
	"github.com/datvvan/affiliate/util"
	"github.com/go-pg/pg/v10/orm"
)

func UserFindOne(where string, params interface{}) (*model.User, error) {
	instance := GetInstance()
	data := &model.User{}
	if err := instance.DB.Model(data).Where(where, params).Select(); err != nil {
		return nil, err
	}
	return data, nil
}

func GetUserSubscriptionByUserID(userID string) (*model.UserSubscription, error) {
	instance := GetInstance()
	user := &model.User{}
	userSub := &model.UserSubscription{}

	err := instance.DB.Model(user).
		ColumnExpr(`"user"."id"`).
		ColumnExpr(`"user"."type"`).
		ColumnExpr(`"user"."email"`).
		ColumnExpr(`"user"."ref_code"`).
		ColumnExpr(`"user"."intermediary"`).
		ColumnExpr(`"s"."id" as subscription_id`).
		ColumnExpr(`"s"."member_type"`).
		ColumnExpr(`"s"."transaction_id"`).
		ColumnExpr(`"s"."end_of_trial_time"`).
		ColumnExpr(`"s"."expired_time"`).
		ColumnExpr(`"s"."update_at" as sub_update_at`).
		Join(`JOIN subscriptions as s ON "s"."user_id" = "user"."id"`).
		Where(`"user"."id" = ?`, userID).
		Select(userSub)
	if err != nil {
		return nil, err
	}
	return userSub, nil
}

func (d *dbQuery) UserUpdateByID(ctx context.Context, data *model.User) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (d *dbTxQuery) UserUpdateByID(ctx context.Context, data *model.User) error {
	if _, err := d.tx.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func GetAllAffiliateCommissionList(limit int, offset int, emailQuery string, isConversion bool, isReachCommission bool) ([]model.UserAffiliate, error) {
	if isReachCommission == true {
		data := model.User{}
		result := []model.UserAffiliate{}
		err := GetInstance().DB.Model(&data).
			ColumnExpr(`"user"."id"`).
			ColumnExpr(`"user"."email" as affiliate_email`).
			ColumnExpr(`"a"."referral"`).
			ColumnExpr(`"a"."is_conversion"`).
			ColumnExpr(`"a"."is_canceled_sub"`).
			ColumnExpr(`"a"."commission_status"`).
			ColumnExpr(`"u"."email" as referral_email`).
			ColumnExpr(`"s"."end_of_trial_time"`).
			Join(`JOIN affiliate_referrals as a ON "a"."affiliate" = "user"."id"`).
			Join(`JOIN users as u ON "u"."id"="a"."referral"`).
			Join(`JOIN subscriptions as s ON "s"."user_id" = "user"."id"`).
			Where("is_conversion", true).
			Where("end_of_trial_time < ?", time.Now().Add(-time.Hour*util.ChallengeTime)).
			WhereGroup(func(q *orm.Query) (*orm.Query, error) {
				q = q.Where(`"user"."email" LIKE ?`, "%"+emailQuery+"%")
				return q, nil
			}).
			Limit(limit).Offset(offset).Select(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else if isConversion == true {
		data := model.User{}
		result := []model.UserAffiliate{}
		err := GetInstance().DB.Model(&data).
			ColumnExpr(`"user"."id"`).
			ColumnExpr(`"user"."email" as affiliate_email`).
			ColumnExpr(`"a"."referral"`).
			ColumnExpr(`"a"."is_conversion"`).
			ColumnExpr(`"a"."is_canceled_sub"`).
			ColumnExpr(`"a"."commission_status"`).
			ColumnExpr(`"u"."email" as referral_email`).
			ColumnExpr(`"s"."end_of_trial_time"`).
			Join(`JOIN affiliate_referrals as a ON "a"."affiliate" = "user"."id"`).
			Join(`JOIN users as u ON "u"."id"="a"."referral"`).
			Join(`JOIN subscriptions as s ON "s"."user_id" = "user"."id"`).
			Where("is_conversion", true).
			WhereGroup(func(q *orm.Query) (*orm.Query, error) {
				q = q.Where(`"user"."email" LIKE ?`, "%"+emailQuery+"%")
				return q, nil
			}).
			Limit(limit).Offset(offset).Select(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		data := model.User{}
		result := []model.UserAffiliate{}
		err := GetInstance().DB.Model(&data).
			ColumnExpr(`"user"."id"`).
			ColumnExpr(`"user"."email" as affiliate_email`).
			ColumnExpr(`"a"."referral"`).
			ColumnExpr(`"a"."is_conversion"`).
			ColumnExpr(`"a"."is_canceled_sub"`).
			ColumnExpr(`"a"."commission_status"`).
			ColumnExpr(`"u"."email" as referral_email`).
			ColumnExpr(`"s"."end_of_trial_time"`).
			Join(`JOIN affiliate_referrals as a ON "a"."affiliate" = "user"."id"`).
			Join(`JOIN users as u ON "u"."id"="a"."referral"`).
			Join(`JOIN subscriptions as s ON "s"."user_id" = "user"."id"`).
			WhereGroup(func(q *orm.Query) (*orm.Query, error) {
				q = q.Where(`"user"."email" LIKE ?`, "%"+emailQuery+"%")
				return q, nil
			}).
			Limit(limit).Offset(offset).Select(&result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
}
