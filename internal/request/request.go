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
