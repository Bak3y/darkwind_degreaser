package puller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bak3y/darkwind_degreaser/internal/request"
	"github.com/Bak3y/darkwind_degreaser/internal/response"
	"github.com/pkg/errors"
)

//Gets Enjin News articles according to the defined incoming request and url
func GetEnjinNews(payload *request.NewsPayload, apiurl string) (*response.EnjinResults, error) {

	ereseults := response.EnjinResults{}

	b, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "Failed json.Marshal()")
	}

	req, err := http.NewRequest("POST", apiurl, bytes.NewBuffer(b))
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
	err = json.Unmarshal(body, &ereseults)

	if err != nil {
		return nil, errors.Wrap(err, "Failed json.Unmarshal(body, &ereseults)")
	}
	if ereseults.Result == nil {
		return nil, errors.Errorf("Unexpected nil results.")
	}
	return &ereseults, nil
}
