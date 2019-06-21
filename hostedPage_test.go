package komoju

import (
	"testing"
)

func TestHostedPageBaseURI(t *testing.T) {
	locale := "en"
	merchantUuid := "degica-mart"
	paymentMethod := PaymentMethodCreditCard

	expected := "https://komoju.com/en/api/degica-mart/transactions/credit_card/new"

	actual := HostedPageBaseURI(locale, merchantUuid, paymentMethod)

	if expected != actual {
		t.Error("incorrect hosted page uri generated")
	}
}

func TestHostedPageURI(t *testing.T) {
	locale := "en"
	merchantUuid := "degica-mart"
	paymentMethod := PaymentMethodCreditCard
	baseUri := HostedPageBaseURI(locale, merchantUuid, paymentMethod)

	secretKey := []byte("ABCD1234567890")
	config := HostedPageConfig{
		Transaction: HostedPageTransaction{
			Amount:   130,
			Currency: "JPY",
			Customer: HostedPageCustomer{
				GivenName:      "John",
				FamilyName:     "Smith",
				GivenNameKana:  "John",
				FamilyNameKana: "Smith",
				Email:          "",
				Phone:          "",
			},
			ExternalOrderNumber: "M8x6U6Z5HEeXv3",
			ReturnUrl:           "http://example.com/?sucess=true",
			CancelUrl:           "http://example.com/?cancel=true",
			Tax:                 0,
		},
		Timestamp: 1561022519,
	}

	expected := "https://komoju.com/en/api/degica-mart/transactions/credit_card/new?timestamp=1561022519&transaction%5Bamount%5D=130&transaction%5Bcancel_url%5D=http%3A%2F%2Fexample.com%2F%3Fcancel%3Dtrue&transaction%5Bcurrency%5D=JPY&transaction%5Bcustomer%5D%5Bemail%5D=&transaction%5Bcustomer%5D%5Bfamily_name%5D=Smith&transaction%5Bcustomer%5D%5Bfamily_name_kana%5D=Smith&transaction%5Bcustomer%5D%5Bgiven_name%5D=John&transaction%5Bcustomer%5D%5Bgiven_name_kana%5D=John&transaction%5Bcustomer%5D%5Bphone%5D=&transaction%5Bexternal_order_num%5D=M8x6U6Z5HEeXv3&transaction%5Bmetadata%5D=map%5B%5D&transaction%5Breturn_url%5D=http%3A%2F%2Fexample.com%2F%3Fsucess%3Dtrue&transaction%5Btax%5D=0&hmac=d4c67129924bbba0cd6e61c97140eb1c10cfa2781617a1e4d8d255b9255ac0f3"
	actual, err := HostedPageURI(secretKey, baseUri, &config)
	if err != nil {
		t.Error(err)
	}

	if expected != actual {
		t.Error("generated URI does not equal expected")
	}
}
