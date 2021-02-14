// Copyright 2020 Matt Ho
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package epochdate

import (
	"time"
)

func isLeapYear(year int) bool {
	switch {
	case year%4 > 0:
		return false
	case year%100 > 0:
		return true
	case year%400 > 0:
		return false
	default:
		return true
	}
}

// FromSeconds converts epochSecs to an epoch date assuming UTC
func FromSeconds(epochSecs int64) int {
	return From(time.Unix(epochSecs, 0).In(time.UTC))
}

// From converts epochSecs to an epoch date using the location associated with the time provided
func From(t time.Time) int {
	var n int
	for year := t.Year() - 1; year >= 1970; year-- {
		n += 365
		if isLeapYear(year) {
			n++
		}
	}
	return n + t.YearDay() - 1
}

// Time converts an epoch date a UTC time
func Time(epochDate int) time.Time {
	return TimeInLocation(epochDate, time.UTC)
}

// Time converts the epoch date to time at midnight on the location provided
func TimeInLocation(epochDate int, loc *time.Location) time.Time {
	year := 1970
	for {
		days := 365
		if isLeapYear(year) {
			days++
		}
		if epochDate < days {
			break
		}
		epochDate -= days
		year++
	}

	date := time.Date(year, time.January, 1, 0, 0, 0, 0, loc)
	return date.AddDate(0, 0, epochDate)
}
