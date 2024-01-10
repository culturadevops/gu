package httprequest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/culturadevops/gu/httpRequest"
)

type TestData struct {
	Message string `json:"message"`
}

func TestPost(t *testing.T) {
	// Crear un servidor HTTP de prueba
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simular una respuesta JSON exitosa
		data := TestData{Message: "Success"}
		jsonBytes, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}))
	defer ts.Close()

	// Preparar datos de prueba
	formData := url.Values{}
	formData.Set("key", "value")

	var result TestData

	// Realizar la solicitud POST al servidor de prueba
	err := httpRequest.Post(ts.URL, formData, &result)
	if err != nil {
		t.Errorf("Error al hacer la solicitud POST: %s", err)
		return
	}

	// Verificar si se obtuvo el mensaje esperado
	expectedMessage := "Success"
	if result.Message != expectedMessage {
		t.Errorf("Se esperaba el mensaje '%s', pero se obtuvo '%s'", expectedMessage, result.Message)
	}
}

func TestGet(t *testing.T) {
	// Crear un servidor HTTP de prueba
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simular una respuesta JSON exitosa
		data := TestData{Message: "Success"}
		jsonBytes, _ := json.Marshal(data)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}))
	defer ts.Close()

	var result TestData

	// Realizar la solicitud GET al servidor de prueba
	err := httpRequest.Get(ts.URL, &result)
	if err != nil {
		t.Errorf("Error al hacer la solicitud GET: %s", err)
		return
	}

	// Verificar si se obtuvo el mensaje esperado
	expectedMessage := "Success"
	if result.Message != expectedMessage {
		t.Errorf("Se esperaba el mensaje '%s', pero se obtuvo '%s'", expectedMessage, result.Message)
	}
}
