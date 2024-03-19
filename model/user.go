package model

import "time"

type (
	UserTypeEnum      string
	TransactionType   string
	TransactionStatus string
)

const (
	UserType  UserTypeEnum = "USER"
	AdminType UserTypeEnum = "ADMIN"
)

type User struct {
	ID           string       `json:"id"`
	Email        string       `json:"email"`
	Type         UserTypeEnum `json:"type"`
	RefCode      string       `json:"ref_code"`
	Intermediary string       `json:"intermediary"`
	CreateAt     time.Time    `json:"create_at"`
	UpdateAt     time.Time    `json:"update_at"`
}

type UserSubscription struct {
	ID             string             `json:"id"`
	Email          string             `json:"email"`
	RefCode        string             `json:"ref_code"`
	Type           UserTypeEnum       `json:"type"`
	EndOfTrialTime time.Time          `json:"end_of_trial_time"`
	Intermediary   string             `json:"intermediary"`
	SubscriptionID uint64             `json:"subscription_id" pg:"subscription_id"`
	TransactionID  uint64             `json:"transaction_id"`
	MemberType     SubscriptionStatus `json:"member_type" `
	ExpiredTime    time.Time          `json:"expired_time"`
	SubUpdateAt    time.Time          `json:"sub_update_at" pg:"sub_update_at"`
}

type UserAffiliate struct {
	ID               string           `json:"id"`
	AffiliateEmail   string           `json:"affiliate_email" pg:"affiliate_email"`
	Referral         string           `json:"referral"`
	IsConversion     bool             `json:"is_conversion"`
	IsCanceledSub    bool             `json:"is_canceled_sub"` //check user used to canceled subscriptions or not
	CommissionStatus CommissionStatus `json:"commission_status"`
	ReferralEmail    string           `json:"referral_email" pg:"referral_email"`
	EndOfTrialTime   time.Time        `json:"end_of_trial_time"`
}
