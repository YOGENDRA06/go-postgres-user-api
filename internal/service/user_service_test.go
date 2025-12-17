package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "birthday already happened this year",
			dob:      time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 1990,
		},
		{
			name:     "birthday not yet happened this year",
			dob:      time.Date(1990, time.December, 31, 0, 0, 0, 0, time.UTC),
			expected: time.Now().Year() - 1990 - 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := calculateAge(tt.dob)
			if age != tt.expected {
				t.Errorf("expected age %d, got %d", tt.expected, age)
			}
		})
	}
}
