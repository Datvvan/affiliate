package model

import "time"

type CommissionStatus string

const (
	CommissionPending  CommissionStatus = "PENDING"
	CommissionApproved CommissionStatus = "APPROVED"
	CommissionReject   CommissionStatus = "REJECT"
)

type AffiliateReferrals struct {
	ID               uint64           `json:"id"`
	Affiliate        string           `json:"affiliate"`
	Referral         string           `json:"referral"`
	IsConversion     bool             `json:"is_conversion" pg:",use_zero"`
	IsCanceledSub    bool             `json:"is_canceled_sub" pg:",use_zero"` //check user used to canceled subscriptions or not
	CommissionStatus CommissionStatus `json:"commission_status"`
	TransactionID    int              `json:"transaction_id"`
	BatchID          string           `json:"batch_id"`
	CreateAt         time.Time        `json:"create_at"`
	UpdateAt         time.Time        `json:"update_at"`
}

type ReferralList struct {
	ID               uint64    `json:"id"`
	Affiliate        string    `json:"affiliate"`
	Referral         string    `json:"referral"`
	IsConversion     bool      `json:"is_conversion"`
	CommissionStatus string    `json:"commission_status"`
	IsCanceledSub    bool      `json:"is_canceled_sub"`
	CreateAt         time.Time `json:"create_at"`
	UpdateAt         time.Time `json:"update_at"`
	Email            string    `json:"email"`
}

type AffiliateTransaction struct {
	ID                uint64            `json:"id"`
	Affiliate         string            `json:"affiliate"`
	Referral          string            `json:"referral"`
	IsConversion      bool              `json:"is_conversion" pg:",use_zero"`
	IsCanceledSub     bool              `json:"is_canceled_sub" pg:",use_zero"` //check user used to canceled subscriptions or not
	CommissionStatus  CommissionStatus  `json:"commission_status"`
	TransactionID     int               `json:"transaction_id"`
	BatchID           string            `json:"batch_id"`
	TransactionType   TransactionType   `json:"transaction_type" pg:"transaction_type"`
	TransactionStatus TransactionStatus `json:"transaction_status" pg:"transaction_status"`
}
