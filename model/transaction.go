package model

import "time"

const (
	SubscriptionTrans TransactionType   = "SUBSCRIPTION"
	CommissionTrans   TransactionType   = "COMMISSION"
	TransComplete     TransactionStatus = "COMPLETE"
	TransProcessing   TransactionStatus = "PROCESSING"
	TransCancel       TransactionStatus = "CANCEL"
)

type UserTransaction struct {
	tableName struct{}          `pg:"user_transaction"`
	ID        uint64            `json:"id"`
	UserID    string            `json:"user_id"`
	Type      TransactionType   `json:"type"`
	Status    TransactionStatus `json:"status"`
	CreateAt  time.Time         `json:"create_at"`
	UpdateAt  time.Time         `json:"update_at"`
}
