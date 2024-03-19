package webhook

import "github.com/datvvan/affiliate/util"

type BodyCheckoutWebhook struct {
	ID           string           `json:"id"`
	CreateTime   string           `json:"create_time"`
	ReSourceType string           `json:"checkout-order"`
	EventType    util.PaypalEvent `json:"event_type"`
	Summary      string           `json:"summary"`
	Resource     Resource         `json:"resource"`
}

type Resource struct {
	UpdateTime  string      `json:"update_time"`
	CreateTime  string      `json:"create_time"`
	ID          string      `json:"id"`
	GrossAmount GrossAmount `json:"gross_amount"`
	Intent      string      `json:"intent"`
	Payer       Payer       `json:"payer"`
	Status      string      `json:"status"`
}

type GrossAmount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}

type Payer struct {
	Name         PayerName `json:"name"`
	EmailAddress string    `json:"email_address"`
	PayerID      string    `json:"payer_id"`
}

type PayerName struct {
	GivenName string `json:"given_name"`
	SurName   string `json:"surname"`
}

type BodyPayoutWebhook struct {
	ID           string           `json:"id"`
	CreateTime   string           `json:"create_time"`
	ReSourceType string           `json:"checkout-order"`
	EventType    util.PaypalEvent `json:"event_type"`
	Summary      string           `json:"summary"`
	Resource     PayoutResource   `json:"resource"`
}

type PayoutResource struct {
	BatchHeader BatchHeader `json:"batch_header"`
}

type BatchHeader struct {
	PayoutBatchID string `json:"payout_batch_id"`
	BatchStatus   string `json:"batch_status"`
}
