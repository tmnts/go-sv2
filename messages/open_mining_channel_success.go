package sv2

import (
	"bytes"
	"encoding/binary"
)

const MsgTypeOpenStandardMiningChannelSuccess uint8 = 0x02

type OpenStandardMiningChannelSuccess struct {
	RequestID     uint32
	ConnectionID  uint32 // Твой уникальный ID в пуле
	TargetSpacing uint32 // Как часто пул хочет получать решения (в секундах)
	InitialTarget []byte // Начальная сложность (32 байта)
}

func (m *OpenStandardMiningChannelSuccess) Deserialize(payload []byte) error {
	buf := bytes.NewReader(payload)

	// Читаем ID запроса, на который это ответ
	binary.Read(buf, binary.LittleEndian, &m.RequestID)
	// Получаем наш ConnectionID
	binary.Read(buf, binary.LittleEndian, &m.ConnectionID)
	// Получаем интервал таргета
	binary.Read(buf, binary.LittleEndian, &m.TargetSpacing)

	// Читаем 32 байта начальной сложности
	m.InitialTarget = make([]byte, 32)
	_, err := buf.Read(m.InitialTarget)
	return err
}
