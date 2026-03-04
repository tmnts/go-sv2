package sv2

import (
	"bytes"
	"encoding/binary"
)

const MsgTypeOpenStandardMiningChannel uint8 = 0x01

type OpenStandardMiningChannel struct {
	RequestID       uint32
	UserIdentity    string  // Логин майнера на пуле (например, "lera.worker1")
	NominalHashRate float32 // Ожидаемый хешрейт майнера
	MaxTarget       []byte  // Сложность (32 байта, BigInt)
}

func (m *OpenStandardMiningChannel) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)

	// 1. Payload
	binary.Write(buf, binary.LittleEndian, m.RequestID)

	// UserIdentity: длина (uint8) + строка
	binary.Write(buf, binary.LittleEndian, uint8(len(m.UserIdentity)))
	buf.WriteString(m.UserIdentity)

	// Хешрейт (4 байта float32)
	binary.Write(buf, binary.LittleEndian, m.NominalHashRate)

	// MaxTarget: в Sv2 это обычно фиксированные 32 байта
	// Если передали меньше — добьем нулями
	target := make([]byte, 32)
	copy(target, m.MaxTarget)
	buf.Write(target)

	payload := buf.Bytes()
	payloadLen := uint32(len(payload))

	// 2. Header (6 байт)
	headerBuf := new(bytes.Buffer)
	binary.Write(headerBuf, binary.LittleEndian, uint16(0)) // Extension
	binary.Write(headerBuf, binary.LittleEndian, MsgTypeOpenStandardMiningChannel)

	// Наш любимый uint24 длины
	lenBytes := []byte{
		byte(payloadLen),
		byte(payloadLen >> 8),
		byte(payloadLen >> 16),
	}
	headerBuf.Write(lenBytes)

	return append(headerBuf.Bytes(), payload...), nil
}
