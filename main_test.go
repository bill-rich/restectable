package main

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"encoding/hex"
	"encoding/json"
)

func TestCreateMessage(t *testing.T) {
	messageContent := "somemessage123"
	hashHex := sha256.Sum256([]byte(messageContent))
	hash := hex.EncodeToString(hashHex[:])

	router := setupRouter()
	w := httptest.NewRecorder()

	// Create the message on the server.
	var messageResp Message

	req := createTestMessageRequest(messageContent)

	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &messageResp); err != nil {
		t.Error(fmt.Errorf("unable to unmarshal json: %s, content: %s", err, string(w.Body.Bytes())))
	}

	// Check that the returned hash matches the expected hash.
	if messageResp.Hash != hash {
		t.Error(fmt.Errorf("incorrect sha256sum. expected %s, got %s", hash, messageResp.Hash))
	}
}

func TestRetrieveMessage(t *testing.T) {
	messageContent := "somemessage123"

	router := setupRouter()

	// First create the message on the server.
	req := createTestMessageRequest(messageContent)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var message Message
	var respMessage Message

	// Unmarshal the response to get the sha256sum.
	if err := json.Unmarshal(w.Body.Bytes(), &message); err != nil {
		t.Error(fmt.Errorf("unable to unmarshal json: %s, content: %s", err, string(w.Body.Bytes())))
	}

	// Create the new request to retrieve the message using the sha256sum.
	req = getTestMessageRequest(message.Hash)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if err := json.Unmarshal(w.Body.Bytes(), &respMessage); err != nil {
		t.Error(fmt.Errorf("unable to unmarshal json: %s, content: %s", err, string(w.Body.Bytes())))
	}

	// Check that the retrieved message matches the one that was sent.
	if respMessage.Content != messageContent {
		t.Error(fmt.Errorf("incorrect message content. expected %s, got %s", messageContent, respMessage.Content))
	}
}

func createTestMessageRequest(messageContent string) *http.Request {
	var message Message
	message = Message{
		Content: messageContent,
	}
	jsonData, _ := json.Marshal(message)
	req, _ := http.NewRequest("POST", "/message", strings.NewReader(string(jsonData)))
	req.Header.Set("Content-Type", "application/json")
	return req
}

func getTestMessageRequest(hash string) *http.Request {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/message/%s", hash), nil)
	return req
}
