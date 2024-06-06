package test

import (
	"testing"
	"time"
)

type TimeChecker struct {
	Now func() time.Time
}

func (tc *TimeChecker) HasDurationPassed(startTime time.Time, duration time.Duration) bool {
	return tc.Now().After(startTime.Add(duration))
}

func TestHasDurationPassed(t *testing.T) {
	fixedTime := time.Date(2023, time.June, 5, 12, 0, 0, 0, time.UTC)

	tc := &TimeChecker{
		Now: func() time.Time {
			return fixedTime
		},
	}

	startTime := fixedTime.Add(-10 * time.Minute)
	duration := 5 * time.Minute
	if !tc.HasDurationPassed(startTime, duration) {
		t.Errorf("expected true, got false")
	}

	startTime = fixedTime.Add(-3 * time.Minute)
	duration = 5 * time.Minute
	if tc.HasDurationPassed(startTime, duration) {
		t.Errorf("expected false, got true")
	}
}
