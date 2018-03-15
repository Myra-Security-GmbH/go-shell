package types

import (
	"encoding/json"
	"strings"
	"time"
)

//
// DateTime ...
//
type DateTime struct {
	time.Time
}

//
// DateTimeNow ...
//
func DateTimeNow() *DateTime {
	ret := &DateTime{}
	ret.Time = time.Now()

	return ret
}

//
// Truncate reappends the time.Truncate method to DateTime
//
func (dt *DateTime) Truncate(d time.Duration) *DateTime {
	return &DateTime{
		Time: dt.Time.Truncate(d),
	}
}

//
// Add reappends the time.Add method to DateTime
//
func (dt *DateTime) Add(duration time.Duration) *DateTime {
	return &DateTime{
		Time: dt.Time.Add(duration),
	}
}

//
// MarshalJSON ...
//
func (dt *DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.Format(time.RFC3339))
}

//
// UnmarshalJSON ...
//
func (dt *DateTime) UnmarshalJSON(b []byte) error {
	date := strings.Trim(string(b), "\"")

	t, err := time.Parse(
		"2006-01-02T15:04:05Z0700",
		date,
	)

	dt.Time = t

	return err
}

//
// ToUnixDate ...
//
func (dt *DateTime) ToUnixDate() string {
	format := "_2. Jan 2006 "

	if dt.Year() == time.Now().Year() {
		format = "_2. Jan 15:04"
	}

	return dt.Format(format)
}
