package redsys

// MerchantParametersResponse struct to read Redsys API responses
type MerchantParametersResponse struct {
	Date              string `json:"Ds_Date"`
	Hour              string `json:"Ds_Hour"`
	SecurePayment     string `json:"Ds_SecurePayment"`
	CardCountry       string `json:"Ds_Card_Country,omitempty"`
	Amount            string `json:"Ds_Amount"`
	Currency          string `json:"Ds_Currency"`
	Order             string `json:"Ds_Order"`
	MerchantCode      string `json:"Ds_MerchantCode"`
	Terminal          string `json:"Ds_Terminal"`
	Response          string `json:"Ds_Response"`
	MerchantData      string `json:"Ds_MerchantData"`
	TransactionType   string `json:"Ds_TransactionType"`
	ConsumerLanguage  string `json:"Ds_ConsumerLanguage,omitempty"`
	AuthorisationCode string `json:"Ds_AuthorisationCode,omitempty"`
	Identifier        string `json:"Ds_Merchant_Identifier,omitempty"`
	CofTxnid          string `json:"Ds_Merchant_Cof_Txnid,omitempty"`
	ExpiryDate        string `json:"Ds_ExpiryDate,omitempty"`
	Description       string `json:"Ds_Response_Description,omitempty"`
	CardNumber        string `json:"Ds_Card_Number,omitempty"`
	CardBrand         string `json:"Ds_Card_Brand,omitempty"`
}

// MerchantParametersRequest struct to construct Redsys API requests
type MerchantParametersRequest struct {
	// Optional fields are tagged with omitempty
	MerchantMerchantCode       string `json:"Ds_Merchant_MerchantCode"`
	MerchantTerminal           string `json:"Ds_Merchant_Terminal"`
	MerchantTransactionType    string `json:"Ds_Merchant_TransactionType"`
	MerchantAmount             string `json:"Ds_Merchant_Amount"`
	MerchantCurrency           string `json:"Ds_Merchant_Currency"`
	MerchantOrder              string `json:"Ds_Merchant_Order"`
	MerchantMerchantUrl        string `json:"Ds_Merchant_MerchantURL,omitempty"`
	MerchantProductDescription string `json:"Ds_Merchant_ProductDescription,omitempty"`
	MerchantTitular            string `json:"Ds_Merchant_Titular,omitempty"`
	MerchantUrlOK              string `json:"Ds_Merchant_UrlOK,omitempty"`
	MerchantUrlKO              string `json:"Ds_Merchant_UrlKO,omitempty"`
	MerchantMerchantName       string `json:"Ds_Merchant_MerchantName,omitempty"`
	MerchantConsumerLanguage   string `json:"Ds_Merchant_ConsumerLanguage,omitempty"`
	MerchantIdentifier         string `json:"Ds_Merchant_Identifier,omitempty"`
	MerchantCofIni             string `json:"Ds_Merchant_Cof_Ini,omitempty"`
	MerchantCofType            string `json:"Ds_Merchant_Cof_Type,omitempty"`
	MerchantData               string `json:"Ds_Merchant_MerchantData,omitempty"`
	MerchantCofTxnid           string `json:"Ds_Merchant_Cof_Txnid,omitempty"`
	MerchantExcepSca           string `json:"Ds_Merchant_Excep_SCA,omitempty"`
	MerchantDirectPayment      string `json:"Ds_Merchant_DirectPayment,omitempty"`
	MerchantGroup              string `json:"Ds_Merchant_MerchantGroup,omitempty"`
}

type RequestPreauthorization struct {
	Ds_MerchantParameters string `json:"Ds_MerchantParameters"`
	Ds_SignatureVersion   string `json:"Ds_SignatureVersion"`
	Ds_Signature          string `json:"Ds_Signature"`
}

type RequestErrorMessage struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorCodeDescription"`
}
