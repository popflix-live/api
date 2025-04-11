package instance

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var Client *resty.Client

func init() {
	Client = resty.New().
		SetBaseURL("https://anitaku.io").
		SetTimeout(5 * time.Second)
}
