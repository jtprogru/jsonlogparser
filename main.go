package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sort"
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
	fName := flag.String("fname", "/var/log/nginx/access.log", "A path to access.log Nginx")
	howmuch := flag.Int("howmuch", 15, "Show a TOP of remote addresses")
	//showRequests := flag.Bool("request", false, "Show a TOP of Requests if enabled")
	//showUpstreams := flag.Bool("upstreams", false, "Show a upstream if enabled")
	flag.Parse()
	threadCount := runtime.NumCPU()

	strCh := make(chan string, threadCount)
	parsedCh := make(chan *LogString, threadCount)

	go readFile(*fName, strCh)
	go parseJson(strCh, parsedCh)

	makeReport(parsedCh, *howmuch)
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

type Pair struct {
	Key   string
	Count int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Count > p[j].Count }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func makeReport(in <-chan *LogString, topRemoteAddr int) {
	remoteAddr := make(map[string]Pair)
	for ls := range in {
		pl := remoteAddr[ls.RemoteAddr]
		pl.Key = ls.RemoteAddr
		pl.Count++
		remoteAddr[ls.RemoteAddr] = pl

	}
	raPL := make(PairList, len(remoteAddr))
	i := 0
	for _, v := range remoteAddr {
		raPL[i] = v
		i++
	}
	sort.Sort(raPL)
	for i, v := range raPL {
		if i > topRemoteAddr-1 {
			break
		}
		fmt.Printf("%d\t\t%s\t%d\n", i+1, v.Key, v.Count)
	}
}
