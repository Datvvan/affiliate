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

// {
// 	"id": "WH-83C777576Y8332450-2L845887S3616745G",
// 	"create_time": "2015-09-07T12:46:46Z",
// 	"resource_type": "payouts",
// 	"event_type": "PAYMENT.PAYOUTSBATCH.SUCCESS",
// 	"summary": "Payouts batch completed successfully.",
// 	"resource": {
// 	  "batch_header": {
// 		"payout_batch_id": "CQGA9SFAU8WSN",
// 		"batch_status": "SUCCESS",
// 		"time_created": "2015-09-07T12:46:41Z",
// 		"time_completed": "2015-09-07T12:46:45Z",
// 		"sender_batch_header": {
// 		  "sender_batch_id": "REL1"
// 		},
// 		"amount": {
// 		  "currency": "CAD",
// 		  "value": "25.0"
// 		},
// 		"fees": {
// 		  "currency": "CAD",
// 		  "value": "0.5"
// 		},
// 		"payments": 1
// 	  },
// 	  "links": [
// 		{
// 		  "href": "https://api.paypal.com/v1/payments/payouts/CQGA9SFAU8WSN",
// 		  "rel": "self",
// 		  "method": "GET"
// 		}
// 	  ]
// 	},
// 	"links": [
// 	  {
// 		"href": "https://api.paypal.com/v1/notifications/webhooks-events/WH-83C777576Y8332450-2L845887S3616745G",
// 		"rel": "self",
// 		"method": "GET",
// 		"encType": "application/json"
// 	  },
// 	  {
// 		"href": "https://api.paypal.com/v1/notifications/webhooks-events/WH-83C777576Y8332450-2L845887S3616745G/resend",
// 		"rel": "resend",
// 		"method": "POST",
// 		"encType": "application/json"
// 	  }
// 	],
// 	"event_version": "1.0"
//   }
//denied

//   {
// 	"id": "WH-1WS771625W821303N-0FG05488C6485202W",
// 	"create_time": "2015-09-07T12:49:17Z",
// 	"resource_type": "payouts",
// 	"event_type": "PAYMENT.PAYOUTSBATCH.DENIED",
// 	"summary": "Payouts batch got denied.",
// 	"resource": {
// 	  "batch_header": {
// 		"payout_batch_id": "CHDK9ANSPBWQU",
// 		"batch_status": "DENIED",
// 		"time_created": "2015-09-07T12:49:17Z",
// 		"time_completed": "2015-09-07T12:49:17Z",
// 		"sender_batch_header": {
// 		  "sender_batch_id": "REL2"
// 		},
// 		"amount": {
// 		  "currency": "CAD",
// 		  "value": "25.0"
// 		},
// 		"fees": {
// 		  "currency": "CAD",
// 		  "value": "0.5"
// 		},
// 		"payments": 1
// 	  },
// 	  "links": [
// 		{
// 		  "href": "https://api.paypal.com/v1/payments/payouts/CHDK9ANSPBWQU",
// 		  "rel": "self",
// 		  "method": "GET"
// 		}
// 	  ]
// 	},
// 	"links": [
// 	  {
// 		"href": "https://api.paypal.com/v1/notifications/webhooks-events/WH-1WS771625W821303N-0FG05488C6485202W",
// 		"rel": "self",
// 		"method": "GET",
// 		"encType": "application/json"
// 	  },
// 	  {
// 		"href": "https://api.paypal.com/v1/notifications/webhooks-events/WH-1WS771625W821303N-0FG05488C6485202W/resend",
// 		"rel": "resend",
// 		"method": "POST",
// 		"encType": "application/json"
// 	  }
// 	],
// 	"event_version": "1.0"
//   }
