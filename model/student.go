package model

import "time"

type student struct {
	division StudentDivision
}

func NewStudent(division StudentDivision) Pricer {
	return student{
		division: division,
	}
}

func (s student) Price(date time.Time) uint {
	switch s.division {
	case StudentUniversity:
		return s.univ(date)
	case StudentElementaryOrLess, StudentJuniorHighSchool, StudentHighSchool:
		return s.other(date)
	default:
		return 0
	}
}

func (s student) univ(date time.Time) uint {
	if isCinemaDay(date) {
		return 1100
	}
	if isLate(date) {
		return 1300
	}
	return 1500
}

func (s student) other(date time.Time) uint {
	return 1000
}
