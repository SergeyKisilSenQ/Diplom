package repo_billing

import (
	"log"
	"os"
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

func NewStorageBilling() *BillingData {
	return &BillingData{}
}

func (*BillingData) ReadFileBilling() *BillingData {
	r, err := os.ReadFile(os.Getenv("BILLING_FILE"))
	if err != nil {
		log.Fatal(err)
	}
	var b []bool
	for i := 0; i < len(r); i++ {
		if r[i] == '1' {
			b = append(b, true)
		} else {
			b = append(b, false)
		}
	}

	return &BillingData{
		CreateCustomer: b[5],
		Purchase:       b[4],
		Payout:         b[3],
		Recurring:      b[2],
		FraudControl:   b[1],
		CheckoutPage:   b[0],
	}
}
