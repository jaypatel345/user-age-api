package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	dob := time.Date(
		2000,
		1,
		1,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	age := CalculateAge(dob)
	expected := time.Now().Year() - 2000

	if age != expected {
		t.Errorf(
			"expected %d got %d",
			expected,
			age,
		)
	}
}
