package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		headers := "Headers=>"
		for k, v := range r.Header {
			headers += " | " + k + " : " + strings.Join(v, ",")
		}

		log.Printf(
			"%s\t%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
			headers,
		)

		fmt.Printf(
			"%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
