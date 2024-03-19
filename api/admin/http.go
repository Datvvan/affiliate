package admin

type InputReviewCommission struct {
	BatchID              string `json:"batch_id"`
	AffiliateReferralsID []int  `json:"affiliate_referral_id"`
	IsApprove            bool   `json:"is_approve"`
}
