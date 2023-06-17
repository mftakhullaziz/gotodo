package accounts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestLoginEndpoint(t *testing.T) {
	// Create a new router
	router := httprouter.New()

	// Define the login endpoint
	router.POST("/login", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		// Create a response map
		response := map[string]string{
			"message": "Login successful",
		}

		// Convert the response to JSON
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Set the response headers and write the response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	})

	// Create a test request
	req, err := http.NewRequest("POST", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Serve the request using the router
	router.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d but got %d", http.StatusOK, rr.Code)
	}

	// Decode the response body
	var response map[string]string
	err = json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		t.Fatal(err)
	}

	// Check the response message
	expectedMessage := "Login successful"
	if response["message"] != expectedMessage {
		t.Errorf("expected message %s but got %s", expectedMessage, response["message"])
	}
}
