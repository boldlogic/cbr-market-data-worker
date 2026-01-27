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

type TaskDTO struct {
	Action string `json:"action"`
	Uuid   string `json:"uuid,omitempty"`

	Params Params `json:"params,omitempty"`
	//ScheduledAt string `json:"scheduledAt,omitempty"`
}

type Params struct {
	CcyCode  string `json:"ccyCode,omitempty"`
	DateFrom string `json:"dateFrom,omitempty"`
	DateTo   string `json:"dateTo,omitempty"`
}
