package wpconverter

import (
	"time"

	"github.com/Bak3y/darkwind_degreaser/internal/request"
	"github.com/Bak3y/darkwind_degreaser/internal/response"
)

func Convert(enjinstuff *response.EnjinResponse) (*request.WPPost, error) {
	var wpstuff = request.WPPost{}

	wpstuff.Title = enjinstuff.Title
	wpstuff.Author = enjinstuff.Username
	wpstuff.Content = enjinstuff.Content

	wptime, err := GetTimestamp(enjinstuff.Timestamp)
	if err != nil {
		return nil, err
	}

	wpstuff.Date = wptime.String()
	return &wpstuff, nil
}

func GetTimestamp(stimestamp string) (time.Time, error) {

	timestamp, err := time.Parse(time.RFC3339, stimestamp)
	if err != nil {
		return time.Now(), err
	}
	return timestamp, nil
}
