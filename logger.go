package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		requestGUID := uuid.NewV4().String()

		headers := "Headers=>"
		for k, v := range r.Header {
			headers += " | " + k + " : " + strings.Join(v, ",")
		}

		requestStr := fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n",
			r.Method,
			r.RequestURI,
			name,
			requestGUID,
			headers,
		)

		log.Printf(requestStr)
		fmt.Printf(requestStr)

		ctx := context.WithValue(r.Context(), "uuid", requestGUID)
		inner.ServeHTTP(w, r.WithContext(ctx))

		responseStr := fmt.Sprintf("%s\t%s\t%s\t%s\t took [%s]\n",
			r.Method,
			r.RequestURI,
			name,
			requestGUID,
			time.Since(start),
		)

		log.Printf(responseStr)
		fmt.Printf(responseStr)
	})
}
