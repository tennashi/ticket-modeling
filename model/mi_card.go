package model

import "time"

type miCard struct{}

func NewMICard() Pricer {
	return miCard{}
}

func (m miCard) Price(date time.Time) uint {
	if isLate(date) {
		return 1300
	}
	return 1600
}
