package db

import (
	"context"

	"github.com/datvvan/affiliate/model"
)

func SubscriptionFindOne(where string, params interface{}) (*model.Subscription, error) {
	instance := GetInstance()
	data := &model.Subscription{}
	if err := instance.DB.Model(data).Where(where, params).Select(); err != nil {
		return nil, err
	}
	return data, nil
}

func (d *dbQuery) SubscriptionUpdateByID(ctx context.Context, data *model.Subscription) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (d *dbTxQuery) SubscriptionUpdateByID(ctx context.Context, data *model.Subscription) error {
	if _, err := d.tx.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}
