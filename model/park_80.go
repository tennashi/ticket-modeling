package model

import "time"

type Park80 struct{}

func NewPark80() Park80 {
	return Park80{}
}

func (m Park80) Price(date time.Time) uint {
	if isLate(date) {
		return 1100
	}
	return 1400
}
