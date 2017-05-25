package entities

import (
	"encoding/xml"
	"fmt"
)

type Email struct {
	Where string `xml:"where,attr"`
	Addr  string
}
type Address struct {
	City, State string
}
type Person struct {
	XMLName xml.Name `xml:"Person"`
	Name    string   `xml:"FullName"`
	Phone   string
	Email   []Email
	Groups  []string `xml:"Group>Value"`
	Address
}
type Persons struct {
	XMLName xml.Name `xml:"Persons"`
	Persons []Person `xml:"Person"`
}

func ExampleUnmarshal() {

	data := `
	<Persons>
  		<Person>
  			<FullName>Grace R. Emlin</FullName>
  			<Company>Example Inc.</Company>
  			<Email where="home">
  				<Addr>gre@example.com</Addr>
  			</Email>
  			<Email where='work'>
  				<Addr>gre@work.com</Addr>
  			</Email>
  			<Group>
  				<Value>Friends</Value>
  				<Value>Squash</Value>
  			</Group>
  			<City>Hanga Roa</City>
  			<State>Easter Island</State>
  		</Person>
	</Persons>
  	`

	//v := Person{Name: "none", Phone: "none"}
	v := Persons{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	for _, person := range v.Persons {
		fmt.Printf("XMLName: %#v\n", person.XMLName)
		fmt.Printf("Name: %q\n", person.Name)
		fmt.Printf("Phone: %q\n", person.Phone)
		fmt.Printf("Email: %v\n", person.Email)
		fmt.Printf("Groups: %v\n", person.Groups)
		fmt.Printf("Address: %v\n", person.Address)
	}
}
func ParseXmlTestBookingAmount() {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns="http://schemas.xmlsoap.org/soap/envelope/">
		<Header>
			<Security d3p1:mustUnderstand="0" xmlns:d3p1="http://schemas.xmlsoap.org/soap/envelope/"
					xmlns="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd">
				<UsernameToken>
					<Username>user name</Username>
					<Password>password</Password>
				</UsernameToken>
			</Security>
			<MessageID xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">1197835</MessageID>
			<From xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">
				<Address>http://schemas.xmlsoap.org/ws/2004/12/addressing/role/anonymous</Address>
			</From>
		</Header>
		<Body>
			<OTA_HotelRateAmountNotifRQ xmlns="http://www.opentravel.org/OTA/2003/05">
				<RateAmountMessages HotelCode="13947">
					<RateAmountMessage>
						<StatusApplicationControl 
						InvTypeCode="STD_Twin" 
						RatePlanCode="BAR" 
						Start="2017-05-18" 
						End="2017-05-18" 
						Sun="False" 
						Mon="False" 
						Tue="False" 
						Weds="False" 
						Thur="True" 
						Fri="False" 
						Sat="False" />
						<Rates>
							<Rate>
								<BaseByGuestAmts>
									<BaseByGuestAmt AmountAfterTax="219.00" />
								</BaseByGuestAmts>
							</Rate>
						</Rates>
					</RateAmountMessage>
				</RateAmountMessages>
			</OTA_HotelRateAmountNotifRQ>
		</Body>
	</Envelope>	
  	`

	//v := Person{Name: "none", Phone: "none"}
	v := HRAEnvelope{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Print

	// Header
	fmt.Printf("Header:\n")
	fmt.Printf("Security - mustUnderstand:%q username:%q password:%q\n",
		v.HRAHeader.HRASecurity.D3p1MustUnderstand,
		v.HRAHeader.HRASecurity.HRAUsernameToken.HRAUsername,
		v.HRAHeader.HRASecurity.HRAUsernameToken.HRAPassword)
	fmt.Printf("MessageID - %q\n", v.HRAHeader.HRAMessageID)
	fmt.Printf("From - %q\n", v.HRAHeader.HRAFrom.Address)

	// Body

	fmt.Printf("\nBody:\n")

	for _, RateAmount := range v.HRABody.HotelRateAmount {
		for _, message := range RateAmount.HRARateAmountMessages {
			fmt.Printf("Message from hotel %q:\n", message.HotelCode)

			appControl := message.HRARateAmountMessage[0].HRAStatusApplicationControl
			fmt.Printf("ApplicationControl:\n InvTypeCode:%q\nRatePlanCode:%q\nStart:%q\nEnd:%q\nSun:%q\nMon:%q\nTue:%q\nWeds:%q\nThur:%q\nFri:%q\nSat:%q\n",
				appControl.InvTypeCode,
				appControl.RatePlanCode,
				appControl.Start,
				appControl.End,
				appControl.Sun,
				appControl.Mon,
				appControl.Tue,
				appControl.Weds,
				appControl.Thur,
				appControl.Fri,
				appControl.Sat)

			fmt.Printf("Rates:")
			for _, rate := range message.HRARateAmountMessage[0].HRARates.HRARate {
				fmt.Printf("AmountAfterTax:", rate.HRABaseByGuestAmts.HRABaseByGuestAmt.AmountAfterTax)
			}

			fmt.Printf("\n")
		}
	}
	// for _, person := range v.Persons {
	// 	fmt.Printf("XMLName: %#v\n", person.XMLName)
	// 	fmt.Printf("Name: %q\n", person.Name)
	// 	fmt.Printf("Phone: %q\n", person.Phone)
	// 	fmt.Printf("Email: %v\n", person.Email)
	// 	fmt.Printf("Groups: %v\n", person.Groups)
	// 	fmt.Printf("Address: %v\n", person.Address)
	// }

}
