package puller

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bak3y/darkwind_degreaser/internal/request"
	"github.com/pkg/errors"
)

func GetEnjinNews(payload *request.NewsPayload, apiurl string) ([]byte, error) {

	b, err := json.Marshal(payload)
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
