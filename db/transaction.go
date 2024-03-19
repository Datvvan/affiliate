package db

import (
	"context"
	"log"

	"github.com/datvvan/affiliate/model"
)

func CountAllCommissionCompleteTransactionByUserID(userID string) (int, error) {
	data := &model.UserTransaction{}
	count, err := GetInstance().DB.Model(data).Where("user_id = ?", userID).Where("type = ?", model.SubscriptionTrans).Where("status = ?", model.TransComplete).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (d *dbQuery) TransactionInsertOne(ctx context.Context, data *model.UserTransaction) error {
	if _, err := GetInstance().DB.Model(data).Returning("*").Insert(); err != nil {
		return err
	}
	return nil
}

func (d *dbTxQuery) TransactionInsertOne(ctx context.Context, data *model.UserTransaction) error {
	log.Println(data)
	if _, err := d.tx.Model(data).Returning("*").Insert(); err != nil {
		return err
	}
	return nil
}

func (d *dbQuery) TransactionUpdateByID(ctx context.Context, data *model.UserTransaction) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (d *dbTxQuery) TransactionUpdateByID(ctx context.Context, data *model.UserTransaction) error {
	if _, err := d.tx.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}
