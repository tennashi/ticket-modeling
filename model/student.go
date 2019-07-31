package model

import "time"

type Student struct {
	division StudentDivision
}

func NewStudent(division StudentDivision) Student {
	return Student{
		division: division,
	}
}

func (s Student) Price(date time.Time) uint {
	switch s.division {
	case StudentUniversity:
		return s.univ(date)
	case StudentElementaryOrLess, StudentJuniorHighSchool, StudentHighSchool:
		return s.other(date)
	default:
		return 0
	}
}

func (s Student) univ(date time.Time) uint {
	if isCinemaDay(date) {
		return 1100
	}
	if isLate(date) {
		return 1300
	}
	return 1500
}

func (s Student) other(date time.Time) uint {
	return 1000
}
