package komoju_go

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/google/go-querystring/query"
	"sort"
	"strings"
)

const (
	paramNameHmac = "hmac"
)

type HostedPageConfig struct {
	Transaction HostedPageTransaction `url:"transaction"`
	Timestamp int64 `url:"timestamp"`
}

type HostedPageTransaction struct {
	Amount float64 `url:"amount"`
	Currency string `url:"currency"`
	Customer HostedPageCustomer `url:"customer"`
	ExternalOrderNumber string `url:"external_order_num"`
	ReturnUrl string `url:"return_url"`
	CancelUrl string `url:"cancel_url"`
	Tax float32 `url:"tax"`
	Metadata map[string]string `url:"metadata"`
}

type HostedPageCustomer struct {
	GivenName string `url:"given_name"`
	FamilyName string `url:"family_name"`
	GivenNameKana string `url:"given_name_kana"`
	FamilyNameKana string `url:"family_name_kana"`
	Email string `url:"email"`
	Phone string `url:"phone"`
}

// HostPageBaseURI create the base URI for Komoju Hosted page api request
func HostedPageBaseURI(locale, merchantUuid string, paymentMethod PaymentMethod) string {
	return komojuDomain + strings.Replace(
		strings.Replace(
			strings.Replace(hostedPageURI, hostedPagePaymentMethod, string(paymentMethod), -1),
			hostedPageMerchant, merchantUuid, -1),
		hostedPageLocale, locale, -1)
}

// HostedPageURI returns the URI for a hosted page
func HostedPageURI(secretKey []byte, baseURI string, config *HostedPageConfig) (string,error) {
	endpoint := baseURI + "?"

	encoded,err := query.Values(config)
	if err != nil {
		return "", err
	}

	//sort alphabetically
	unmangled := strings.Split(encoded.Encode(), "&")
	sort.Strings(unmangled)
	endpoint += strings.Join(unmangled, "&")

	mac := hmac.New(sha256.New, secretKey)
	mac.Write([]byte(endpoint))
	checkSum := hex.EncodeToString(mac.Sum(nil))

	endpoint += "&" + paramNameHmac + "=" + string(checkSum)

	return endpoint, nil
}