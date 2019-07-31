package main

import (
	"fmt"
	"time"

	"github.com/tennashi/ticket-modeling/model"
)

func main() {
	ticket := model.NewTicket(model.NewCinemaCitizen(30))
	fmt.Println(ticket.Price(time.Now()))
}
