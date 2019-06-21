package komoju

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

// HostedPageConfig provides a struct that holds all parameters
// that can be passed to the HostedPage API
type HostedPageConfig struct {
	// Transaction contains information about the order
	Transaction HostedPageTransaction `url:"transaction"`
	// Timestamp should be the current time
	Timestamp int64 `url:"timestamp"`
}

// HostedPageTransaction contains transaction information for the hosted page
type HostedPageTransaction struct {
	// Amount is how much this transaction is for
	Amount float64 `url:"amount"`
	// Currency is the currency this transaction will use
	Currency string `url:"currency"`
	// Customer provides customer personal information
	Customer HostedPageCustomer `url:"customer"`
	// ExternalOrderNumber is a unique order id
	ExternalOrderNumber string `url:"external_order_num"`
	// ReturnUrl provides a URL to redirect after completion
	ReturnUrl string `url:"return_url"`
	// cancelUrl provides a return URL for transaction cancellation
	CancelUrl string `url:"cancel_url"`
	// Tax is the amount applied for tax
	Tax float32 `url:"tax"`
	// Metadata is optional.
	Metadata map[string]string `url:"metadata"`
}

// HostedPageCustimer contains customer information, used to pre-fill
// the Hosted Page form
type HostedPageCustomer struct {
	// GivenName is given/first name
	GivenName string `url:"given_name"`
	// FamilyName is family/last name
	FamilyName string `url:"family_name"`
	// GivenNameKana is given/first name represented as Kana
	GivenNameKana string `url:"given_name_kana"`
	// FamilyNameKana is family/last name represented as Kana
	FamilyNameKana string `url:"family_name_kana"`
	// Email is the cutomers email address
	Email string `url:"email"`
	// Phone is the customers phone number
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
func HostedPageURI(secretKey []byte, baseURI string, config *HostedPageConfig) (string, error) {
	endpoint := baseURI + "?"

	encoded, err := query.Values(config)
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
