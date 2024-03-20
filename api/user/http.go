package user

import "github.com/datvvan/affiliate/model"

type InputAddRefCode struct {
	UserId  string `json:"user_id"`
	RefCode string `json:"ref_code"`
}

type InputSubscription struct {
	UserID      string `json:"user_id"`
	IsSubscribe bool   `json:"is_subscribe"`
}

type OutputReferralList struct {
	TotalPendingAmount float32
	TotalPaidAmount    float32
	ReferralList       []model.ReferralList
}
