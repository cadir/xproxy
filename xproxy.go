// TODO: Add SSL support and certificate
package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var kLocalUrl string

func copyHeaders(dst, src http.Header) {
	for k, _ := range dst {
		dst.Del(k)
	}
	for k, vs := range src {
		for _, v := range vs {
			if k == "Location" {
				dst.Add(k, strings.Replace(v, "https://www.xpiron.com", kLocalUrl, 1))
			} else {
				dst.Add(k, v)
			}
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%v %v", r.Method, r.URL)

	tr := &http.Transport{}
	req, err := http.NewRequest(r.Method, "https://www.xpiron.com"+r.URL.String(), r.Body)
	req.Header = r.Header
	resp, err := tr.RoundTrip(req)
	if err != nil {
		log.Printf("%v", err)
	}
	copyHeaders(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	defer resp.Body.Close()
	nr, err := io.Copy(w, resp.Body)
	log.Printf("Wrote %v bytes to client error=%v", nr, err)
}

func main() {
	port := ""
	kLocalUrl = "http://ec2-54-201-152-136.us-west-2.compute.amazonaws.com"
	hostname, _ := os.Hostname()
	if strings.Contains(hostname, "Mac") {
		kLocalUrl = "http://localhost:8081"
		port = ":8081"
	}
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
