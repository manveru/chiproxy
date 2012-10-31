package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.Host, "manveru.name") {
			proxy := httputil.ReverseProxy{Director: manveruDotName}
			proxy.ServeHTTP(w, r)
		} else if strings.HasSuffix(r.Host, "nsans.eu") {
			proxy := httputil.ReverseProxy{Director: nsansDotEu}
			proxy.ServeHTTP(w, r)
		} else {
			w.WriteHeader(404)
		}
	})

	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}

func manveruDotName(r *http.Request) {
	server, err := url.Parse("http://localhost:8000" + r.RequestURI)
	if err != nil {
		log.Printf("%#v\n", r)
		log.Println("couldn't create manveru.name url", err)
	}
	r.URL = server
}

func nsansDotEu(r *http.Request) {
	server, err := url.Parse("http://localhost:9000" + r.RequestURI)
	if err != nil {
		log.Printf("%#v\n", r)
		log.Println("couldn't create nsans.eu url", err)
	}
	r.URL = server
}
