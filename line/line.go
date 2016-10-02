package line

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const NOTIFY_API = "https://notify-api.line.me/api/notify"
const HEADER_AUTHORIZATION = "Authorization"

type line struct {
	token string
}

func New(token string) *line {
	return &line{token}
}

func (l *line) Post(str string) error {
	const MESSAGE = "message"
	const BEARER = "Bearer "

	values := url.Values{}
	values.Set(MESSAGE, str)

	req, _ := http.NewRequest(
		"POST",
		NOTIFY_API,
		strings.NewReader(values.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add(HEADER_AUTHORIZATION, BEARER + l.token)


	client := &http.Client{Timeout: time.Duration(5 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//fmt.Println(resp.StatusCode)
	if(resp.StatusCode != http.StatusOK){
		fmt.Println("Error Status Code")
		fmt.Println(resp.StatusCode)
	}

	return nil
}
