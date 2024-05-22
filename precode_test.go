package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Equal(t, totalCount, len(list))

}

func TestMainHandlerWhenWrongCity(t *testing.T) {

	req := httptest.NewRequest("GET", "/cafe?count=10&city=tokyo", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	realBody := responseRecorder.Body.String()
	expectedBody := "wrong city value"
	realStatus := responseRecorder.Code
	expectedStatus := http.StatusBadRequest

	assert.Equal(t, realStatus, expectedStatus)
	assert.Equal(t, realBody, expectedBody)

}
