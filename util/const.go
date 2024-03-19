package util

const (
	SubscriptionTime = 24 * 30 //hour
	ChallengeTime    = 14 * 24
)

const CommissionAmount = 19 //$

type PaypalEvent string

const (
	PayoutBatchSuccess    PaypalEvent = "PAYMENT.PAYOUTSBATCH.SUCCESS"
	PayoutBatchDenied     PaypalEvent = "PAYMENT.PAYOUTSBATCH.DENIED"
	CheckoutOrderComplete PaypalEvent = "CHECKOUT.ORDER.COMPLETED"
)
