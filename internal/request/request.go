package request

type NewsPayload struct {
	JSONRPC string
	ID      string
	Method  string
	Params  *NewsParams
}

type NewsParams struct {
	Api_key string
	Site_id string
	Limit   string
}

type WPPost struct {
	Date    string
	Status  string
	Title   string
	Content string
	Author  string
}
