package main

import (
	"GoApi/common"
	"log"
	"net/http"
	"strconv"
)

func mainRun() {

	//f, err := os.OpenFile("./logs/trace/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//if err != nil {
	//	log.Fatalf("error opening file: %v", err)
	//}
	//defer f.Close()
	//log.SetOutput(f)

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(common.ServicePortNumber), router))
}
