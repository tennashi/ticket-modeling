package model

import "time"

func isHoliday(date time.Time) bool {
	wd := date.Weekday()
	if wd == time.Sunday || wd == time.Saturday {
		return true
	}
	return false
}

func isLate(date time.Time) bool {
	if date.Hour() > 20 {
		return true
	}
	return false
}

func isCinemaDay(date time.Time) bool {
	if date.Day() == 1 {
		return true
	}
	return false
}
