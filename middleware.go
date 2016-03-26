package main

import (
	"log"
	"net/http"
)

// SingleHost middleware struct
type SingleHost struct {
	handler     http.Handler
	allowedHost string
}

// NewSingleHost create SingleHost middleware
func NewSingleHost(handler http.Handler, allowedHost string) *SingleHost {
	return &SingleHost{handler: handler, allowedHost: allowedHost}
}

// ServeHTTP server http
func (s *SingleHost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	host := r.Host
	log.Printf("host: %s", host)
	if host == s.allowedHost {
		s.handler.ServeHTTP(w, r)
	} else {
		log.Printf("host [%s] is allowed host [%s]", host, s.allowedHost)
		w.WriteHeader(http.StatusForbidden)
	}
}

// SingleHost2 single host 2
func SingleHost2(handler http.Handler, allowedHost string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		log.Printf("host: %s", host)
		if host == allowedHost {
			handler.ServeHTTP(w, r)
		} else {
			log.Printf("host [%s] is allowed host [%s]", host, allowedHost)
			w.WriteHeader(http.StatusForbidden)
		}
	}
	return http.HandlerFunc(fn)
}
