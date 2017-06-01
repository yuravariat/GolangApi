package entities

import "encoding/xml"

///////////////////////////////// Wrapper //////////////////////////////////////

// HRAEnvelope ...
type HRAEnvelope struct {
	Xmlns     string     `xml:" xmlns,attr"  json:",omitempty"`
	Xsd       string     `xml:"xmlns xsd,attr"  json:",omitempty"`
	Xsi       string     `xml:"xmlns xsi,attr"  json:",omitempty"`
	HRAHeader *HRAHeader `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header,omitempty" json:"Header,omitempty"`
	HRABody   *HRABody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName   xml.Name   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

//////////////////////////////// Header part ///////////////////////////////////

// HRAHeader ...
type HRAHeader struct {
	HRAFrom      *HRAFrom     `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing From,omitempty" json:"From,omitempty"`
	HRAMessageID string       `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing MessageID,omitempty" json:"MessageID,omitempty"`
	HRASecurity  *HRASecurity `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security,omitempty" json:"Security,omitempty"`
	XMLName      xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header,omitempty" json:"Header,omitempty"`
}

// HRASecurity ...
type HRASecurity struct {
	D3p1               string            `xml:"xmlns d3p1,attr"  json:",omitempty"`
	D3p1MustUnderstand string            `xml:"http://schemas.xmlsoap.org/soap/envelope/ mustUnderstand,attr"  json:",omitempty"`
	HRAUsernameToken   *HRAUsernameToken `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd UsernameToken,omitempty" json:"UsernameToken,omitempty"`
	XMLName            xml.Name          `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security,omitempty" json:"Security,omitempty"`
}

// HRAUsernameToken ...
type HRAUsernameToken struct {
	HRAPassword string   `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Password,omitempty" json:"Password,omitempty"`
	HRAUsername string   `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Username,omitempty" json:"Username,omitempty"`
	XMLName     xml.Name `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd UsernameToken,omitempty" json:"UsernameToken,omitempty"`
}

// HRAFrom ...
type HRAFrom struct {
	Address string   `xml:"Address,omitempty" json:"Address,omitempty"`
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing From,omitempty" json:"Header,omitempty"`
}

// HRAAddress ...
type HRAAddress struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address,omitempty" json:"Address,omitempty"`
}

//////////////////////////////////// Body part //////////////////////////////////////

// HRABody ...
type HRABody struct {
	HotelRateAmount []*HotelRateAmount `xml:"http://www.opentravel.org/OTA/2003/05 OTA_HotelRateAmountNotifRQ,omitempty" json:"OTA_HotelRateAmountNotifRQ,omitempty"`
	XMLName         xml.Name           `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

// HotelRateAmount ...
type HotelRateAmount struct {
	HRARateAmountMessages []*HRARateAmountMessages `xml:"http://www.opentravel.org/OTA/2003/05 RateAmountMessages,omitempty" json:"RateAmountMessages,omitempty"`
	XMLName               xml.Name                 `xml:"http://www.opentravel.org/OTA/2003/05 OTA_HotelRateAmountNotifRQ,omitempty" json:"OTA_HotelRateAmountNotifRQ,omitempty"`
}

// HRARateAmountMessages ...
type HRARateAmountMessages struct {
	HotelCode            string                  `xml:" HotelCode,attr"  json:",omitempty"`
	HRARateAmountMessage []*HRARateAmountMessage `xml:"http://www.opentravel.org/OTA/2003/05 RateAmountMessage,omitempty" json:"RateAmountMessage,omitempty"`
	XMLName              xml.Name                `xml:"http://www.opentravel.org/OTA/2003/05 RateAmountMessages,omitempty" json:"RateAmountMessages,omitempty"`
}

// HRARateAmountMessage ...
type HRARateAmountMessage struct {
	HRARates                    *HRARates                    `xml:"http://www.opentravel.org/OTA/2003/05 Rates,omitempty" json:"Rates,omitempty"`
	HRAStatusApplicationControl *HRAStatusApplicationControl `xml:"http://www.opentravel.org/OTA/2003/05 StatusApplicationControl,omitempty" json:"StatusApplicationControl,omitempty"`
	XMLName                     xml.Name                     `xml:"http://www.opentravel.org/OTA/2003/05 RateAmountMessage,omitempty" json:"RateAmountMessage,omitempty"`
}

// HRAStatusApplicationControl ...
type HRAStatusApplicationControl struct {
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

// HRARates ...
type HRARates struct {
	HRARate []*HRARate `xml:"http://www.opentravel.org/OTA/2003/05 Rate,omitempty" json:"Rate,omitempty"`
	XMLName xml.Name   `xml:"http://www.opentravel.org/OTA/2003/05 Rates,omitempty" json:"Rates,omitempty"`
}

// HRARate ...
type HRARate struct {
	HRABaseByGuestAmts *HRABaseByGuestAmts `xml:"http://www.opentravel.org/OTA/2003/05 BaseByGuestAmts,omitempty" json:"BaseByGuestAmts,omitempty"`
	XMLName            xml.Name            `xml:"http://www.opentravel.org/OTA/2003/05 Rate,omitempty" json:"Rate,omitempty"`
}

// HRABaseByGuestAmts ...
type HRABaseByGuestAmts struct {
	HRABaseByGuestAmt *HRABaseByGuestAmt `xml:"http://www.opentravel.org/OTA/2003/05 BaseByGuestAmt,omitempty" json:"BaseByGuestAmt,omitempty"`
	XMLName           xml.Name           `xml:"http://www.opentravel.org/OTA/2003/05 BaseByGuestAmts,omitempty" json:"BaseByGuestAmts,omitempty"`
}

// HRABaseByGuestAmt ...
type HRABaseByGuestAmt struct {
	AmountAfterTax string   `xml:" AmountAfterTax,attr"  json:",omitempty"`
	XMLName        xml.Name `xml:"http://www.opentravel.org/OTA/2003/05 BaseByGuestAmt,omitempty" json:"BaseByGuestAmt,omitempty"`
}
