package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenReqOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Status should be OK")
	body := responseRecorder.Body
	assert.NotEmpty(t, body, "Body should be not empty")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.Equal(t, http.StatusOK, responseRecorder.Code, "Status code should be OK")

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount, fmt.Sprintf("Total count: %d, should be %d", len(list), totalCount))
}

func TestMainHandlerUnkownCity(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=berlin", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest, "Status code should be BadRequest")
	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body, "Body response should be \"wrong city value\" ")
}
