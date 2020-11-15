package response

//describes the structure of the data coming back from Enjin
type EnjinNews struct {
	Preset_Id   string `json:"preset_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Timestamp   string `json:"timestamp"`
	UserID      string `json:"user_id"`
	ArticleID   string `json:"article_id"`
	NumComments string `json:"num_comments"`
	Username    string `json:"username"`
	Avatar      string `json:"avatar"`
}

//shoutout to Keith again - go wizard
type EnjinResults struct {
	Result  []*EnjinNews `json:"result"`
	ID      string       `json:"id"`
	JSONRPC string       `json:"jsonrpc"`
}
