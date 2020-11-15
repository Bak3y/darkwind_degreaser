package main

import (
	"fmt"
	"os"

	"github.com/Bak3y/darkwind_degreaser/internal/config"
	"github.com/Bak3y/darkwind_degreaser/internal/puller"
	"github.com/Bak3y/darkwind_degreaser/internal/pusher"
	"github.com/Bak3y/darkwind_degreaser/internal/request"
)

var status string

func main() {
	// make sure all env vars are set
	eapikey, eapiurl, esiteid, euserid, elimit, wpurl, err := config.CheckEnvVars()
	if err != nil {
		fmt.Println("Error getting env vars:", err)
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
			Limit:   elimit,
		},
	}

	// get "Limit" amount of items from News.getLatest Enjin api
	enjinstuff, err := puller.GetEnjinNews(payloadchunk, eapiurl)
	if err != nil {
		fmt.Println("Error getting Enjin news:", err)
		os.Exit(1)
	}

	// create WordPress news posts from Enjin data
	for _, enjinpost := range enjinstuff.Result {

		status, err = pusher.CreateWPNews(enjinpost, wpurl)
		if err != nil {
			fmt.Println("Error creating WordPress news:", err)
			os.Exit(1)
		}
		fmt.Printf(status)

	}
}
