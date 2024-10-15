package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJson(t *testing.T) {

	// create a request

	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)

	// create a response recorder

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(testApp.GetAllDogBreedsJson)

	// server handler

	handler.ServeHTTP(rr, req)

	// check response

	if rr.Code != http.StatusOK {
		t.Errorf("Wrong response code; expected %d, got %d", http.StatusOK, rr.Code)
	}

}
