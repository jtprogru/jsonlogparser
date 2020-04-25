package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
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
	// Check length of args
	if len(os.Args) < 2 { //first arg is binary name
		fmt.Println("ERROR: lost path to logfile \n" +
			"./jsonlogparser ./path/to/access.log")
	}
	fName := os.Args[1]
	// Try to open file

	threadCount := runtime.NumCPU()

	strCh := make(chan string, threadCount)
	parsedCh := make(chan *LogString, threadCount)

	go readFile(fName, strCh)
	go parseJson(strCh, parsedCh)

	someLogic(parsedCh)
}

func readFile(fName string, ch chan string) {
	defer close(ch)
	file, err := os.Open(fName)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
	if err = scanner.Err(); err != nil {
		log.Printf("error: %v", err)
	}
}

func parseJson(in <-chan string, out chan *LogString) {
	defer close(out)
	for str := range in {
		res := new(LogString)
		err := json.Unmarshal([]byte(str), res)
		if err != nil {
			log.Printf("Error on parse '%s': %v", str, err)
		} else {
			out <- res
		}
	}
}

func someLogic(in <-chan *LogString) {
	for ls := range in {
		log.Printf("%v", ls)
	}
}
