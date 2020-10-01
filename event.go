package eventsource

// Event describes a single Server-Send Event
type Event struct {
	Type  string
	ID    string
	Retry string
	Data  []byte
}
