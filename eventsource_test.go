package eventsource

import (
	"bytes"
	"testing"
)

func TestEncoderEncode(t *testing.T) {
	table := []struct {
		Event
		expected string
	}{
		{Event{Type: "type"}, "event: type\ndata\n\n"},
		{Event{ID: "123"}, "id: 123\ndata\n\n"},
		{Event{Retry: "10000"}, "retry: 10000\ndata\n\n"},
		{Event{Data: []byte("data")}, "data: data\n\n"},
		{Event{ResetID: true}, "id\ndata\n\n"},
	}
	for i, tt := range table {
		b := bytes.Buffer{}
		if err := WriteEvent(&b, tt.Event); err != nil {
			t.Fatalf("%d. write error: %q", i, err)
		}
		if b.String() != tt.expected {
			t.Errorf("%d. expected %q, got %q", i, tt.expected, b.String())
		}
	}
}
