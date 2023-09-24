package backend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {

	server := CreateServer()
	server.RegisterRoutes()

	go func() {
		err := http.ListenAndServe(":8080", server.Router)
		if err != nil {
			t.Fail()
		}
	}()

	time.Sleep(100 * time.Millisecond)
	// Define the request body
	requestBody := RequestNote{
		Content: "Hello World",
		Tags:    []string{"Hello"},
	}

	// Convert the request body to JSON
	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Define the API endpoint URL
	apiURL := "http://localhost:8080/note/"

	// Send the POST request
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Fatal(resp.StatusCode)
	}

}
