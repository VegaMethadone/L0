package client

import "testing"

func TestClient(t *testing.T) {
	message := Message{
		ID:      1,
		Title:   "Заголовок сообщения",
		Content: "Содержание сообщения",
	}
	Client(message)
}
