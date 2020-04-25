package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"flag"
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

//type Report struct {
//	IP string `json:"remote_addr"`
//}

func main() {

	fname := flag.String("fname", "/var/log/nginx/access.log", "A path to access.log Nginx")
	howmuch := flag.Int("howmuch", 15, "Show a TOP of remote addresses")
	//showRequests := flag.Bool("request", false, "Show a TOP of Requests if enabled")
	//showUpstreams := flag.Bool("upstreams", false, "Show a upstream if enabled")
	flag.Parse()

	//var report []LogString

	// Check file name
	if *fname == "" {
		fmt.Println("Error: Please enter a file name \n" +
			"./jsonlogparser -fname ./path/to/access.log")
	}
	if *howmuch == 0 {
		fmt.Println("Error: Please enter a file name \n" +
			"./jsonlogparser -fname ./path/to/access.log -howmuch 15")
	}
	c := *howmuch
	// Try to open file
	FName := *fname
	file, err := os.Open(FName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// prep - array of string
	var prep []string
	// Some magic from StackOverflow
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prep = append(prep, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(prep))
	res := prepareData(prep)
	rep := makeReport(res, c)
	fmt.Println(rep)
	fmt.Println(len(rep))
}

func prepareData(data []string) []LogString {
	var result []LogString

	for k,_ := range data {
		str := []byte(string(data[k]))
		res := LogString{}
		json.Unmarshal([]byte(str), &res)
		//fmt.Println(res.RemoteAddr)
		result = append(result, res)
	}
	return result
}

func makeReport(data []LogString, count int) []string {
	var res []string
	i := 0
	for _,v := range data {
		i += 1
		res = append(res, v.RemoteAddr)
	}
	return res[0:count]
}

