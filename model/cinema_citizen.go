package model

import "time"

type CinemaCitizen struct {
	age uint8
}

func NewCinemaCitizen(age uint8) CinemaCitizen {
	return CinemaCitizen{
		age: age,
	}
}

func (c CinemaCitizen) Price(date time.Time) uint {
	if c.age >= 60 {
		return c.cc60(date)
	}
	return c.cc(date)
}

func (c CinemaCitizen) cc60(date time.Time) uint {
	return 1000
}

func (c CinemaCitizen) cc(date time.Time) uint {
	if !isHoliday(date) && !isLate(date) {
		return 1300
	}
	return 1000
}
