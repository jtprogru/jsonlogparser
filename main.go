package main

import (
	"fmt"
	//"io"
	"os"
	"net"
	"path"
	//"strings"
	//"time"
)

const iso8601DateFormat = "2006-01-02T15:04:05-07:00"
//commonlogDateFormat := "2/Jan/2006:15:04:05 -0700"
//parsedDate, err := time.Parse(commonlogDateFormat, "18/Oct/2014:08:53:14 +0200")
//if err != nil {
//	fmt.Println(err)
//	return
//}
//fmt.Println(parsedDate.Format(iso8601DateFormat))

var logstring struct {
	// TODO:
	//timestamp time.RFC3339  //time_iso8601
	remote_addr net.IPAddr
	cookie_bar string
	set_cookie string
	body_bytes_sent int
	status int
	request string
	url string
	request_method string
	upstream net.IPAddr
	response_time float32
	http_referrer string
	http_user_agent string
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
