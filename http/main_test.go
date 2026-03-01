package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*
We test the handler function directly.
No need to start a real server.
*/

func TestEchoHandler_Success(t *testing.T) {

	// Create request body
	body := map[string]string{
		"message": "hello",
	}

	jsonBody, _ := json.Marshal(body)

	// Create fake HTTP request
	req := httptest.NewRequest(
		http.MethodPost,
		"/echo",
		bytes.NewReader(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	// Create response recorder (captures output)
	rr := httptest.NewRecorder()

	// Call handler
	echoHandler(rr, req)

	// Check status code
	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}

	// Parse response JSON
	var resp EchoResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("invalid json response: %v", err)
	}

	if resp.Echo != "hello" {
		t.Fatalf("expected echo 'hello', got '%s'", resp.Echo)
	}
}

func TestEchoHandler_InvalidMethod(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodGet, // wrong method
		"/echo",
		nil,
	)

	rr := httptest.NewRecorder()

	echoHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rr.Code)
	}
}

func TestEchoHandler_InvalidJSON(t *testing.T) {

	req := httptest.NewRequest(
		http.MethodPost,
		"/echo",
		bytes.NewBufferString(`{invalid}`),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	echoHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}

func TestEchoHandler_EmptyMessage(t *testing.T) {

	body := map[string]string{
		"message": "",
	}

	jsonBody, _ := json.Marshal(body)

	req := httptest.NewRequest(
		http.MethodPost,
		"/echo",
		bytes.NewReader(jsonBody),
	)

	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	echoHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", rr.Code)
	}
}
