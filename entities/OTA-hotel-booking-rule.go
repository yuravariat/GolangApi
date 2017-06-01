package entities

import "encoding/xml"

///////////////////////////////// Wrapper //////////////////////////////////////

// HBREnvelope ...
type HBREnvelope struct {
	SoapEnv   string     `xml:"xmlns SOAP-ENV,attr"  json:",omitempty"`
	Ns1       string     `xml:"xmlns ns1,attr"  json:",omitempty"`
	Ns2       string     `xml:"xmlns ns2,attr"  json:",omitempty"`
	Ns3       string     `xml:"xmlns ns3,attr"  json:",omitempty"`
	HBRHeader *HBRHeader `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header,omitempty" json:"Header,omitempty"`
	HBRBody   *HBRBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName   xml.Name   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

//////////////////////////////// Header part ///////////////////////////////////

// HBRHeader ...
type HBRHeader struct {
	HBRFrom      *HBRFrom     `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing From,omitempty" json:"From,omitempty"`
	HBRMessageID string       `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing MessageID,omitempty" json:"MessageID,omitempty"`
	HBRSecurity  *HBRSecurity `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security,omitempty" json:"Security,omitempty"`
	XMLName      xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header,omitempty" json:"Header,omitempty"`
}

// HBRSecurity ...
type HBRSecurity struct {
	D3p1               string            `xml:"xmlns d3p1,attr"  json:",omitempty"`
	D3p1MustUnderstand string            `xml:"http://schemas.xmlsoap.org/soap/envelope/ mustUnderstand,attr"  json:",omitempty"`
	HBRUsernameToken   *HBRUsernameToken `xml:"UsernameToken,omitempty" json:"UsernameToken,omitempty"`
	XMLName            xml.Name          `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security,omitempty" json:"Security,omitempty"`
}

// HBRUsernameToken ...
type HBRUsernameToken struct {
	HBRPassword string   `xml:"Password,omitempty" json:"Password,omitempty"`
	HBRUsername string   `xml:"Username,omitempty" json:"Username,omitempty"`
	XMLName     xml.Name `xml:"UsernameToken,omitempty" json:"UsernameToken,omitempty"`
}

// HBRFrom ...
type HBRFrom struct {
	Address string   `xml:"Address,omitempty" json:"Address,omitempty"`
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing From,omitempty" json:"Header,omitempty"`
}

// HBRAddress ...
type HBRAddress struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address,omitempty" json:"Address,omitempty"`
}

//////////////////////////////////// Body part //////////////////////////////////////

