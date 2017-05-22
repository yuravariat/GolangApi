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
		body, er := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if er != nil {
			return er
		}

		f, er := os.Create("./logs/requests/recived_" + time.Now().Format("2006-01-02_15-04-04_05Z07.00") + ".xml")
		if er != nil {
			serr := errors.Wrap(er, 1)
			log.Printf(serr.ErrorStack())
			return er
		}
		defer f.Close()

		f.Write(body)
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
