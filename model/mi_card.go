package model

import "time"

type MICard struct{}

func NewMICard() MICard {
	return MICard{}
}

func (m MICard) Price(date time.Time) uint {
	if isLate(date) {
		return 1300
	}
	return 1600
}
