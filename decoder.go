package eventsource

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

//Decoder wraps and buffers an io.Reader
//and extracts Events from the stream
type Decoder struct {
	r *bufio.Reader
}

//NewDecoder returns a Decoder
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: bufio.NewReader(r)}
}

//Decode a single Event from the stream
func (d *Decoder) Decode(event *Event) error {
	if event == nil {
		return errors.New("event is nil")
	}
	for {
		line, err := d.r.ReadString('\n')
		if err != nil {
			return err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			break //blank line is the end of the event
		}
		i := strings.Index(line, ": ")
		if i == -1 || i == 0 || i == len(line)+2 {
			return errors.New("line has invalid field")
		}
		field := line[0:i]
		value := line[i+2:]
		switch field {
		case "id":
			event.ID = value
		case "event":
			event.Type = value
		case "data":
			event.Data = []byte(value)
		case "retry":
			event.Retry = value
		default:
			//unknown, spec says to ignore
		}
	}
	return nil
}
