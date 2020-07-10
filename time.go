package mono

import (
	"strconv"
	"time"
)

// Time defines a timestamp encoded as epoch seconds in JSON
// Most of it's code taken from https://github.com/pieterclaerhout/example-json-unixtimestamp/blob/master/cmd/example-json-timestamp/time.go
type Time struct {
	time.Time
}

// MarshalJSON is used to convert the timestamp to JSON
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(t.Unix(), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *Time) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*t = Time{time.Unix(q, 0).UTC()}
	return nil
}
