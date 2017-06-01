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
			fmt.Printf("ApplicationControl:\nInvTypeCode:%q\nRatePlanCode:%q\nStart:%q\nEnd:%q\nSun:%q\nMon:%q\nTue:%q\nWeds:%q\nThur:%q\nFri:%q\nSat:%q\n",
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
}
func ParseXmlTestInvCount() {
	data := `
	<?xml version="1.0" encoding="utf-8"?>
	<SOAP-ENV:Envelope 
		xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/"
		xmlns:ns1="http://www.opentravel.org/OTA/2003/05" 
		xmlns:ns2="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" 
		xmlns:ns3="http://schemas.xmlsoap.org/ws/2004/08/addressing">
		<SOAP-ENV:Header>
		<ns2:Security SOAP-ENV:mustUnderstand="1">
			<UsernameToken>
				<Username>user name</Username>
				<Password>password</Password>
			</UsernameToken>
		</ns2:Security>
		<ns3:MessageID>uuid:41839242-45bc-102d-b85d-0030487c664c</ns3:MessageID>
		<ns3:From>
			<ns3:Address>http://schemas.xmlsoap.org/ws/2004/12/addressing/role/anonymous</ns3:Address>
		</ns3:From>
		</SOAP-ENV:Header>
		<SOAP-ENV:Body>
		<OTA_HotelInvCountNotifRQ TimeStamp = "2016-11-14T15:43:25"  Version="1.000">
			<Inventories  HotelCode = "13947" HotelName = "">
			<Inventory>
				<ns1:StatusApplicationControl  InvTypeCode = "STD_Triple"  Start = "2017-08-30"  End = "2017-08-30"  />
				<InvCounts>
				<InvCount CountType = "1"  Count = "5" />
				</InvCounts>
			</Inventory>
			</Inventories>
		</OTA_HotelInvCountNotifRQ>
		</SOAP-ENV:Body>
	</SOAP-ENV:Envelope>
  	`

	//v := Person{Name: "none", Phone: "none"}
	v := HICEnvelope{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Print

	// Header
	fmt.Printf("Header:\n")
	fmt.Printf("Security - mustUnderstand:%q username:%q password:%q\n",
		v.HICHeader.HICSecurity.D3p1MustUnderstand,
		v.HICHeader.HICSecurity.HICUsernameToken.HICUsername,
		v.HICHeader.HICSecurity.HICUsernameToken.HICPassword)
	fmt.Printf("MessageID - %q\n", v.HICHeader.HICMessageID)
	fmt.Printf("From - %q\n", v.HICHeader.HICFrom.Address)

	// Body

	fmt.Printf("\nBody:\n")

	for _, invertList := range v.HICBody.HotelInvCountNotifRQ.HICInventoriesList {
		fmt.Printf("Message from hotel: %q %q\n", invertList.HotelCode, invertList.HotelName)

		for _, invertory := range invertList.HICInventories {
			appControl := invertory.StatusApplicationControl
			fmt.Printf("ApplicationControl:\nInvTypeCode:%q\nStart:%q\nEnd:%q",
				appControl.InvTypeCode,
				appControl.Start,
				appControl.End)

			fmt.Printf("Counts:")
			for _, count := range invertory.HICInvCounts.HICInvCount {
				fmt.Printf("CountType: %q  Count: %q", count.CountType, count.Count)
			}

			fmt.Printf("\n")
		}
	}
}
func ParseXmlTestBookingRules() {
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
			<MessageID xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">1197730</MessageID>
			<From xmlns="http://schemas.xmlsoap.org/ws/2004/08/addressing">
				<Address>http://schemas.xmlsoap.org/ws/2004/12/addressing/role/anonymous</Address>
			</From>
		</Header>
		<Body>
			<OTA_HotelBookingRuleNotifRQ TimeStamp="2017-38-14T02:05:47" xmlns="http://www.opentravel.org/OTA/2003/05">
				<RuleMessages HotelCode="13947">
					<RuleMessage>
						<StatusApplicationControl InvTypeCode="STD_Twin" RatePlanCode="BAR" Start="2017-05-19" End="2017-05-19" Sun="True" Mon="True" Tue="True" Weds="True" Thur="True" Fri="True" Sat="True" />
						<BookingRules>
							<BookingRule Start="2017-05-19" End="2017-05-19">
								<LengthsOfStay ArrivalDateBased="true">
									<LengthOfStay Time="2" TimeUnit="Day" MinMaxMessageType="SetMinLOS" />
								</LengthsOfStay>
							</BookingRule>
							<BookingRule Start="2017-05-19" End="2017-05-19">
								<RestrictionStatus Restriction="Master" Status="Open" />
							</BookingRule>
							<BookingRule Start="2017-05-19" End="2017-05-19">
								<RestrictionStatus Restriction="Arrival" Status="Open" />
							</BookingRule>
							<BookingRule Start="2017-05-19" End="2017-05-19">
								<RestrictionStatus Restriction="Departure" Status="Open" />
							</BookingRule>
						</BookingRules>
					</RuleMessage>
				</RuleMessages>
			</OTA_HotelBookingRuleNotifRQ>
		</Body>
	</Envelope>
  	`

	//v := Person{Name: "none", Phone: "none"}
	v := HBREnvelope{}

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	// Print

	// Header
	fmt.Printf("Header:\n")
	fmt.Printf("Security - mustUnderstand:%q username:%q password:%q\n",
		v.HBRHeader.HBRSecurity.D3p1MustUnderstand,
		v.HBRHeader.HBRSecurity.HBRUsernameToken.HBRUsername,
		v.HBRHeader.HBRSecurity.HBRUsernameToken.HBRPassword)
	fmt.Printf("MessageID - %q\n", v.HBRHeader.HBRMessageID)
	fmt.Printf("From - %q\n", v.HBRHeader.HBRFrom.Address)

	// Body

	fmt.Printf("\nBody:\n")

	for _, ruleMessage := range v.HBRBody.HotelBookingRuleNotifRQ.HRBRuleMessages {
		fmt.Printf("Message from hotel: %q \n", ruleMessage.HotelCode)

		for _, ruleMessageItem := range ruleMessage.HRBMessagesItems {
			appControl := ruleMessageItem.StatusApplicationControl
			fmt.Printf("ApplicationControl:\nInvTypeCode:%q\nRatePlanCode:%q\nStart:%q\nEnd:%q\nSun:%q\nMon:%q\nTue:%q\nWeds:%q\nThur:%q\nFri:%q\nSat:%q\n",
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

			fmt.Printf("Booking Rules:")
			for _, rule := range ruleMessageItem.HRBBookingRules.HRBBookingRuleItems {
				fmt.Printf("Booking Rule: %q  End: %q", rule.Start, rule.End)
				if rule.HRBLengthsOfStay != nil {
					fmt.Printf(" HRBLengthsOfStay:  ArrivalDateBased:%q  Time:%q TimeUnit:%q MinMaxMessageType:%q",
						rule.HRBLengthsOfStay.ArrivalDateBased,
						rule.HRBLengthsOfStay.HRBLengthOfStay.Time,
						rule.HRBLengthsOfStay.HRBLengthOfStay.TimeUnit,
						rule.HRBLengthsOfStay.HRBLengthOfStay.MinMaxMessageType)
				}
				if rule.HRBRestrictionStatus != nil {
					fmt.Printf(" HRBRestrictionStatus:  Restriction:%q  Restriction:%q",
						rule.HRBRestrictionStatus.Restriction,
						rule.HRBRestrictionStatus.Status)
				}
				fmt.Printf("\n")
			}

			fmt.Printf("\n")
		}
	}
}
