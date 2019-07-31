package model

import "time"

type Handicapped struct {
	division StudentDivision
}

func NewHandicapped(division StudentDivision) Handicapped {
	return Handicapped{
		division: division,
	}
}

func (h Handicapped) Price(date time.Time) uint {
	switch h.division {
	case StudentNotStudent, StudentUniversity:
		return h.univAndOver(date)
	default:
		return h.other(date)
	}
}

func (h Handicapped) univAndOver(date time.Time) uint {
	return 1000
}

func (h Handicapped) other(date time.Time) uint {
	return 900
}
