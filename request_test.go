package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCEP(t *testing.T) {
	req, err := http.NewRequest("GET", "/cep?cep=50151250", nil)

	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	weatherHandler(rr, req)
	fmt.Println(rr.Body.String())

	assert.Contains(t, rr.Body.String(), "temp_C")
}
