package db

import (
	"context"

	"github.com/datvvan/affiliate/model"
	"github.com/go-pg/pg/v10"
)

type SubscriptionQuery interface {
	SubscriptionUpdateByID(ctx context.Context, data *model.Subscription) error
}

type subscriptionQuery struct {
}

type subscriptionTxQuery struct {
	tx *pg.Tx
}

func NewSubscriptionQuery(tx *pg.Tx) SubscriptionQuery {
	if tx == nil {
		return &subscriptionQuery{}
	} else {
		return &subscriptionTxQuery{
			tx: tx,
		}
	}
}

func SubscriptionFindOne(where string, params interface{}) (*model.Subscription, error) {
	instance := GetInstance()
	data := &model.Subscription{}
	if err := instance.DB.Model(data).Where(where, params).Select(); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *subscriptionQuery) SubscriptionUpdateByID(ctx context.Context, data *model.Subscription) error {
	instance := GetInstance()
	if _, err := instance.DB.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}

func (s *subscriptionTxQuery) SubscriptionUpdateByID(ctx context.Context, data *model.Subscription) error {
	if _, err := s.tx.Model(data).WherePK().Update(); err != nil {
		return err
	}
	return nil
}
