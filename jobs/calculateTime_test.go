package jobs

import (
	"fmt"
	"testing"
	"time"
)

func TestCalculateNextTuesdayAtNine(t *testing.T) {

	testCase := []struct {
		name     string
		input    time.Time
		expected time.Duration
	}{
		{
			name:     "First week of December",
			input:    time.Date(2023, time.December, 5, 8, 0, 0, 0, time.UTC),
			expected: 1 * time.Hour,
		},
		{
			name:     "Christmas Day",
			input:    time.Date(2023, time.December, 25, 8, 0, 0, 0, time.UTC),
			expected: 24*time.Hour + 1*time.Hour,
		},
	}

	for _, tc := range testCase {
		t.Run(tc.name, func(t *testing.T) {
			now = tc.input

			result := CalculateNextTuesdayAtNine()

			fmt.Println(result)

			if result != tc.expected {
				t.Errorf("CalculateNextTuesdayAtNine() for %s = %v, want %v", tc.name, result, tc.expected)
			}
		})
	}

}
