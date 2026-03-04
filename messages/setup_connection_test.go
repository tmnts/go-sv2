package sv2

import (
	"bytes"
	"testing"
)

func TestSetupConnection_Serialize(t *testing.T) {
	// 1. Создаем тестовое сообщение (как будто мы майнер)
	msg := SetupConnection{
		Protocol:     MiningProtocol,
		MinVersion:   2,
		MaxVersion:   2,
		Flags:        0,
		EndpointHost: "localhost",
		EndpointPort: 3333,
	}

	// 2. Запускаем сериализацию
	data, err := msg.Serialize()
	if err != nil {
		t.Fatalf("Failed to serialize: %v", err)
	}

	// 3. Проверяем заголовок (первые 6 байт)
	// Extension(2) + MsgType(1) + Length(3)
	if data[2] != MsgTypeSetupConnection {
		t.Errorf("Expected message type %d, got %d", MsgTypeSetupConnection, data[2])
	}

	// 4. Проверяем "магию" длины (uint24)
	// Длина данных после хедера: 1+2+2+4+1+9+2 = 21 байт (0x15)
	expectedLen := byte(21)
	if data[3] != expectedLen {
		t.Errorf("Expected payload length %d, got %d", expectedLen, data[3])
	}

	// 5. Проверяем, что хост записался корректно
	if !bytes.Contains(data, []byte("localhost")) {
		t.Error("Serialized data does not contain endpoint host")
	}
}

func TestSetupConnection_Validation(t *testing.T) {
	// Тест на слишком длинный хост (> 255 символов)
	longHost := ""
	for i := 0; i < 256; i++ {
		longHost += "a"
	}

	msg := SetupConnection{EndpointHost: longHost}
	_, err := msg.Serialize()

	if err == nil {
		t.Error("Expected error for long hostname, but got nil")
	}
}
