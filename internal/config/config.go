package config

import (
	"fmt"
	"os"
)

var eapikey = os.Getenv("ENJIN_API_KEY")
var eapiurl = os.Getenv("ENJIN_API_URL")
var esiteid = os.Getenv("ENJIN_SITE_ID")
var euserid = os.Getenv("ENJIN_USER_ID")
var elimit = os.Getenv("ENJIN_POST_LIMIT")
var wpurl = os.Getenv("WP_URL")
var wpauth = os.Getenv("WP_AUTH")

func CheckEnvVars() (string, string, string, string, string, string, string, error) {
	if eapikey == "" {
		return "", "", "", "", "", "", "", fmt.Errorf("ENJIN_API_KEY not set")

	}
	if eapiurl == "" {
		return "", "", "", "", "", "", "", fmt.Errorf("ENJIN_API_URL not set")

	}
	if esiteid == "" {
		return "", "", "", "", "", "", "", fmt.Errorf("ENJIN_SITE_ID not set")

	}
	if euserid == "" {
		return "", "", "", "", "", "", "", fmt.Errorf("ENJIN_USER_ID not set")

	}
	if elimit == "" {
		return "", "", "", "", "", "", "", fmt.Errorf("ENJIN_POST_LIMIT not set")

	}
	if wpurl == "" {
		return "", "", "", "", "", "", "", fmt.Errorf("WP_URL not set")

	}
	if wpauth == "" {
		return "", "", "", "", "", "", "", fmt.Errorf("WP_AUTH not set")

	}
	return eapikey, eapiurl, esiteid, euserid, elimit, wpurl, wpauth, nil
}
