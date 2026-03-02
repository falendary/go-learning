package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

/*
This is the simplest possible
server-to-server request example.
*/

func main() {

	// =========================
	// Simple GET request
	// =========================

	resp, err := http.Get("http://localhost:8080/health")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var health map[string]any

	// Decode JSON directly from response body
	err = json.NewDecoder(resp.Body).Decode(&health)
	if err != nil {
		panic(err)
	}

	fmt.Println("GET /health response:")
	fmt.Println(health)

	// =========================
	// Simple POST request
	// =========================

	// Create request body
	payload := map[string]string{
		"message": "hello from go",
	}

	bodyBytes, _ := json.Marshal(payload)

	resp2, err := http.Post(
		"http://localhost:8080/echo",
		"application/json",
		bytes.NewReader(bodyBytes),
	)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	var echoResp map[string]any

	err = json.NewDecoder(resp2.Body).Decode(&echoResp)
	if err != nil {
		panic(err)
	}

	fmt.Println("POST /echo response:")
	fmt.Println(echoResp)
}
