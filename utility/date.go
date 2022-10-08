package utility

import (
	"time"
)

type Dater interface {
	getFrom() string
	getTo() string
}

type Dates struct {
	from string
	to   string
}

func (d Dates) GetFrom() string {
	return d.from
}

func (d Dates) GetTo() string {
	return d.to
}

// CreateDates Offset is number of days from timestamp
func CreateDates(offset int64, timestamp time.Time) *Dates {
	secondsOffset := offset * 86400
	newUnixTime := timestamp.Unix() + secondsOffset
	newTimestamp := time.Unix(newUnixTime, 0)

	return &Dates{
		from: timestamp.Format("2006-01-02"),
		to:   newTimestamp.Format("2006-01-02"),
	}
}
