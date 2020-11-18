package pusher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Bak3y/darkwind_degreaser/internal/response"
	"github.com/Bak3y/darkwind_degreaser/internal/wpconverter"
	"github.com/pkg/errors"
)

func CreateWPNews(eresponse *response.EnjinNews, wpurl string, wpauth string) (string, error) {
	wpdata, err := wpconverter.Convert(eresponse)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(wpdata)
	if err != nil {
		return "", errors.Wrap(err, "Failed json.Marshal()")
	}

	req, err := http.NewRequest("POST", wpurl, bytes.NewBuffer(b))

	if err != nil {
		return "", errors.Wrap(err, "Failed creating request to send to wordpress.")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+wpauth)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "Failed https.DefaultClient.Do(req)")
	}

	defer res.Body.Close()

	if res.StatusCode != 201 {
		return "", errors.Errorf("Unexpected status code from wordpress: %v", res.StatusCode)
	}
	postresponse := fmt.Sprintf("Posted %s on %s\n", wpdata.Title, wpdata.Date)
	return postresponse, nil
}
