package main

import (
	"GoApi/common"
	"log"
	"net/http"
	"strconv"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(common.ServicePortNumber), router))
}
