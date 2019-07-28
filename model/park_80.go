package model

import "time"

type park80 struct{}

func NewPark80() Pricer {
	return park80{}
}

func (m park80) Price(date time.Time) uint {
	if isLate(date) {
		return 1100
	}
	return 1400
}
