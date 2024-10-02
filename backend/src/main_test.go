package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomeHandler(t *testing.T) {
	// Prueba para el código 200 en la ruta principal.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)
	handler.ServeHTTP(rr, req)

	// Verifica que la respuesta sea 200 OK.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler devolvió status incorrecto: obtuvo %v esperaba %v", status, http.StatusOK)
	}

	// Verifica que el contenido sea correcto.
	expected := `{"message": "Welcome to the Backend API!"}`
	if rr.Body.String() != expected {
		t.Errorf("handler devolvió body incorrecto: obtuvo %v esperaba %v", rr.Body.String(), expected)
	}
}

func TestNotFoundHandler(t *testing.T) {
	// Prueba para una ruta no encontrada.
	req, err := http.NewRequest("GET", "/non-existent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)
	handler.ServeHTTP(rr, req)

	// Verifica que el código de respuesta sea 404.
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler devolvió status incorrecto: obtuvo %v esperaba %v", status, http.StatusNotFound)
	}
}
