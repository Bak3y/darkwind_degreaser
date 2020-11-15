package main

import (
	"fmt"
	"os"

	"github.com/Bak3y/darkwind_degreaser/internal/puller"
	"github.com/Bak3y/darkwind_degreaser/internal/pusher"
	"github.com/Bak3y/darkwind_degreaser/internal/request"
)

const apiurl = "https://www.darkwindgaming.com/api/v1/api.php"
const siteid = "108898"

var apikey = os.Getenv("ENJIN_API_KEY")

func main() {
	payloadchunk := &request.NewsPayload{
		JSONRPC: "2.0",
		ID:      siteid,
		Method:  "News.getLatest",
		Params: &request.NewsParams{
			Api_key: apikey,
			Site_id: siteid,
			Limit:   "999",
		},
	}

	enjinstuff, err := puller.GetEnjinNews(payloadchunk, apiurl)
	if err != nil {
		fmt.Println("Error getting Enjin News: %w", err)
	}
	status, err := pusher.CreateWPNews(enjinstuff)
	if err != nil {
		fmt.Println("Error creating wordpress news: %w", err)
	}
	fmt.Printf("Created WP articles: %s", status)

}
