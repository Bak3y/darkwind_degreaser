package puller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

type Enjin struct{}

const apiurl = "https://www.darkwindgaming.com/api/v1/api.php"
const siteid = "108898"

var apikey = os.Getenv("ENJIN_API_KEY")

type NewsPayload struct {
	jsonrpc string
	id      string
	method  string
	params  *NewsParams
}

type NewsParams struct {
	api_key string
	site_id string
	limit   string
}

func GetEnjinNews() ([]byte, error) {

	payloadchunk := &NewsPayload{
		jsonrpc: "2.0",
		id:      siteid,
		method:  "News.getLatest",
		params: &NewsParams{
			api_key: apikey,
			site_id: siteid,
			limit:   "999",
		},
	}

	b, err := json.Marshal(payloadchunk)
	if err != nil {
		return nil, errors.Wrap(err, "Failed json.Marshal()")
	}

	buf := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", apiurl, buf)
	if err != nil {
		return nil, errors.Wrap(err, "Failed http.NewRequest()")
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "Failed https.DefaultClient.Do(req)")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
