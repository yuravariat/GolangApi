package controllers

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"GoApi/entities"
	"encoding/xml"

	"strings"

	"github.com/go-errors/errors"
)

func PushRecieve(w http.ResponseWriter, r *http.Request) {

	responseType := "OTA_HotelInvCountNotifRS"
	success := true

	err := func() error {

		w.Header().Set("Content-Type", "text/xml")

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
		uuidStr := r.Context().Value("uuid").(string)
		go pushRecieveWriteFile(body, uuidStr)
		//fmt.Fprint(w, string(body))

		if strings.Contains(string(body), "OTA_HotelInvCountNotifRQ") {
			responseType = "OTA_HotelInvCountNotifRS"
		} else if strings.Contains(string(body), "OTA_HotelBookingRuleNotifRQ") {
			responseType = "OTA_HotelBookingRuleNotifRS"
		} else if strings.Contains(string(body), "OTA_HotelRateAmountNotifRQ") {
			responseType = "OTA_HotelRateAmountNotifRS"
		}

		return nil
	}()

	if err != nil {
		w.WriteHeader(400)
		serr := errors.Wrap(err, 1)
		log.Printf(serr.ErrorStack())
		io.WriteString(w, err.Error())
		success = false
	}

	responseBody := "<Success/>"
	if !success {
		responseBody = `<Errors>
		<Error Type="Unknown">
			` + err.Error() + `
		<Error>
		</Errors>`
	}
	response := `<Envelope 
		xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" 
		xmlns:ns1="http://schemas.xmlsoap.org/ws/2004/08/addressing"
		xmlns:ns2="http://www.w3.org/2005/08/addressing" 
		xmlns:ns3="http://www.opentravel.org/OTA/2003/05">
	<Header>
		<MessageID>uuid:82e8a910-e7f1-2044-d8f0-d69493a35126</MessageID>
		<RelatesTo>0</RelatesTo>
	</Header>
	<Body>
		<` + responseType + ` TimeStamp="` + time.Now().Format(time.RFC3339) + `">
		` + responseBody + `
		</` + responseType + `>
	</Body>
	</Envelope>`

	fmt.Fprint(w, response)
}
func pushRecieveWriteFile(content []byte, uuidStr string) {

	f, er := os.Create("./logs/requests/recived_" + time.Now().Format("2006-01-02_15-04-04_05Z07.00") + "-" + uuidStr + ".xml")
	if er != nil {
		serr := errors.Wrap(er, 1)
		log.Printf(serr.ErrorStack())
	}
	defer f.Close()

	f.Write(content)
}
func parseHRAXML(content []byte, xmlObj *entities.HRAEnvelope) error {

	if er := xml.Unmarshal(content, &xmlObj); er != nil {
		return er
	}
	return nil
}
