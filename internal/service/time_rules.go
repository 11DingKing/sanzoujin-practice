package service

import (
	"github.com/11DingKing/sanzoujin-practice/internal/domain"
	"time"
)

type TimeRule struct {
	Location *time.Location
	Grace    time.Duration
}

func NewTimeRule(name string, grace time.Duration) TimeRule {
	loc, err := time.LoadLocation(name)
	if err != nil {
		loc = time.UTC
	}
	return TimeRule{Location: loc, Grace: grace}
}
func (r TimeRule) Normalize(t time.Time) time.Time { return t.In(r.Location) }
func (r TimeRule) CanCheckIn(now, start time.Time) bool {
	return !r.Normalize(now).Before(r.Normalize(start).Add(-r.Grace))
}
func (r TimeRule) CanCheckOut(now, start, end time.Time) bool {
	return !r.Normalize(now).Before(r.Normalize(start)) && r.Normalize(now).Before(r.Normalize(end).Add(r.Grace))
}
func AttendanceStateFor(now, start time.Time, grace time.Duration) domain.AttendanceState {
	if now.Before(start) {
		return domain.AttendanceAbsent
	}
	if now.Sub(start) <= grace {
		return domain.AttendancePresent
	}
	return domain.AttendanceLate
}
