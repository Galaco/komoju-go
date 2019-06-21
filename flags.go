package komoju

// PaymentMethod is an alias to String type to
// restrict to valid payment options
type PaymentMethod string

const (
	// PaymentMethodBankTransfer is bank transfer payment method
	PaymentMethodBankTransfer = PaymentMethod("bank_transfer")
	// PaymentMethodBitCash is BitCash payment method
	PaymentMethodBitCash = PaymentMethod("bit_cash")
	// PaymentMethodCreditCard is Credit Card payment method
	PaymentMethodCreditCard = PaymentMethod("credit_card")
	// PaymentMethodKonbini is Konbini payment method
	PaymentMethodKonbini = PaymentMethod("konbini")
	// PaymentMethodNanaco is Nanaco payment method
	PaymentMethodNanaco = PaymentMethod("nanaco")
	// PaymentMethodNetCash is Net Cash payment method
	PaymentMethodNetCash = PaymentMethod("net_cash")
	// PaymentMethodPayEasy is Pay Easy payment method
	PaymentMethodPayEasy = PaymentMethod("pay_easy")
	// PaymentMethodWebMoney is Web Money payment method
	PaymentMethodWebMoney = PaymentMethod("web_money")
)
