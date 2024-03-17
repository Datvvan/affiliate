package model

import "time"

type AffiliateReferrals struct {
	ID               uint64           `json:"id"`
	Affiliate        string           `json:"affiliate"`
	Referral         string           `json:"referral"`
	IsConversion     bool             `json:"is_conversion"`
	IsCanceledSub    bool             `json:"is_canceled_sub"` //check user used to canceled subscriptions or not
	CommissionStatus CommissionStatus `json:"commission_status"`
	CreateAt         time.Time        `json:"create_at"`
	UpdateAt         time.Time        `json:"update_at"`
}

type CommissionStatus string

const (
	CommissionPending  CommissionStatus = "PENDING"
	CommissionApproved CommissionStatus = "APPROVED"
	CommissionReject   CommissionStatus = "REJECT"
)
