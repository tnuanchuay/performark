package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
	"errors"
)

func main(){
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Println("redirect", req.URL)
			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		},
	}

	req, err := http.NewRequest("GET", "https://sis-phuket5.psu.ac.th/WebRegist2005/Login.aspx", nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36")
	//req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
	//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	//req.Header.Add("Accept-Language", "th-TH,th;q=0.8,en;q=0.6")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	cookie := resp.Cookies()[0]
	fmt.Println(cookie)
	for {
		if resp.StatusCode == 302{
			req, err := http.NewRequest("GET", "https://sis-phuket5.psu.ac.th/WebRegist2005/Login.aspx", nil)
			if err != nil {
				fmt.Println(err)
			}

			req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/52.0.2743.116 Safari/537.36")
			//req.Header.Add("Connection", "keep-alive")
			//req.Header.Add("Accept-Encoding", "gzip, deflate, sdch")
			//req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
			//req.Header.Add("Accept-Language", "th-TH,th;q=0.8,en;q=0.6")
			req.AddCookie(cookie)
			resp, err = client.Do(req)
		}else{
			break
		}
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {

	}
	defer resp.Body.Close()
	res := string(buf)

	fmt.Println(res)
}
