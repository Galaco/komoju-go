package komoju

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// ApiResponseError maps to a response error body
type ApiResponseError struct {
	// Data
	Data ApiResponseErrorMessage `json:"error"`
}

// ApiResponseErrorMessage
type ApiResponseErrorMessage struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Param   string `json:"param"`
}

// NewApiResponseError creates a new response error
func NewApiResponseError(code, param, message string) *ApiResponseError {
	return &ApiResponseError{
		Data: ApiResponseErrorMessage{
			Message: message,
			Code:    code,
			Param:   param,
		},
	}
}

// Error returns the error in a simplfied manner
func (err ApiResponseError) Error() string {
	return fmt.Sprintf("Code: %s, Param: %s, Message: %s", err.Data.Code, err.Data.Param, err.Data.Message)
}

// ApiResponseList provides properties across all
// api response body lists
type ApiResponseList struct {
	// Resource
	Resource string `json:"list"`
	// Total
	Total int64 `json:"total"`
	// Page
	Page int64 `json:"page"`
	// PerPage
	PerPage int64 `json:"per_page"`
	// lastPage
	LastPage int64 `json:"last_page"`
}

// Events
// ApiResponseEvents maps to a response to list events
type ApiResponseEvents struct {
	ApiResponseList
	Data []ApiResponseEvent `json:"data"`
}

// ApiResponseEvent maps to a response for a single event
type ApiResponseEvent struct {
	// Id
	Id string `json:"id"`
	// Type
	Type string `json:"type"`
	// Resource
	Resource string `json:"resource"`
	// Data
	Data ApiResponsePaymentDetail `json:"data"`
	// CreatedAt
	CreatedAt string `json:"created_at"`
}

// Payment
// ApiResponsePayments maps to a response for a list of payments
type ApiResponsePayments struct {
	ApiResponseList
	// Data
	Data []ApiResponsePayment `json:"data"`
}

// ApiResponsePayment maps to a response for a single payment
type ApiResponsePayment struct {
	// Id
	Id string `json:"id"`
	// Resource
	Resource string `json:"resource"`
	// Status
	Status string `json:"status"`
	// Amount
	Amount float64 `json:"amount"`
	// Tax
	Tax float64 `json:"tax"`
	// Customer
	Customer ApiResponseCustomer `json:"customer"`
	// PaymentDeadline
	PaymentDeadline string `json:"payment_deadline"`
	// PaymentDetails
	PaymentDetails ApiResponsePaymentDetail `json:"payment_details"`
	// PaymentMethodFee
	PaymentMethodFee float64 `json:"payment_method_fee"`
	// Total
	Total float64 `json:"total"`
	// Currency
	Currency string `json:"currency"`
	// Description
	Description string `json:"description"`
	// CapturedAt
	CapturedAt string `json:"captured_at"`
	// ExternalOrderNumber
	ExternalOrderNumber string `json:"external_order_num"`
	// Metadata
	MetaData map[string]string `json:"metadata"`
	// CreatedAt
	CreatedAt string `json:"created_at"`
	// AmountRefunded
	AmountRefunded float64 `json:"amount_refunded"`
	// Locale
	Locale string `json:"locale"`
	// Refunds
	Refunds []ApiResponseRefund `json:"refunds"`
	// RefundRequests
	RefundRequests []string `json:"refund_requests"`
}

// ApiResponsePaymentDetail maps to a response for a Payment
type ApiResponsePaymentDetail struct {
	// Type
	Type PaymentMethod `json:"type"`
	// Email
	Email string `json:"email"`
	// GivenName
	GivenName string `json:"given_name"`
	// FamilyName
	FamilyName string `json:"family_name"`
	// GivenNamek\Kana
	GivenNameKana string `json:"given_name_kana"`
	// FamilyNameKana
	FamilyNameKana string `json:"family_name_kana"`
	// Phone
	Phone string `json:"phone"`

	// Credit card
	// Brand
	Brand string `json:"brand"`
	// last4Digits
	Last4Digits string `json:"last_four_digits"`
	// Month
	Month int64 `json:"month"`
	// Year
	Year int64 `json:"year"`

	// Konbini
	// Store
	Store string `json:"store"`
	// ConfirmationCode
	ConfirmationCode string `json:"confirmation_code"`
	// Receipt
	Receipt string `json:"receipt"`
	// InstructionsUrl
	InstructionsUrl string `json:"instructions_url"`

	// Bank transfer
	// orderId
	OrderId string `json:"order_id"`
	// BankName
	BankName string `json:"bank_name"`
	// AccountBranchName
	AccountBranchName string `json:"account_branch_name"`
	// AccountNumber
	AccountNumber string `json:"account_number"`
	// AccountType
	AccountType string `json:"account_type"`
	// AccountName
	AccountName string `json:"account_name"`

	// PayEasy
	// BankId
	BankId string `json:"bank_id"`
	// CustomerId
	CustomerId string `json:"customer_id"`
	// ConfirmationId
	ConfirmationId string `json:"confirmation_id"`

	// WebMoney/Bitcash/Nanaco/Net Cash
	// ShortAmount
	ShortAmount float64 `json:"short_amount"`
	// PrepaidCards
	PrepaidCards []ApiResponsePaymentPrepaidCard `json:"prepaid_cards"`
}

// ApiResponsePaymentPrepaidCard maps to prepaid card
// properties on payment details
type ApiResponsePaymentPrepaidCard struct {
	// Last4Digits
	Last4Digits string `json:"last_four_digits"`
	// Points
	Points int64 `json:"points"`
}

// ApiResponseRefund maps to refund property on payment details
type ApiResponseRefund struct {
	// id
	Id string `json:"id"`
	// Resource
	Resource string `json:"resource"`
	// Amount
	Amount float64 `json:"amount"`
	// Currency
	Currency string `json:"currency"`
	// Payment
	Payment string `json:"payment"`
	// Description
	Description string `json:"description"`
	// CreatedAt
	CreatedAt string `json:"created_at"`
	// Chargeback
	Chargeback string `json:"chargeback"`
}

// Customers

// ApiResponseCustomers maps to response body of customer list
type ApiResponseCustomers struct {
	ApiResponseList
	// Data
	Data []ApiResponseCustomer `json:"data"`
}

// ApiResponseCustomer maps to response body of single customer
type ApiResponseCustomer struct {
	// id
	Id string `json:"id"`
	// Resource
	Resource string `json:"resource"`
	// Email
	Email string `json:"email"`
	// Source
	Source string `json:"source"`
	// Metadata
	Metadata map[string]string `json:"metadata"`
	// CreatedAt
	CreatedAt string `json:"created_at"`
}

// ParseApiResponse
func ParseApiResponse(payload io.ReadCloser, resp interface{}) (err error) {
	if err := unmarshallApiResponse(payload, resp); err != nil {
		return err
	}
	return nil
}

func unmarshallApiResponse(payload io.ReadCloser, target interface{}) error {
	body, err := ioutil.ReadAll(payload)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, target); err != nil {
		apiResponseErr := ApiResponseError{}
		if err = json.Unmarshal(body, &apiResponseErr); err != nil {
			return err
		}
		return apiResponseErr
	}

	return err
}
