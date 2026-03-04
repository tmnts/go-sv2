package sv2

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Константы из спецификации
const (
	MiningProtocol         uint8 = 0
	JobNegotiation         uint8 = 1
	MsgTypeSetupConnection uint8 = 0x00
	HeaderLength           int   = 6
)

// SetupConnection — первое сообщение от клиента к серверу
type SetupConnection struct {
	Protocol     uint8
	MinVersion   uint16
	MaxVersion   uint16
	Flags        uint32
	EndpointHost string
	EndpointPort uint16
}

// Header — те самые 6 байт перед каждым сообщением
type Header struct {
	ExtensionType uint16
	MsgType       uint8
	MsgLength     uint32 // Используем uint32 для хранения 24-битного значения
}

func (s *SetupConnection) Serialize() ([]byte, error) {
	if len(s.EndpointHost) > 255 {
		return nil, fmt.Errorf("hostname too long: %d (max 255)", len(s.EndpointHost))
	}

	buf := new(bytes.Buffer)

	// 1. Payload
	binary.Write(buf, binary.LittleEndian, s.Protocol)
	binary.Write(buf, binary.LittleEndian, s.MinVersion)
	binary.Write(buf, binary.LittleEndian, s.MaxVersion)
	binary.Write(buf, binary.LittleEndian, s.Flags)

	hostLen := uint8(len(s.EndpointHost))
	binary.Write(buf, binary.LittleEndian, hostLen)
	buf.WriteString(s.EndpointHost)

	binary.Write(buf, binary.LittleEndian, s.EndpointPort)

	payload := buf.Bytes()
	payloadLen := uint32(len(payload))

	// 2. Header
	headerBuf := new(bytes.Buffer)
	binary.Write(headerBuf, binary.LittleEndian, uint16(0)) // Extension
	binary.Write(headerBuf, binary.LittleEndian, MsgTypeSetupConnection)

	// 24-bit length
	lenBytes := []byte{
		byte(payloadLen),
		byte(payloadLen >> 8),
		byte(payloadLen >> 16),
	}
	headerBuf.Write(lenBytes)

	return append(headerBuf.Bytes(), payload...), nil
}

func DeserializeHeader(data []byte) (Header, error) {
	if len(data) < HeaderLength {
		return Header{}, fmt.Errorf("data too short: %d", len(data))
	}

	// Собираем 24 бита длины
	length := uint32(data[3]) | uint32(data[4])<<8 | uint32(data[5])<<16

	return Header{
		ExtensionType: binary.LittleEndian.Uint16(data[0:2]),
		MsgType:       data[2],
		MsgLength:     length,
	}, nil
}

func (s *SetupConnection) Deserialize(payload []byte) error {
	buf := bytes.NewReader(payload)

	if err := binary.Read(buf, binary.LittleEndian, &s.Protocol); err != nil {
		return err
	}
	binary.Read(buf, binary.LittleEndian, &s.MinVersion)
	binary.Read(buf, binary.LittleEndian, &s.MaxVersion)
	binary.Read(buf, binary.LittleEndian, &s.Flags)

	var hostLen uint8
	if err := binary.Read(buf, binary.LittleEndian, &hostLen); err != nil {
		return err
	}

	hostBytes := make([]byte, hostLen)
	if _, err := buf.Read(hostBytes); err != nil {
		return err
	}
	s.EndpointHost = string(hostBytes)

	return binary.Read(buf, binary.LittleEndian, &s.EndpointPort)
}
