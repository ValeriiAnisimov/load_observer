package observer

import (
	"errors"
	"io/ioutil"
	"math"
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

func getResultBilling() (BillingData, error) {
	result := BillingData{}

	f, err := os.Open("../simulator/billing.data")
	if err != nil {
		return result, err
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return result, err
	}

	var bitMask int8 = 0

	if len(data) != 6 {
		return result, errors.New("billing mask string not valid")
	}

	for i, v := range data {
		if v == '1' {
			position := len(data) - 1 - i
			bitMask += int8(math.Pow(2, float64(position)))
		}
	}

	result.CreateCustomer = bitMask>>0&1 == 1
	result.Purchase = bitMask>>1&1 == 1
	result.Payout = bitMask>>2&1 == 1
	result.Recurring = bitMask>>3&1 == 1
	result.FraudControl = bitMask>>4&1 == 1
	result.CheckoutPage = bitMask>>5&1 == 1

	return result, nil
}
