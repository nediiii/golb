package utils

import (
	"testing"
	"time"
)

func TestTime2UnixString(t *testing.T) {
	var testCase = []struct {
		inputTime time.Time
		expectStr string
	}{
		{time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC), "0"},
		{time.Date(2016, time.August, 15, 0, 0, 0, 0, time.UTC), "1471219200"},
		{time.Date(2017, time.February, 16, 0, 0, 0, 0, time.UTC), "1487203200"},
	}

	for _, tC := range testCase {
		unixTimeStr := Time2UnixString(&tC.inputTime)
		if unixTimeStr != tC.expectStr {
			t.Error("Time2UnixString error:", unixTimeStr)
		}
	}
}
