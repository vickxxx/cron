package cron

import "time"

type TimeSchedule struct {
	Tm       time.Time
	Location *time.Location
}

func (s *TimeSchedule) Next(t time.Time) time.Time {

	if s.Tm.Before(t) {
		return time.Time{}
	}

	return s.Tm

}
