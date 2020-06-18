// endpoints_test.go
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheckHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheckHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"alive": true}`

	assert.Equal(t, rr.Body.String(), expected, "Got back alive response")
}

func TestSearchHandlerEmpty(t *testing.T) {
	req, err := http.NewRequest("POST", "/search", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusBadRequest, "Got StatusBadRequest status")
}

func TestSearchHandlerQuery(t *testing.T) {

	payload := []byte(`{"search":"luke"}`)

	req, err := http.NewRequest("POST", "/search", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "Got OK status")

}

func TestIndexHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	data, err := ioutil.ReadFile(`data/documents/test.json`)
	if err != nil {
		log.Fatalf("Error opening file to index: %s", err)
	}

	jsonString := string(data)
	payload := []byte(jsonString)

	req, err := http.NewRequest("POST", "/index", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(SearchHandler)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusOK, "Got OK status")

}

func TestCreateIndex(t *testing.T) {
	createIndex()
}
