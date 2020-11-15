package main

import (
	"fmt"
	"os"

	"github.com/Bak3y/darkwind_degreaser/internal/config"
	"github.com/Bak3y/darkwind_degreaser/internal/puller"
	"github.com/Bak3y/darkwind_degreaser/internal/pusher"
	"github.com/Bak3y/darkwind_degreaser/internal/request"
)

func main() {
	// make sure all env vars are set
	eapikey, eapiurl, esiteid, euserid, wpurl, err := config.CheckEnvVars()
	if err != nil {
		fmt.Println("Error getting env vars: %w", err)
		os.Exit(1)
	}

	// format request for enjin
	payloadchunk := &request.NewsPayload{
		JSONRPC: "2.0",
		ID:      euserid,
		Method:  "News.getLatest",
		Params: &request.NewsParams{
			Api_key: eapikey,
			Site_id: esiteid,
			Limit:   "999",
		},
	}

	// get "Limit" amount of items from News.getLatest Enjin api
	enjinstuff, err := puller.GetEnjinNews(payloadchunk, eapiurl)
	if err != nil {
		fmt.Println("Error getting Enjin News: %w", err)
		os.Exit(1)
	}

	// create WordPress news posts from Enjin data
	status, err := pusher.CreateWPNews(enjinstuff)
	if err != nil {
		fmt.Println("Error creating wordpress news: %w", err)
		os.Exit(1)
	}

	//show something meaniningful to the user
	fmt.Printf("Created WP articles: %s", status)
}
