package sv2

import (
	"bytes"
	"testing"
)

func TestOpenStandardMiningChannel_Serialize(t *testing.T) {
	// 1. Готовим данные (как будто подключаем твой воркер)
	msg := OpenStandardMiningChannel{
		RequestID:       12345,
		UserIdentity:    "lera.worker1",
		NominalHashRate: 150.5,            // Тот самый флоат!
		MaxTarget:       make([]byte, 32), // Пустой таргет для начала
	}

	// 2. Сериализуем
	data, err := msg.Serialize()
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// 3. Проверяем тип сообщения (должен быть 0x01)
	if data[2] != MsgTypeOpenStandardMiningChannel {
		t.Errorf("Expected msg type 0x01, got %d", data[2])
	}

	// 4. Проверяем, что наше имя воркера попало в байты
	if !bytes.Contains(data, []byte("lera.worker1")) {
		t.Error("Worker identity not found in serialized data")
	}
}
