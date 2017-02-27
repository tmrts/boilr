package template

import (
	"fmt"
	"time"

	"github.com/docker/go-units"
)

// Metadata contains the information for a template.
type Metadata struct {
	Tag        string
	Repository string

	Created JSONTime
}

// String returns the string slice form of Metadata.
func (m Metadata) String() []string {
	tDelta := time.Now().Sub(time.Time(m.Created))
	return []string{m.Tag, m.Repository, units.HumanDuration(tDelta) + " ago"}
}

// JSONTime is time.Time with JSON marshaling and unmarshaling implementations.
type JSONTime time.Time

const (
	// "Mon, 02 Jan 2006 15:04:05 -0700"
	timeFormat = time.RFC1123Z
)

// NewTime returns a new JSONTime containing the current time.
func NewTime() JSONTime {
	return JSONTime(time.Now())
}

// MarshalJSON marshals JSONTime to JSON.
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`"%s"`, time.Time(*t).Format(timeFormat))

	return []byte(stamp), nil
}

// UnmarshalJSON unmarshals JSON to JSONTime.
func (t *JSONTime) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(timeFormat, string(b)[1:len(b)-1])
	if err != nil {
		return err
	}

	*t = JSONTime(time)

	return nil
}

// String returns the string form of JSONTime.
func (t JSONTime) String() string {
	return fmt.Sprintf("%s", time.Time(t).Format(timeFormat))
}
