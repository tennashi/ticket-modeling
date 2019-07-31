package model

import "time"

type Ticket struct {
	p Pricer
}

func NewTicket(p Pricer) Ticket {
	return Ticket{p}
}

func (t Ticket) Price(date time.Time) uint {
	return t.p.Price(date)
}
