package model

import "time"

type handicapped struct {
	division StudentDivision
}

func NewHandicapped(division StudentDivision) Pricer {
	return handicapped{
		division: division,
	}
}

func (h handicapped) Price(date time.Time) uint {
	switch h.division {
	case StudentNotStudent, StudentUniversity:
		return h.univAndOver(date)
	default:
		return h.other(date)
	}
}

func (h handicapped) univAndOver(date time.Time) uint {
	return 1000
}

func (h handicapped) other(date time.Time) uint {
	return 900
}
