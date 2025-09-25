package tickets

import "fmt"

type Ticket struct {
	ID    int
	Event string // public property
}

func (ticket Ticket) PrintEvent() { // public method
	fmt.Println(ticket.Event)
}

func (ticket Ticket) printID() {
	fmt.Println(ticket.ID)
}
