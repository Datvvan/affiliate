package db

import (
	"context"

	"github.com/datvvan/affiliate/model"
	"github.com/go-pg/pg/v10"
)

type UserQuery interface {
	UserUpdateByID(ctx context.Context, data *model.User) error
}

type userQuery struct{}
type userTxQuery struct {
	tx *pg.Tx
}

func NewUserQuery(tx *pg.Tx) UserQuery {
	if tx == nil {
		return &userQuery{}
	} else {
		return &userTxQuery{
			tx: tx,
		}
	}

}

func UserFindOne(where string, params interface{}) (*model.User, error) {
	instance := GetInstance()
	data := &model.User{}
	if err := instance.DB.Model(data).Where(where, params).Select(); err != nil {
		return nil, err
	}
	return data, nil
}

func GetUserCommissionByUserID(userID string) (*model.UserSubscription, error) {
	instance := GetInstance()
	user := &model.User{}
	userSub := &model.UserSubscription{}

	err := instance.DB.Model(user).
		ColumnExpr(`"user"."id"`).
		ColumnExpr(`"user"."type"`).
		ColumnExpr(`"user"."email"`).
		ColumnExpr(`"user"."ref_code"`).
		ColumnExpr(`"user"."intermediary"`).
		ColumnExpr(`"user"."commission_amount"`).
		ColumnExpr(`"s"."id" as subscription_id`).
		ColumnExpr(`"s"."member_type"`).
		ColumnExpr(`"s"."last_paid_date"`).
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

func (u *userQuery) UserUpdateByID(ctx context.Context, data *model.User) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (u *userTxQuery) UserUpdateByID(ctx context.Context, data *model.User) error {
	if _, err := u.tx.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}
