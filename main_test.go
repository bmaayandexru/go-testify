package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerCorrectRequest(t *testing.T) {
	// Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
	// формируем корректный запрос
	req := httptest.NewRequest("GET", "/cafe?count=3&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверяем статус на http.StatusOK
	assert.Equal(t, responseRecorder.Code, http.StatusOK)

	// читаем тело в строку
	body := responseRecorder.Body.String()
	// проверяем не пустое ли тело
	assert.NotEmpty(t, body)
}

func TestMainHandlerCityIsMissing(t *testing.T) {
	// Город, который передаётся в параметре city, не поддерживается.
	// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
	// формируем запрос с несуществующим городом
	req := httptest.NewRequest("GET", "/cafe?count=23&city=piter", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверяем ответ на 400 http.StatusBadRequest
	assert.Equal(t, responseRecorder.Code, http.StatusBadRequest)

	// и тело == wrong city value
	body := responseRecorder.Body.String()
	assert.Equal(t, body, "wrong city value")
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	// Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
	// формируем запрос с count > 4 и корректным городом
	req := httptest.NewRequest("GET", "/cafe?count=12&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	// проверяем что вернулось totalCount
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")
	require.Len(t, list, totalCount)
}
