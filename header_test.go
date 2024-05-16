package guic

import (
	"fmt"
	"testing"
)

func TestLongHeaderBytes(t *testing.T) {
	tests := []struct {
		form                          bool
		versionSpecific               uint8
		version                       uint32
		destinationConnectionIDLength uint8
		destinationConnectionID       []byte
		sourceConnectionIDLength      uint8
		sourceConnectionID            []byte
		versionSpecificData           []byte
		expected                      []byte
	}{
		{
			form:                          true,
			versionSpecific:               0x01,
			version:                       0x01,
			destinationConnectionIDLength: 4,
			destinationConnectionID:       []byte{0xde, 0xad, 0xbe, 0xef},
			sourceConnectionIDLength:      4,
			sourceConnectionID:            []byte{0xba, 0xad, 0xf0, 0x0d},
			versionSpecificData:           []byte{0xab, 0xcd},
			expected: []byte{
				0x80,
				0x00,
				0x00,
				0x00,
				0x01,
				0x04,
				0xde,
				0xad,
				0xbe,
				0xef,
				0x04,
				0xba,
				0xad,
				0xf0,
				0x0d,
				0xab,
				0xcd,
			},
		},
	}

	for _, tt := range tests {
		l := newlongHeader(tt.form, tt.versionSpecific, tt.version, tt.destinationConnectionIDLength, tt.destinationConnectionID, tt.sourceConnectionIDLength, tt.sourceConnectionID, tt.versionSpecificData)
		b := l.Bytes()

		if len(b) != len(tt.expected) {
			t.Errorf("expected %d bytes, got %d", len(tt.expected), len(b))
			return
		}

		for i := 0; i < len(b); i++ {
			if b[i] != tt.expected[i] {
				t.Errorf("expected %x, got %x", tt.expected[i], b[i])
			}
		}
	}
}

func TestShortHeaderBytes(t *testing.T) {
	tests := []struct {
		form                    bool
		versionSpecific         uint8
		DestinationConnectionID []byte
		VersionSpecificData     []byte
		expected                []byte
	}{
		{
			form:                    false,
			versionSpecific:         0x01,
			DestinationConnectionID: []byte{0xde, 0xad, 0xbe, 0xef},
			VersionSpecificData:     []byte{0xab, 0xcd},
			expected: []byte{
				0x01,
				0xde,
				0xad,
				0xbe,
				0xef,
				0xab,
				0xcd,
			},
		},
	}

	for _, tt := range tests {
		s := newshortHeader(tt.form, tt.versionSpecific, tt.DestinationConnectionID, tt.VersionSpecificData)
		b := s.Bytes()
		fmt.Printf("%v\n", tt.expected)
		fmt.Printf("%v\n", b)

		if len(b) != len(tt.expected) {
			t.Errorf("expected %d bytes, got %d", len(tt.expected), len(b))
			return
		}

		for i := 0; i < len(b); i++ {
			if b[i] != tt.expected[i] {
				t.Errorf("expected %x, got %x", tt.expected[i], b[i])
			}
		}
	}
}
