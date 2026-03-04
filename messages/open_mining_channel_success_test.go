package sv2

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestOpenStandardMiningChannelSuccess_Deserialize(t *testing.T) {
	// 1. Имитируем байты от Luxor (Payload: ReqID(4) + ConnID(4) + Spacing(4) + Target(32))
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, uint32(12345)) // RequestID
	binary.Write(buf, binary.LittleEndian, uint32(999))   // ConnectionID
	binary.Write(buf, binary.LittleEndian, uint32(10))    // TargetSpacing (10 сек)

	target := make([]byte, 32)
	target[0] = 0xff // Задаем тестовую сложность
	buf.Write(target)

	payload := buf.Bytes()

	// 2. Пробуем десериализовать
	msg := OpenStandardMiningChannelSuccess{}
	err := msg.Deserialize(payload)
	if err != nil {
		t.Fatalf("Bambaklaat! Deserialize failed: %v", err)
	}

	// 3. Сверяем результаты
	if msg.ConnectionID != 999 {
		t.Errorf("Expected ConnectionID 999, got %d", msg.ConnectionID)
	}

	if msg.TargetSpacing != 10 {
		t.Errorf("Expected Spacing 10, got %d", msg.TargetSpacing)
	}

	if msg.InitialTarget[0] != 0xff {
		t.Error("InitialTarget data mismatch")
	}
}
