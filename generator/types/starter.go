package types

type StaterResponse struct {
	Token               string              `json:"token"`
	Checksum            string              `json:"checksum"`
	DeviceID            string              `json:"deviceId"`
	ClientConfiguration ClientConfiguration `json:"clientConfiguration"`
}
type PropertyDescriptors struct {
	Chrome     []string `json:"chrome"`
	Navigator  []string `json:"Navigator"`
	Screen     []string `json:"Screen"`
	NavigatorL []string `json:"navigator"`
	Window     []string `json:"window"`
}
type CheckoutDeliveryOption0 struct {
	Selector string `json:"selector"`
}
type PostcodeObfuscated struct {
	ValueSelector string `json:"valueSelector"`
	Context       string `json:"context"`
	Selector      string `json:"selector"`
	Operation     string `json:"operation"`
}
type CheckoutDeliveryOption1 struct {
	Selector string `json:"selector"`
}
type AddedGiftCard struct {
	Selector string `json:"selector"`
}
type CheckoutDeliveryOption2 struct {
	Selector string `json:"selector"`
}
type MainContactPreferencesClick struct {
	Selector string `json:"selector"`
}
type CheckoutDeliveryOption3 struct {
	Selector string `json:"selector"`
}
type CheckoutDeliveryOption4 struct {
	Selector string `json:"selector"`
}
type CheckoutDeliveryOption5 struct {
	Selector string `json:"selector"`
}
type ContactPreferencesLifestyleSms struct {
	Selector string `json:"selector"`
}
type ContactPreferencesSelectionAll struct {
	Selector string `json:"selector"`
}
type ContactPreferencesNewnessSms struct {
	Selector string `json:"selector"`
}
type MainMyBagCheckoutClick struct {
	Selector string `json:"selector"`
}
type MainMyOrdersClick struct {
	Selector string `json:"selector"`
}
type ContactPreferencesNewnessEmail struct {
	Selector string `json:"selector"`
}
type CheckoutDeliveryOption6 struct {
	Selector string `json:"selector"`
}
type CheckoutDeliveryOption7 struct {
	Selector string `json:"selector"`
}
type CheckoutDeliveryOption8 struct {
	Selector string `json:"selector"`
}
type ContactPreferencesPartnerSms struct {
	Selector string `json:"selector"`
}
type MainCheckoutClick struct {
	Selector string `json:"selector"`
}
type EmailAddressObfuscated struct {
	ValueSelector string `json:"valueSelector"`
	Context       string `json:"context"`
	Selector      string `json:"selector"`
	Operation     string `json:"operation"`
}
type ContactPreferencesPromosEmail struct {
	Selector string `json:"selector"`
}
type MainMyBagSaveForLaterClick struct {
	Selector string `json:"selector"`
}
type RemovePaymentMethodConfirm struct {
	Selector string `json:"selector"`
}
type AddedGiftVoucher struct {
	Selector string `json:"selector"`
}
type MainAddToSavedListsClick struct {
	Selector []string `json:"selector"`
}
type MyAccountDeleteAddress struct {
	Selector string `json:"selector"`
}
type ContactTweeter struct {
	Selector string `json:"selector"`
}
type MainMyAccountClick struct {
	Selector string `json:"selector"`
}
type ContactChat struct {
	Selector string `json:"selector"`
}
type MainMyBagDeleteItemClick struct {
	Selector string `json:"selector"`
}
type MyAccountSaveAddress struct {
	Selector string `json:"selector"`
}
type BuyGiftVoucher struct {
	Selector string `json:"selector"`
}
type RemovePaymentMethodPaypal struct {
	Selector string `json:"selector"`
}
type EmailAddressDomain struct {
	ValueSelector string `json:"valueSelector"`
	Context       string `json:"context"`
	Selector      string `json:"selector"`
	Operation     string `json:"operation"`
}
type MainAddToBagClick struct {
	Selector string `json:"selector"`
}
type MainSavedListsDeleteItemClick struct {
	Selector string `json:"selector"`
}
type AddPaymentMethodPaypal struct {
	Selector string `json:"selector"`
}
type ContactPreferencesPromosSms struct {
	Selector string `json:"selector"`
}
type ContactEmail struct {
	Selector string `json:"selector"`
}
type ContactPreferencesSave struct {
	Selector string `json:"selector"`
}
type ContactPreferencesPartnerEmail struct {
	Selector string `json:"selector"`
}
type AddPaymentMethodCard struct {
	Selector string `json:"selector"`
}
type MainSavedListsClick struct {
	Selector string `json:"selector"`
}
type MainGiftVouturePaySecurelyNowClick struct {
	Selector string `json:"selector"`
}
type MainReturnsInfoClick struct {
	Selector string `json:"selector"`
}
type ContactMessenger struct {
	Selector string `json:"selector"`
}
type MainDeliverDeliverToChangeClick struct {
	Selector string `json:"selector"`
}
type MainViewBagClick struct {
	Selector string `json:"selector"`
}
type MainShoppingFromChangedClick struct {
	Selector string `json:"selector"`
}
type ContactPreferencesLifestyleEmail struct {
	Selector string `json:"selector"`
}
type BuyGiftVoucherPay struct {
	Selector string `json:"selector"`
}
type MainBagItemDeleteClick struct {
	Selector string `json:"selector"`
}
type RemoteTags struct {
	CheckoutDeliveryOption0            CheckoutDeliveryOption0            `json:"checkout-delivery-option-0"`
	PostcodeObfuscated                 PostcodeObfuscated                 `json:"postcode_obfuscated"`
	CheckoutDeliveryOption1            CheckoutDeliveryOption1            `json:"checkout-delivery-option-1"`
	AddedGiftCard                      AddedGiftCard                      `json:"added-gift-card"`
	CheckoutDeliveryOption2            CheckoutDeliveryOption2            `json:"checkout-delivery-option-2"`
	MainContactPreferencesClick        MainContactPreferencesClick        `json:"main-contactPreferences-click"`
	CheckoutDeliveryOption3            CheckoutDeliveryOption3            `json:"checkout-delivery-option-3"`
	CheckoutDeliveryOption4            CheckoutDeliveryOption4            `json:"checkout-delivery-option-4"`
	CheckoutDeliveryOption5            CheckoutDeliveryOption5            `json:"checkout-delivery-option-5"`
	ContactPreferencesLifestyleSms     ContactPreferencesLifestyleSms     `json:"contact-preferences-lifestyle-sms"`
	ContactPreferencesSelectionAll     ContactPreferencesSelectionAll     `json:"contact-preferences-selection-all"`
	ContactPreferencesNewnessSms       ContactPreferencesNewnessSms       `json:"contact-preferences-newness-sms"`
	MainMyBagCheckoutClick             MainMyBagCheckoutClick             `json:"main-myBag-checkout-click"`
	MainMyOrdersClick                  MainMyOrdersClick                  `json:"main-myOrders-click"`
	ContactPreferencesNewnessEmail     ContactPreferencesNewnessEmail     `json:"contact-preferences-newness-email"`
	CheckoutDeliveryOption6            CheckoutDeliveryOption6            `json:"checkout-delivery-option-6"`
	CheckoutDeliveryOption7            CheckoutDeliveryOption7            `json:"checkout-delivery-option-7"`
	CheckoutDeliveryOption8            CheckoutDeliveryOption8            `json:"checkout-delivery-option-8"`
	ContactPreferencesPartnerSms       ContactPreferencesPartnerSms       `json:"contact-preferences-partner-sms"`
	MainCheckoutClick                  MainCheckoutClick                  `json:"main-checkout-click"`
	EmailAddressObfuscated             EmailAddressObfuscated             `json:"emailAddress_obfuscated"`
	ContactPreferencesPromosEmail      ContactPreferencesPromosEmail      `json:"contact-preferences-promos-email"`
	MainMyBagSaveForLaterClick         MainMyBagSaveForLaterClick         `json:"main-myBag-saveForLater-click"`
	RemovePaymentMethodConfirm         RemovePaymentMethodConfirm         `json:"remove-payment-method-confirm"`
	AddedGiftVoucher                   AddedGiftVoucher                   `json:"added-gift-voucher"`
	MainAddToSavedListsClick           MainAddToSavedListsClick           `json:"main-addToSavedLists-click"`
	MyAccountDeleteAddress             MyAccountDeleteAddress             `json:"my-account-delete-address"`
	ContactTweeter                     ContactTweeter                     `json:"contact_tweeter"`
	MainMyAccountClick                 MainMyAccountClick                 `json:"main-myAccount-click"`
	ContactChat                        ContactChat                        `json:"contact_chat"`
	MainMyBagDeleteItemClick           MainMyBagDeleteItemClick           `json:"main-myBag-deleteItem-click"`
	MyAccountSaveAddress               MyAccountSaveAddress               `json:"my-account-save-address"`
	BuyGiftVoucher                     BuyGiftVoucher                     `json:"buy-gift-voucher"`
	RemovePaymentMethodPaypal          RemovePaymentMethodPaypal          `json:"remove-payment-method-paypal"`
	EmailAddressDomain                 EmailAddressDomain                 `json:"emailAddress_domain"`
	MainAddToBagClick                  MainAddToBagClick                  `json:"main-addToBag-click"`
	MainSavedListsDeleteItemClick      MainSavedListsDeleteItemClick      `json:"main-savedLists-deleteItem-click"`
	AddPaymentMethodPaypal             AddPaymentMethodPaypal             `json:"add-payment-method-paypal"`
	ContactPreferencesPromosSms        ContactPreferencesPromosSms        `json:"contact-preferences-promos-sms"`
	ContactEmail                       ContactEmail                       `json:"contact_email"`
	ContactPreferencesSave             ContactPreferencesSave             `json:"contact-preferences-save"`
	ContactPreferencesPartnerEmail     ContactPreferencesPartnerEmail     `json:"contact-preferences-partner-email"`
	AddPaymentMethodCard               AddPaymentMethodCard               `json:"add-payment-method-card"`
	MainSavedListsClick                MainSavedListsClick                `json:"main-savedLists-click"`
	MainGiftVouturePaySecurelyNowClick MainGiftVouturePaySecurelyNowClick `json:"main-giftVouture-paySecurelyNow-click"`
	MainReturnsInfoClick               MainReturnsInfoClick               `json:"main-returnsInfo-click"`
	ContactMessenger                   ContactMessenger                   `json:"contact_messenger"`
	MainDeliverDeliverToChangeClick    MainDeliverDeliverToChangeClick    `json:"main-deliver-deliverToChange-click"`
	MainViewBagClick                   MainViewBagClick                   `json:"main-viewBag-click"`
	MainShoppingFromChangedClick       MainShoppingFromChangedClick       `json:"main-shoppingFromChanged-click"`
	ContactPreferencesLifestyleEmail   ContactPreferencesLifestyleEmail   `json:"contact-preferences-lifestyle-email"`
	BuyGiftVoucherPay                  BuyGiftVoucherPay                  `json:"buy-gift-voucher-pay"`
	MainBagItemDeleteClick             MainBagItemDeleteClick             `json:"main-bagItemDelete-click"`
}
type EventsReduceFactorMap struct {
	Pointermove float64 `json:"pointermove"`
	Mousemove   float64 `json:"mousemove"`
}
type Sensors struct {
	MaxSensorSamples     int `json:"maxSensorSamples"`
	SensorsDeltaInMillis int `json:"sensorsDeltaInMillis"`
}
type KeyboardCSSSelectors struct {
	LastName                  string `json:"lastName"`
	CountyStateProvinceOrArea string `json:"countyStateProvinceOrArea"`
	Address2                  string `json:"address2"`
	TelephoneMobile           string `json:"telephoneMobile"`
	Address1                  string `json:"address1"`
	Postcode                  string `json:"postcode"`
	Locality                  string `json:"locality"`
	GiftVoucherToEmail        string `json:"giftVoucherToEmail"`
	GiftVoucherFromName       string `json:"giftVoucherFromName"`
	FirstName                 string `json:"firstName"`
	EmailAddress              string `json:"emailAddress"`
	GiftVoucherToName         string `json:"giftVoucherToName"`
	MainSearch                string `json:"mainSearch"`
	GiftVoucherValue          string `json:"giftVoucherValue"`
}
type Num0 struct {
	TagsToFlushRegex             string                `json:"tagsToFlushRegex"`
	PropertyDescriptors          PropertyDescriptors   `json:"propertyDescriptors"`
	EventsBlackList              []string              `json:"eventsBlackList"`
	MetadataCacheTTLMinutes      int                   `json:"metadataCacheTtlMinutes"`
	Enabled                      bool                  `json:"enabled"`
	URL                          string                `json:"url"`
	RemoteTags                   RemoteTags            `json:"remoteTags"`
	KeyboardFieldBlackList       []string              `json:"keyboardFieldBlackList"`
	EventsReduceFactorMap        EventsReduceFactorMap `json:"eventsReduceFactorMap"`
	Sensors                      Sensors               `json:"sensors"`
	KeyboardIdentifierAttributes interface{}           `json:"keyboardIdentifierAttributes"`
	EventsToCountList            []string              `json:"eventsToCountList"`
	HeartBeatFrequencySeconds    int                   `json:"heartBeatFrequencySeconds"`
	KeyboardCSSSelectors         KeyboardCSSSelectors  `json:"keyboardCssSelectors"`
	BufferSize                   int                   `json:"bufferSize"`
}
type BehavioralBlacklist struct {
	Mouse    string `json:"mouse"`
	Indirect string `json:"indirect"`
	Gestures string `json:"gestures"`
}
type Num1 struct {
	BehavioralBlacklist BehavioralBlacklist `json:"behavioralBlacklist"`
}
type BehavioralBlacklist3 struct {
	Mouse         string `json:"mouse"`
	EventKeyboard string `json:"eventKeyboard"`
	Indirect      string `json:"indirect"`
	Gestures      string `json:"gestures"`
}
type Num2 struct {
	BehavioralBlacklist BehavioralBlacklist3 `json:"behavioralBlacklist"`
}
type Num3 struct {
	SendMetadata        bool                 `json:"sendMetadata"`
	BehavioralBlacklist BehavioralBlacklist3 `json:"behavioralBlacklist"`
}
type ClientConfiguration struct {
	Num0 Num0 `json:"0"`
	Num1 Num1 `json:"1"`
	Num2 Num2 `json:"2"`
	Num3 Num3 `json:"3"`
}
