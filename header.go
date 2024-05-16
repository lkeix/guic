package guic

import (
	"bytes"
	"encoding/binary"
)

type longHeader struct {
	Form                          bool
	VersionSpecific               uint8
	Version                       uint32
	DestinationConnectionIDLength uint8
	DestinationConnectionID       []byte
	SourceConnectionIDLength      uint8
	SourceConnectionID            []byte
	VersionSpecificData           []byte
}

func newlongHeader(form bool, versionSpecific uint8, version uint32, destinationConnectionIDLength uint8,
	destinationConnectionID []byte, sourceConnectionIDLength uint8, sourceConnectionID []byte, versionSpecificData []byte) *longHeader {
	return &longHeader{
		Form:                          form,
		VersionSpecific:               versionSpecific,
		Version:                       version,
		DestinationConnectionIDLength: destinationConnectionIDLength,
		DestinationConnectionID:       destinationConnectionID,
		SourceConnectionIDLength:      sourceConnectionIDLength,
		SourceConnectionID:            sourceConnectionID,
		VersionSpecificData:           versionSpecificData,
	}
}

func (l *longHeader) Bytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 1+4+1+int(l.DestinationConnectionIDLength)+1+int(l.SourceConnectionIDLength)+len(l.VersionSpecificData)))

	var firstByte byte
	if l.Form {
		firstByte |= 0x7F
	}
	firstByte = firstByte + l.VersionSpecific
	buf.WriteByte(firstByte)

	binary.Write(buf, binary.BigEndian, l.Version)

	buf.WriteByte(l.DestinationConnectionIDLength)
	buf.Write(l.DestinationConnectionID)
	buf.WriteByte(l.SourceConnectionIDLength)
	buf.Write(l.SourceConnectionID)
	buf.Write(l.VersionSpecificData)

	return buf.Bytes()
}

type shortHeader struct {
	Form                    bool
	VersionSpecific         uint8
	DestinationConnectionID []byte
	VersionSpecificData     []byte
}

func newshortHeader(form bool, versionSpecific uint8, destinationConnectionID []byte, versionSpecificData []byte) *shortHeader {
	return &shortHeader{
		Form:                    form,
		VersionSpecific:         versionSpecific,
		DestinationConnectionID: destinationConnectionID,
		VersionSpecificData:     versionSpecificData,
	}
}

func (s *shortHeader) Bytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 1+len(s.DestinationConnectionID)+len(s.VersionSpecificData)))

	var firstByte byte
	if !s.Form {
		firstByte |= 0x00
	}
	firstByte |= s.VersionSpecific & 0x7F
	buf.WriteByte(firstByte)
	buf.Write(s.DestinationConnectionID)
	buf.Write(s.VersionSpecificData)

	return buf.Bytes()
}

type negotiationHeader struct {
	Form                          bool
	Unused                        uint8
	Version                       uint32
	DestinationConnectionIDLength uint8
	DestinationConnectionID       []byte
	SourceConnectionIDLength      uint8
	SourceConnectioniD            []byte
	VersionSpecificData           []byte
}

func newnegotiationHeader(form bool, unused uint8, version uint32, destinationConnectionIDLength uint8,
	destinationConnectionID []byte, sourceConnectionIDLength uint8, sourceConnectionID []byte, versionSpecificData []byte) *negotiationHeader {
	return &negotiationHeader{
		Form:                          form,
		Unused:                        unused,
		Version:                       version,
		DestinationConnectionIDLength: destinationConnectionIDLength,
		DestinationConnectionID:       destinationConnectionID,
		SourceConnectionIDLength:      sourceConnectionIDLength,
		SourceConnectioniD:            sourceConnectionID,
		VersionSpecificData:           versionSpecificData,
	}
}

func (n *negotiationHeader) Bytes() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 1+4+1+int(n.DestinationConnectionIDLength)+int(n.SourceConnectionIDLength)+len(n.VersionSpecificData)))

	var firstByte byte
	if n.Form {
		firstByte |= 0x7F
	}
	firstByte = firstByte + n.Unused

	buf.WriteByte(firstByte)
	binary.Write(buf, binary.BigEndian, n.Version)

	buf.WriteByte(n.DestinationConnectionIDLength)
	buf.Write(n.DestinationConnectionID)
	buf.WriteByte(n.SourceConnectionIDLength)
	buf.Write(n.SourceConnectioniD)
	buf.Write(n.VersionSpecificData)

	return buf.Bytes()
}
