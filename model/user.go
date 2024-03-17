package model

import "time"

type UserTypeEnum string

const (
	UserType  UserTypeEnum = "USER"
	AdminType UserTypeEnum = "ADMIN"
)

type User struct {
	ID               string       `json:"id"`
	Email            string       `json:"email"`
	RefCode          string       `json:"ref_code"`
	Type             UserTypeEnum `json:"type"`
	CommissionAmount uint16       `json:"commission_amount"`
	Intermediary     string       `json:"intermediary"`
	CreateAt         time.Time    `json:"create_at"`
	UpdateAt         time.Time    `json:"update_at"`
}

type UserSubscription struct {
	ID               string             `json:"id"`
	Email            string             `json:"email"`
	RefCode          string             `json:"ref_code"`
	Type             UserTypeEnum       `json:"type"`
	CommissionAmount uint16             `json:"commission_amount"`
	Intermediary     string             `json:"intermediary"`
	SubscriptionID   uint64             `json:"subscription_id" pg:"subscription_id"`
	UserID           string             `json:"user_id"`
	MemberType       SubscriptionStatus `json:"member_type" `
	LastPaidDate     time.Time          `json:"last_paid_date"`
	ExpiredTime      time.Time          `json:"expired_time"`
	SubUpdateAt      time.Time          `json:"sub_update_at" pg:"sub_update_at"`
}
