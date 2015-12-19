package template

import (
	"fmt"
	"time"

	"github.com/docker/go-units"
)

type Metadata struct {
	Tag        string
	Repository string

	Created JSONTime
}

func (m Metadata) String() []string {
	tDelta := time.Now().Sub(time.Time(m.Created))
	return []string{m.Tag, m.Repository, units.HumanDuration(tDelta) + " ago"}
}

type JSONTime time.Time

const (
	timeFormat = "Mon Jan 2 15:04 -0700 MST 2006"
)

func NewTime() JSONTime {
	return JSONTime(time.Now())
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf(`"%s"`, time.Time(*t).Format(timeFormat))

	return []byte(stamp), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	time, err := time.Parse(timeFormat, string(b)[1:len(b)-1])
	if err != nil {
		return err
	}

	*t = JSONTime(time)

	return nil
}

func (t JSONTime) String() string {
	return fmt.Sprintf("%s", time.Time(t).Format(timeFormat))
}
