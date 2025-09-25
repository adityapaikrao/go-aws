package main

import (
	ticket "basics/tickets" //alias for the imported package
	"fmt"
)

func main() {

	fmt.Println("hello world")

	newTicket := ticket.Ticket{
		ID:    2,
		Event: "Concert",
	}
	newTicket.PrintEvent()
}
