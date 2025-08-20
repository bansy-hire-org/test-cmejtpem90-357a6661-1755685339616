package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"encoding/json"
)

func TestGetLicenses(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/licenses", nil)
	if err != nil {
		t.Fatal(err)
	}

	rw := httptest.NewRecorder()
	handler := http.HandlerFunc(getLicenses)

	handler.ServeHTTP(rw, req)

	if status := rw.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var licenses []License
	err = json.Unmarshal(rw.Body.Bytes(), &licenses)
	if err != nil {
		t.Fatalf("Error unmarshaling response body: %v", err)
	}

	if len(licenses) != 3 {
		t.Errorf("handler returned wrong number of licenses: got %v want %v", len(licenses), 3)
	}

	if licenses[0].Name != "License A" {
		t.Errorf("handler returned wrong license name: got %v want %v", licenses[0].Name, "License A")
	}
}
