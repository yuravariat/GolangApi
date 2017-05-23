package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-errors/errors"
)

func PushRecieve(w http.ResponseWriter, r *http.Request) {

	err := func() error {

		// Check that the server actually sent compressed data
		// var reader io.ReadCloser
		// var er error
		// switch r.Header.Get("Content-Encoding") {
		// case "gzip":
		// 	reader, er = gzip.NewReader(r.Body)
		// 	if er != nil {
		// 		return er
		// 	}
		// 	defer reader.Close()
		// default:
		// 	reader = r.Body
		// }

		//body, er := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		body, er := ioutil.ReadAll(r.Body)
		//body, er := ioutil.ReadAll(reader)
		if er != nil {
			return er
		}

		// write file in background thread
		go pushRecieveWriteFile(body)
		//fmt.Fprint(w, string(body))

		return nil
	}()

	if err != nil {
		w.WriteHeader(400)
		serr := errors.Wrap(err, 1)
		log.Printf(serr.ErrorStack())
		io.WriteString(w, err.Error())
		return
	}

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	fmt.Fprint(w, "{\"success\":\"true\"}")
}
func pushRecieveWriteFile(content []byte) {

	f, er := os.Create("./logs/requests/recived_" + time.Now().Format("2006-01-02_15-04-04_05Z07.00") + ".xml")
	if er != nil {
		serr := errors.Wrap(er, 1)
		log.Printf(serr.ErrorStack())
	}
	defer f.Close()

	f.Write(content)
}
