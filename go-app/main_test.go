package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHandler(t *testing.T) {
    // Create a new HTTP request with any method and path (in this case, "/")
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatalf("failed to create request: %v", err)
    }

    // Create a ResponseRecorder (which implements http.ResponseWriter) to record the response
    rr := httptest.NewRecorder()

    // Call the handler function directly, passing in the ResponseRecorder and the dummy request
    handler(rr, req)

    // Check the status code of the response (should be http.StatusOK)
    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Check the body of the response (should be "Hello, Go!")
    expected := "Hello, Go!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

