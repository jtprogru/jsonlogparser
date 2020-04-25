package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

const iso8601DateFormat = "2006-01-02T15:04:05-07:00"

type LogString struct {
	Timestamp     string `json:"@timestamp"`
	RemoteAddr    string `json:"remote_addr"`
	CookieBar     string `json:"cookie_bar,omitempty"`
	SetCookie     string `json:"set_cookie,omitempty"`
	BodyByteSent  string `json:"body_bytes_sent"`
	Status        string `json:"status"`
	Request       string `json:"request"`
	URL           string `json:"url"`
	RequestMethod string `json:"request_method"`
	Upstream      string `json:"upstream"`
	ResponseTime  string `json:"response_time"`
	HTTPReferer   string `json:"http_referrer,omitempty"`
	HTTPUserAgent string `json:"http_user_agent"`
}

func main() {
	if len(os.Args) < 1 {
		fmt.Println("ERROR")
	}
	fName := os.Args[1]
	file, err := os.Open(fName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	prep := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prep = append(prep, scanner.Text())
		//fmt.Println(scanner.Text())
		//fmt.Println(len(prep))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(prep))
	for k,_ := range prep {
		str := []byte(string(prep[k]))
		res := LogString{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println(res)
	}
}
