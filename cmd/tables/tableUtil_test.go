package table

import (
	"testing"
	"time"
)

type mockTimeable struct {
	unixTime int
	delay    int
}

func (m mockTimeable) GetUnixDepartureTime() int {
	return m.unixTime
}

func (m mockTimeable) GetDelayInSeconds() int {
	return m.delay
}

// Test for CalculateHumanRelativeTime function
func TestCalculateHumanRelativeTime(t *testing.T) {
	now := time.Now()

	tests := []struct {
		unixTime int
		delay    int
		expected string
	}{
		// Event is happening now (no delay)
		{int(now.Unix()), 0, "now"},
		// Event in 5 minutes (no delay)
		{int(now.Add(5 * time.Minute).Unix()), 0, "4 min"},
		// Event in 1 hour (no delay)
		{int(now.Add(1 * time.Hour).Unix()), 0, "59 min"},
		// Event in 1 hour and 30 minutes (no delay)
		{int(now.Add(1*time.Hour + 30*time.Minute).Unix()), 0, "1 hour 29 min"},
		// Event in 2 hours (no delay)
		{int(now.Add(2 * time.Hour).Unix()), 0, "1 hour 59 min"},
		// Event is 30 minutes ago but with 1-hour delay
		{int(now.Add(-30 * time.Minute).Unix()), 3600, "29 min"},
		// Event in 45 minutes (30 minutes delay)
		{int(now.Add(15 * time.Minute).Unix()), 1800, "44 min"},
	}

	for _, test := range tests {
		mock := mockTimeable{unixTime: test.unixTime, delay: test.delay}
		result := CalculateHumanRelativeTime(mock)
		if result != test.expected {
			t.Errorf("CalculateHumanRelativeTime(%v) = %s; want %s", mock, result, test.expected)
		}
	}
}
