package repo_billing

import (
	"io/ioutil"
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

type StorageBilling map[int]*BillingData

func NewStorageBilling() StorageBilling {
	return make(map[int]*BillingData)
}
func (BD StorageBilling) Put(Country *BillingData) {
	BD[len(BD)] = Country
}

func (BD StorageBilling) ReadFileBilling() {
	r, err := ioutil.ReadFile(os.Getenv("BILLING_FILE"))
	if err != nil {
		panic(err)
	}
	var b []bool
	for i := 0; i < len(r); i++ {
		if r[i] == 49 {
			b = append(b, true)
		} else {
			b = append(b, false)
		}
	}
	NewCreateCustomer := b[5]
	NewPurchase := b[4]
	NewPayout := b[3]
	NewRecurring := b[2]
	NewFraudControl := b[1]
	NewCheckoutPage := b[0]
	NewBillingData := BillingData{
		CreateCustomer: NewCreateCustomer,
		Purchase:       NewPurchase,
		Payout:         NewPayout,
		Recurring:      NewRecurring,
		FraudControl:   NewFraudControl,
		CheckoutPage:   NewCheckoutPage,
	}
	BD.Put(&NewBillingData)

	//	fmt.Println(BD[0])
}
