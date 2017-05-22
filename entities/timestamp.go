package entities

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Timestamp time.Time

var formats = []string{
	"2006-01-02",
	"2006-01-02 15:04",
	"2006-01-02 15:04:05Z07",
	"2006-01-02 15:04:05Z07:00",
	"2006-01-02T15:04",
	"2006-01-02T15:04:04",
	"2006-01-02T15:04:04:05Z07",
	"2006-01-02T15:04:04:05Z07:00",
	time.RFC3339,     //"2006-01-02T15:04:05Z07:00"
	time.RFC3339Nano, // "2006-01-02T15:04:05.999999999Z07:00"
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	tm := time.Time(*t)
	return tm.MarshalJSON()
}
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts := string(b)
	ts = strings.Replace(ts, "\"", "", -1)
	var tm time.Time
	var err error
	for index := range formats {
		tm, err = time.Parse(formats[index], ts)
		if err == nil {
			*t = Timestamp(tm)
			s := (*t).String()
			_ = s
			return nil
		}
	}
	return err
}
func (t *Timestamp) GetBSON() (interface{}, error) {
	if time.Time(*t).IsZero() {
		return nil, nil
	}

	return time.Time(*t), nil
}
func (t *Timestamp) SetBSON(raw bson.Raw) error {
	var tm time.Time

	if err := raw.Unmarshal(&tm); err != nil {
		return err
	}

	*t = Timestamp(tm)

	return nil
}

func (t *Timestamp) String() string {
	return time.Time(*t).String()
}
