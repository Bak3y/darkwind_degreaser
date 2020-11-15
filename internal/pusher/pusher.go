package pusher

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Bak3y/darkwind_degreaser/internal/response"
	"github.com/Bak3y/darkwind_degreaser/internal/wpconverter"
	"github.com/pkg/errors"
)

func CreateWPNews(eresponse *response.EnjinResponse, wpurl string) (string, error) {
	wpdata, err := wpconverter.Convert(eresponse)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(wpdata)
	if err != nil {
		return "", errors.Wrap(err, "Failed json.Marshal()")
	}

	buf := bytes.NewBuffer(b)

	http.NewRequest("post", wpurl, buf)
	return "posted", nil
}
