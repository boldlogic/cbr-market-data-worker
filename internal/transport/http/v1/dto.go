package v1

type CbRequest struct {
	Type     string `json:"type"`
	CharCode string `json:"code,omitempty"`
	DateFrom string `json:"dateFrom,omitempty"`
	DateTo   string `json:"dateTo,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

type APIResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       interface{}
}

type Body struct {
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

type CurrencyDTO struct {
	ISOCode  int    `json:"isoCode"`
	CharCode string `json:"code"`
	NameRu   string `json:"nameRu,omitempty"`
	NameEn   string `json:"nameEn,omitempty"`
}
