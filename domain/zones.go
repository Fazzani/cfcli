package domain

type ZonesResult struct {
	Result     []Zone        `json:"result"`
	ResultInfo ResultInfo    `json:"result_info"`
	Success    bool          `json:"success"`
	Errors     []interface{} `json:"errors"`
	Messages   []interface{} `json:"messages"`
}

type ZoneResult struct {
	Result   Zone          `json:"result"`
	Success  bool          `json:"success"`
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
}

type Zone struct {
	ID                  string      `json:"id"`
	Name                string      `json:"name"`
	Status              string      `json:"status"`
	Paused              bool        `json:"paused"`
	Type                string      `json:"type"`
	DevelopmentMode     int64       `json:"development_mode"`
	NameServers         []string    `json:"name_servers"`
	OriginalNameServers []string    `json:"original_name_servers"`
	OriginalRegistrar   string      `json:"original_registrar"`
	OriginalDnshost     interface{} `json:"original_dnshost"`
	ModifiedOn          string      `json:"modified_on"`
	CreatedOn           string      `json:"created_on"`
	ActivatedOn         string      `json:"activated_on"`
	Meta                Meta        `json:"meta"`
	Owner               Owner       `json:"owner"`
	Account             Account     `json:"account"`
	Permissions         []string    `json:"permissions"`
	Plan                Plan        `json:"plan"`
}

type Account struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Meta struct {
	Step                    int64 `json:"step"`
	WildcardProxiable       bool  `json:"wildcard_proxiable"`
	CustomCertificateQuota  int64 `json:"custom_certificate_quota"`
	PageRuleQuota           int64 `json:"page_rule_quota"`
	PhishingDetected        bool  `json:"phishing_detected"`
	MultipleRailgunsAllowed bool  `json:"multiple_railguns_allowed"`
}

type Owner struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Email string `json:"email"`
}

type Plan struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	Price             int64  `json:"price"`
	Currency          string `json:"currency"`
	Frequency         string `json:"frequency"`
	IsSubscribed      bool   `json:"is_subscribed"`
	CanSubscribe      bool   `json:"can_subscribe"`
	LegacyID          string `json:"legacy_id"`
	LegacyDiscount    bool   `json:"legacy_discount"`
	ExternallyManaged bool   `json:"externally_managed"`
}

type ResultInfo struct {
	Page       int64 `json:"page"`
	PerPage    int64 `json:"per_page"`
	TotalPages int64 `json:"total_pages"`
	Count      int64 `json:"count"`
	TotalCount int64 `json:"total_count"`
}