// HBRBody ...
type HBRBody struct {
	HotelBookingRuleNotifRQ *HotelBookingRuleNotifRQ `xml:"http://www.opentravel.org/OTA/2003/05 OTA_HotelBookingRuleNotifRQ,omitempty" json:"OTA_HotelBookingRuleNotifRQ,omitempty"`
	XMLName                 xml.Name                 `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

// HotelBookingRuleNotifRQ ...
type HotelBookingRuleNotifRQ struct {
	TimeStamp       string             `xml:" TimeStamp,attr"  json:",omitempty"`
	Xmlns           string             `xml:" xmlns,attr"  json:",omitempty"`
	HRBRuleMessages []*HRBRuleMessages `xml:"http://www.opentravel.org/OTA/2003/05 RuleMessages,omitempty" json:"RuleMessages,omitempty"`
	XMLName         xml.Name           `xml:"http://www.opentravel.org/OTA/2003/05 OTA_HotelBookingRuleNotifRQ,omitempty" json:"OTA_HotelBookingRuleNotifRQ,omitempty"`
}

// HRBRuleMessages ...
type HRBRuleMessages struct {
	HotelCode        string            `xml:" HotelCode,attr"  json:",omitempty"`
	HRBMessagesItems []*HRBRuleMessage `xml:"http://www.opentravel.org/OTA/2003/05 RuleMessage,omitempty" json:"RuleMessage,omitempty"`
	XMLName          xml.Name          `xml:"http://www.opentravel.org/OTA/2003/05 RuleMessages,omitempty" json:"RuleMessages,omitempty"`
}

// HRBRuleMessage ...
type HRBRuleMessage struct {
	HRBBookingRules          *HRBBookingRules             `xml:"http://www.opentravel.org/OTA/2003/05 BookingRules,omitempty" json:"BookingRules,omitempty"`
	StatusApplicationControl *HRBStatusApplicationControl `xml:"http://www.opentravel.org/OTA/2003/05 StatusApplicationControl,omitempty" json:"StatusApplicationControl,omitempty"`
	XMLName                  xml.Name                     `xml:"http://www.opentravel.org/OTA/2003/05 RuleMessage,omitempty" json:"RuleMessage,omitempty"`
}

// HRBStatusApplicationControl ...
type HRBStatusApplicationControl struct {
	End          string   `xml:" End,attr"  json:",omitempty"`
	Fri          string   `xml:" Fri,attr"  json:",omitempty"`
	InvTypeCode  string   `xml:" InvTypeCode,attr"  json:",omitempty"`
	Mon          string   `xml:" Mon,attr"  json:",omitempty"`
	RatePlanCode string   `xml:" RatePlanCode,attr"  json:",omitempty"`
	Sat          string   `xml:" Sat,attr"  json:",omitempty"`
	Start        string   `xml:" Start,attr"  json:",omitempty"`
	Sun          string   `xml:" Sun,attr"  json:",omitempty"`
	Thur         string   `xml:" Thur,attr"  json:",omitempty"`
	Tue          string   `xml:" Tue,attr"  json:",omitempty"`
	Weds         string   `xml:" Weds,attr"  json:",omitempty"`
	XMLName      xml.Name `xml:"http://www.opentravel.org/OTA/2003/05 StatusApplicationControl,omitempty" json:"StatusApplicationControl,omitempty"`
}

// HRBBookingRules ...
type HRBBookingRules struct {
	HRBBookingRuleItems []*HRBBookingRule `xml:"http://www.opentravel.org/OTA/2003/05 BookingRule,omitempty" json:"BookingRule,omitempty"`
	XMLName             xml.Name          `xml:"http://www.opentravel.org/OTA/2003/05 BookingRules,omitempty" json:"BookingRules,omitempty"`
}

// HRBBookingRule ...
type HRBBookingRule struct {
	End                  string                `xml:" End,attr"  json:",omitempty"`
	Start                string                `xml:" Start,attr"  json:",omitempty"`
	HRBLengthsOfStay     *HRBLengthsOfStay     `xml:"http://www.opentravel.org/OTA/2003/05 LengthsOfStay,omitempty" json:"LengthsOfStay,omitempty"`
	HRBRestrictionStatus *HRBRestrictionStatus `xml:"http://www.opentravel.org/OTA/2003/05 RestrictionStatus,omitempty" json:"RestrictionStatus,omitempty"`
	XMLName              xml.Name              `xml:"http://www.opentravel.org/OTA/2003/05 BookingRule,omitempty" json:"BookingRule,omitempty"`
}

// HRBLengthsOfStay ...
type HRBLengthsOfStay struct {
	ArrivalDateBased string                 `xml:" ArrivalDateBased,attr"  json:",omitempty"`
	HRBLengthOfStay  *HRBLengthOfStayNested `xml:"http://www.opentravel.org/OTA/2003/05 LengthOfStay,omitempty" json:"LengthOfStay,omitempty"`
	XMLName          xml.Name               `xml:"http://www.opentravel.org/OTA/2003/05 LengthsOfStay,omitempty" json:"LengthsOfStay,omitempty"`
}

// HRBLengthOfStayNested ...
type HRBLengthOfStayNested struct {
	MinMaxMessageType string   `xml:" MinMaxMessageType,attr"  json:",omitempty"`
	Time              string   `xml:" Time,attr"  json:",omitempty"`
	TimeUnit          string   `xml:" TimeUnit,attr"  json:",omitempty"`
	XMLName           xml.Name `xml:"http://www.opentravel.org/OTA/2003/05 LengthOfStay,omitempty" json:"LengthOfStay,omitempty"`
}

// HRBRestrictionStatus ...
type HRBRestrictionStatus struct {
	Restriction string   `xml:" Restriction,attr"  json:",omitempty"`
	Status      string   `xml:" Status,attr"  json:",omitempty"`
	XMLName     xml.Name `xml:"http://www.opentravel.org/OTA/2003/05 RestrictionStatus,omitempty" json:"RestrictionStatus,omitempty"`
}
