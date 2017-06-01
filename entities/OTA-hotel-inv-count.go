package entities

import "encoding/xml"

///////////////////////////////// Wrapper //////////////////////////////////////

// HICEnvelope ...
type HICEnvelope struct {
	SoapEnv   string     `xml:"xmlns SOAP-ENV,attr"  json:",omitempty"`
	Ns1       string     `xml:"xmlns ns1,attr"  json:",omitempty"`
	Ns2       string     `xml:"xmlns ns2,attr"  json:",omitempty"`
	Ns3       string     `xml:"xmlns ns3,attr"  json:",omitempty"`
	HICHeader *HICHeader `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header,omitempty" json:"Header,omitempty"`
	HICBody   *HICBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
	XMLName   xml.Name   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope,omitempty" json:"Envelope,omitempty"`
}

//////////////////////////////// Header part ///////////////////////////////////

// HICHeader ...
type HICHeader struct {
	HICFrom      *HICFrom     `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing From,omitempty" json:"From,omitempty"`
	HICMessageID string       `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing MessageID,omitempty" json:"MessageID,omitempty"`
	HICSecurity  *HICSecurity `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security,omitempty" json:"Security,omitempty"`
	XMLName      xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header,omitempty" json:"Header,omitempty"`
}

// HICSecurity ...
type HICSecurity struct {
	D3p1               string            `xml:"xmlns d3p1,attr"  json:",omitempty"`
	D3p1MustUnderstand string            `xml:"http://schemas.xmlsoap.org/soap/envelope/ mustUnderstand,attr"  json:",omitempty"`
	HICUsernameToken   *HICUsernameToken `xml:"UsernameToken,omitempty" json:"UsernameToken,omitempty"`
	XMLName            xml.Name          `xml:"http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd Security,omitempty" json:"Security,omitempty"`
}

// HICUsernameToken ...
type HICUsernameToken struct {
	HICPassword string   `xml:"Password,omitempty" json:"Password,omitempty"`
	HICUsername string   `xml:"Username,omitempty" json:"Username,omitempty"`
	XMLName     xml.Name `xml:"UsernameToken,omitempty" json:"UsernameToken,omitempty"`
}

// HICFrom ...
type HICFrom struct {
	Address string   `xml:"Address,omitempty" json:"Address,omitempty"`
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing From,omitempty" json:"Header,omitempty"`
}

// HICAddress ...
type HICAddress struct {
	Text    string   `xml:",chardata" json:",omitempty"`
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/ws/2004/08/addressing Address,omitempty" json:"Address,omitempty"`
}

//////////////////////////////////// Body part //////////////////////////////////////

// HICBody ...
type HICBody struct {
	HotelInvCountNotifRQ *HotelInvCountNotifRQ `xml:" OTA_HotelInvCountNotifRQ,omitempty" json:"OTA_HotelInvCountNotifRQ,omitempty"`
	XMLName              xml.Name              `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body,omitempty" json:"Body,omitempty"`
}

// HotelInvCountNotifRQ ...
type HotelInvCountNotifRQ struct {
	TimeStamp          string            `xml:" TimeStamp,attr"  json:",omitempty"`
	Version            string            `xml:" Version,attr"  json:",omitempty"`
	HICInventoriesList []*HICInventories `xml:" Inventories,omitempty" json:"Inventories,omitempty"`
}

// HICInventories ...
type HICInventories struct {
	HotelCode      string          `xml:" HotelCode,attr"  json:",omitempty"`
	HotelName      string          `xml:" HotelName,attr"  json:",omitempty"`
	HICInventories []*HICInventory `xml:" Inventory,omitempty" json:"Inventory,omitempty"`
}

// HICInventory ...
type HICInventory struct {
	HICInvCounts             *HICInvCounts                `xml:" InvCounts,omitempty" json:"InvCounts,omitempty"`
	StatusApplicationControl *HICStatusApplicationControl `xml:"http://www.opentravel.org/OTA/2003/05 StatusApplicationControl,omitempty" json:"StatusApplicationControl,omitempty"`
}

// HICStatusApplicationControl ...
type HICStatusApplicationControl struct {
	End         string   `xml:" End,attr"  json:",omitempty"`
	InvTypeCode string   `xml:" InvTypeCode,attr"  json:",omitempty"`
	Start       string   `xml:" Start,attr"  json:",omitempty"`
	XMLName     xml.Name `xml:"http://www.opentravel.org/OTA/2003/05 StatusApplicationControl,omitempty" json:"StatusApplicationControl,omitempty"`
}

// HICInvCounts ...
type HICInvCounts struct {
	HICInvCount []*HICInvCount `xml:" InvCount,omitempty" json:"InvCount,omitempty"`
}

// HICInvCount ...
type HICInvCount struct {
	Count     string `xml:" Count,attr"  json:",omitempty"`
	CountType string `xml:" CountType,attr"  json:",omitempty"`
}
