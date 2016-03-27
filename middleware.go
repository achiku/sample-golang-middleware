package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}
	return http.HandlerFunc(fn)
}

func recoverMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %s\n", err)
				w.Header().Set("Content-type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				res := map[string]string{"message": "error"}
				json.NewEncoder(w).Encode(res)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

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

// AppendMiddleware append msg
func AppendMiddleware(handler http.Handler, msg string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		w.Write([]byte(msg))
	}
	return http.HandlerFunc(fn)
}
