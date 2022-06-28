package protocol_test

import (
	"reflect"
	"testing"

	"github.com/vulpemventures/neutrino-elements/pkg/protocol"
)

func TestNewVerackMsg(t *testing.T) {
	tests := []struct {
		name     string
		input    protocol.Magic
		err      error
		expected *protocol.Message
	}{
		{name: "ok",
			input: protocol.MagicRegtest,
			err:   nil,
			expected: &protocol.Message{
				MessageHeader: protocol.MessageHeader{
					Magic:    [4]byte{0x12, 0x34, 0x56, 0x78},
					Command:  [12]byte{118, 101, 114, 97, 99, 107, 0, 0, 0, 0, 0, 0},
					Length:   uint32(0),
					Checksum: [4]byte{93, 246, 224, 226},
				},
				Payload: []byte{},
			}},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			actual, err := protocol.NewVerackMsg(test.input)
			if err != nil && test.err == nil {
				t.Errorf("unexpected error: %+v", err)
			}

			if err == nil && test.err != nil {
				t.Errorf("expected error: %+v, got: %+v", test.err, err)
			}

			if err != nil && test.err != nil && err.Error() != test.err.Error() {
				t.Errorf("expected error: %+v, got: %+v", err, test.err)
			}

			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("expected: %+v, actual: %+v", test.expected, actual)
			}
		})
	}
}
