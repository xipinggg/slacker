package record

import (
	"time"
)

type Record struct {
	ID         string
	Type       string
	CreatorID  string
	BeginTime  time.Time
	UpdateTime *time.Time
	EndTime    *time.Time
}

func (r *Record) RecordTime() time.Duration {
	if r == nil {
		return 0
	}

	if r.EndTime == nil {
		return 0
	}

	return r.BeginTime.Sub(*r.EndTime)
}
