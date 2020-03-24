package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
)

/*
	Reverse Proxy Logic
*/

// Serve a reverse proxy for a given url
func serveReverseProxy(target string, res http.ResponseWriter, req *http.Request) {

}

// Given a request send it to the appropriate url
func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {

	//incomingurl := req.URL.Port

	log.Printf(req.Host)

	re := regexp.MustCompile(`^[a-z0-9_]+`)
	host := re.FindString(req.Host)
	port := ""
	matched, _ := regexp.MatchString("_mailhog$", host)
	if matched {
		port = ":8025"
	}

	url, _ := url.Parse("http://" + host + port)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	//TODO Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}

/*
	Entry
*/

func main() {

	// start server
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
