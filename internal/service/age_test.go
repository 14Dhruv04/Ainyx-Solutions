package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {

	tests := []struct {
		name string
		dob  time.Time
	}{
		{
			name: "Born in 2000",
			dob: time.Date(
				2000,
				time.January,
				1,
				0, 0, 0, 0,
				time.UTC,
			),
		},
		{
			name: "Born in 1990",
			dob: time.Date(
				1990,
				time.May,
				10,
				0, 0, 0, 0,
				time.UTC,
			),
		},
	}

	for _, test := range tests {

		age := CalculateAge(test.dob)

		if age < 0 {
			t.Errorf(
				"%s: age cannot be negative",
				test.name,
			)
		}
	}
}