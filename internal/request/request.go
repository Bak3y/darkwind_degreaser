package request

type NewsPayload struct {
	JSONRPC string `json:"jsonrpc"`
	ID      string `json:"id"`
	Method  string `json:"method"`
	Params  *NewsParams
}

type NewsParams struct {
	Api_key string `json:"api_key"`
	Site_id string `json:"site_id"`
	Limit   string `json:"limit"`
}

type WPPost struct {
	Date    string `json:"date"`
	Status  string `json:"status"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
