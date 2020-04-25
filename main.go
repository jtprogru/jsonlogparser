package main

import (
	"fmt"
	"os"
	"net"
	"path"
)

const iso8601DateFormat = "2006-01-02T15:04:05-07:00"

type LogString struct {
	timestamp string `json:"@timestamp"`
	remote_addr string `json:"remote_addr"`
	cookie_bar string `json:"cookie_bar"`
	set_cookie string `json:"set_cookie"`
	body_bytes_sent string `json:"body_bytes_sent"`
	status string `json:"status"`
	request string `json:"request"`
	url string `json:"url"`
	request_method string `json:"request_method"`
	upstream string `json:"upstream"`
	response_time string `json:"response_time"`
	http_referrer string `json:"http_referrer"`
	http_user_agent string `json:"http_user_agent"`
}

func main() {
	f_name := os.Args[1]

	_, f := path.Split(f_name)
	fmt.Println(f)
	file, err := os.Open(f_name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
}
