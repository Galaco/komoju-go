package komoju

const (
	komojuDomain            = "https://komoju.com/"
	hostedPageLocale        = "{locale}"
	hostedPageMerchant      = "{merchant_uuid}"
	hostedPagePaymentMethod = "{payment_method}"
	hostedPageURI           = hostedPageLocale + "/api/" + hostedPageMerchant + "/transactions/" + hostedPagePaymentMethod + "/new"
)
