package model

import "time"

type Pricer interface {
	Price(date time.Time) uint
}
