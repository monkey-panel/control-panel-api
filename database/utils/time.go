package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time time.Time

const (
	// ISO860: YYYY-MM-DDThh:mm:ss.sssZ
	TimeFormat       = "2006-01-02T15:04:05.999Z"
	StringTimeFormat = `"` + TimeFormat + `"`
)

func (t *Time) UnmarshalJSON(data []byte) error {
	now, err := time.Parse(StringTimeFormat, string(data))
	*t = Time(now)
	return err
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t Time) Time() time.Time { return time.Time(t) }
func (t Time) String() string  { return t.Time().Format(TimeFormat) }

func Now() Time               { return New(time.Now()) }
func New(time time.Time) Time { return Time(time.UTC()) }

// for sql
func (t *Time) Scan(src any) error {
	if value, ok := src.(time.Time); ok {
		*t = Time(value.UTC())
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", src)
}

// for sql
func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := t.Time()
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}
