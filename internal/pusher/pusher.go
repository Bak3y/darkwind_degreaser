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

func CreateWPNews(eresponse *response.EnjinNews, wpurl string) (string, error) {
	wpdata, err := wpconverter.Convert(eresponse)
	if err != nil {
		return "", err
	}

	b, err := json.Marshal(wpdata)
	if err != nil {
		return "", errors.Wrap(err, "Failed json.Marshal()")
	}

	http.NewRequest("post", wpurl, bytes.NewBuffer(b))
	postresponse := fmt.Sprintf("Posted %s by %s on %s\n", wpdata.Title, wpdata.Author, wpdata.Date)
	return postresponse, nil
}
