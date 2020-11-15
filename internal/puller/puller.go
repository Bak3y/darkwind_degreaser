package puller

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetEnjinNews() ([]byte, error) {

	url := "https://www.darkwindgaming.com/api/v1/api.php"

	payload := strings.NewReader("{\n\t\"jsonrpc\":\"2.0\",\n\t\"id\":\"262230\",\n\t\"params\":{\n\t\t\"api_key\": \"75375895eec487b0090d5c4060accea2a089764ff681e06a\",\n\t\t\"site_id\": \"108898\",\n\t\t\"limit\": \"999\"\n\t},\n\t\"method\":\"News.getLatest\"\n}\n")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cookie", "__cfduid=d0fe7cd1b4c6b29d79394286998e7d1261605408849; api_auth=hif7skej2s4aheh8g65gdas9k4")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
