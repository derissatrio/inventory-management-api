package enum

type TicketSeverity string

const (
	SeverityLow      TicketSeverity = "low"
	SeverityMedium   TicketSeverity = "medium"
	SeverityHigh     TicketSeverity = "high"
	SeverityCritical TicketSeverity = "critical"
)

func (s TicketSeverity) IsValid() bool {
	switch s {
	case SeverityLow, SeverityMedium, SeverityHigh, SeverityCritical:
		return true
	default:
		return false
	}
}