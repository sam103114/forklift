package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

//ServeProxy : Send request to origin, modify request to be sent
func ServeProxy(origin string, w http.ResponseWriter, r *http.Request) {
	url, err := url.Parse(origin)
	if err != nil {
		log.Fatal(err)
	}

	newProxy := httputil.NewSingleHostReverseProxy(url)
	r.Host = url.Host
	r.URL.Host = url.Host
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	newProxy.ServeHTTP(w, r)
}

//Start proxy
func Start() {
	fmt.Println("Proxy Started")
	http.HandleFunc("/", redirectRequest)
	http.ListenAndServe(":"+Bal.Port, nil)
}

//Get next host to redirect to, forward request to proxy
func redirectRequest(w http.ResponseWriter, r *http.Request) {
	requestDest := selectRedirectURL()
	ServeProxy(requestDest, w, r)
}
