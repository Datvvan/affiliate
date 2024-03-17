package model

import "time"

type Subscription struct {
	ID             uint64             `json:"id"`
	UserID         string             `json:"user_id"`
	MemberType     SubscriptionStatus `json:"member_type"`
	EndOfTrialTime time.Time          `json:"end_of_trial_time"`
	LastPaidDate   time.Time          `json:"last_paid_date"`
	ExpiredTime    time.Time          `json:"expired_time"`
	UpdateAt       time.Time          `json:"update_at"`
}

type SubscriptionStatus string

const (
	TrialMember   SubscriptionStatus = "TRIAL_MEMBER"
	PaidMember    SubscriptionStatus = "PAID_MEMBER"
	NonPaidMember SubscriptionStatus = "NON_PAID_MEMBER"
)
