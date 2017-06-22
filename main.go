package main

import (
	"GoApi/MSMQ"
)

func main() {

	//MSMQ.WriteQueue("localhost", "queue", "test", "test", "This is body")
	results := MSMQ.ReadQueue("localhost", "queue", "test", "")

	if results != nil {
		for _, message := range results {
			println(message.String())
		}
	}

	// f, err := os.OpenFile("./logs/trace/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer f.Close()
	// log.SetOutput(f)

	// //entities.ParseXmlTestBookingAmount()
	// //entities.ParseXmlTestInvCount()
	// //entities.ParseXmlTestBookingRules()

	// router := NewRouter()

	// log.Printf("Server started")
	// if err := http.ListenAndServe(":"+strconv.Itoa(common.ServicePortNumber), router); err != nil {
	// 	log.Fatal(err)
	// }
}
