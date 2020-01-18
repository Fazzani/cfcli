package domain

type RecordsResult struct {
	Result     []Record      `json:"result"`
	ResultInfo ResultInfo    `json:"result_info"`
	Success    bool          `json:"success"`
	Errors     []interface{} `json:"errors"`
	Messages   []interface{} `json:"messages"`
}

type Record struct {
	ID         string      `json:"id"`
	Type       Type        `json:"type"`
	Name       string      `json:"name"`
	Content    string      `json:"content"`
	Proxiable  bool        `json:"proxiable"`
	Proxied    bool        `json:"proxied"`
	TTL        int64       `json:"ttl"`
	Locked     bool        `json:"locked"`
	ZoneID     string      `json:"zone_id"`
	ZoneName   string      `json:"zone_name"`
	ModifiedOn string      `json:"modified_on"`
	CreatedOn  string      `json:"created_on"`
	Meta       MetaRecords `json:"meta"`
	Priority   *int64      `json:"priority,omitempty"`
}

type MetaRecords struct {
	AutoAdded           bool `json:"auto_added"`
	ManagedByApps       bool `json:"managed_by_apps"`
	ManagedByArgoTunnel bool `json:"managed_by_argo_tunnel"`
}

type Type string

const (
	A     Type = "A"
	MX    Type = "MX"
	Txt   Type = "TXT"
	CNAME Type = "CNAME"
	AAAA  Type = "AAAA"
)
