package komoju_go

// PaymentMethod
type PaymentMethod string

const (
	// PaymentMethodBankTransfer
	PaymentMethodBankTransfer = PaymentMethod("bank_transfer")
	// PaymentMethodBitCash
	PaymentMethodBitCash = PaymentMethod("bit_cash")
	// PaymentMethodCreditCard
	PaymentMethodCreditCard = PaymentMethod("credit_card")
	// PaymentMethodKonbini
	PaymentMethodKonbini = PaymentMethod("konbini")
	// PaymentMethodNanaco
	PaymentMethodNanaco = PaymentMethod("nanaco")
	// PaymentMethodNetCash
	PaymentMethodNetCash = PaymentMethod("net_cash")
	// PaymentMethodPayEasy
	PaymentMethodPayEasy = PaymentMethod("pay_easy")
	// PaymentMethodWebMoney
	PaymentMethodWebMoney = PaymentMethod("web_money")
)