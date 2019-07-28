package model

import "time"

type cinemaCitizen struct {
	age uint8
}

func NewCinemaCitizen(age uint8) Pricer {
	return cinemaCitizen{
		age: age,
	}
}

func (c cinemaCitizen) Price(date time.Time) uint {
	if c.age >= 60 {
		return c.cc60(date)
	}
	return c.cc(date)
}

func (c cinemaCitizen) cc60(date time.Time) uint {
	return 1000
}

func (c cinemaCitizen) cc(date time.Time) uint {
	if !isHoliday(date) && !isLate(date) {
		return 1300
	}
	return 1000
}
