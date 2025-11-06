package enum

type TicketStatus string

const (
	TicketStatusOpen       TicketStatus = "open"
	TicketStatusInProgress TicketStatus = "in_progress"
	TicketStatusResolved   TicketStatus = "resolved"
	TicketStatusClosed     TicketStatus = "closed"
)

func (s TicketStatus) IsValid() bool {
	switch s {
	case TicketStatusOpen, TicketStatusInProgress, TicketStatusResolved, TicketStatusClosed:
		return true
	default:
		return false
	}
}