package wpconverter

import (
	"strconv"
	"time"

	"github.com/Bak3y/darkwind_degreaser/internal/request"
	"github.com/Bak3y/darkwind_degreaser/internal/response"
)

func Convert(enjinstuff *response.EnjinNews) (*request.WPPost, error) {
	var wpstuff = request.WPPost{}

	wpstuff.Title = enjinstuff.Title + " by " + enjinstuff.Username
	wpstuff.Content = enjinstuff.Content
	wpstuff.Date = GetTimestamp(enjinstuff.Timestamp)
	wpstuff.Status = "draft"
	return &wpstuff, nil
}

//shoutout to Keith for the help
func GetTimestamp(stimestamp string) string {
	i, err := strconv.ParseInt(stimestamp, 10, 64)
	if err != nil {
		return ""
	}
	unixtime := time.Unix(i, 0)
	return unixtime.Format(time.RFC3339)
}
