package main

import (
    "bytes"
    "encoding/json" // Importar el paquete json	
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHomeHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(HomeHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler devolvió status incorrecto: obtuvo %v esperaba %v", status, http.StatusOK)
    }

    expected := `{"message": "Welcome to the Backend API!"}`
    if rr.Body.String() != expected {
        t.Errorf("handler devolvió respuesta inesperada: obtuvo %v esperaba %v", rr.Body.String(), expected)
    }
}

func TestContactHandler(t *testing.T) {
    contactJSON := []byte(`{"name":"John Doe", "email":"john@example.com", "message":"Hello!"}`)
    req, err := http.NewRequest("POST", "/api/contact", bytes.NewBuffer(contactJSON))
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(contactHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler devolvió status incorrecto: obtuvo %v esperaba %v", status, http.StatusOK)
    }

    // Convertir la respuesta esperada y la obtenida en JSON
    var expected map[string]string
    var actual map[string]string

    expectedStr := `{"message":"Thank you, John Doe. Your message has been received."}`
    err = json.Unmarshal([]byte(expectedStr), &expected)
    if err != nil {
        t.Fatalf("No se pudo parsear la respuesta esperada: %v", err)
    }

    err = json.Unmarshal(rr.Body.Bytes(), &actual)
    if err != nil {
        t.Fatalf("No se pudo parsear la respuesta obtenida: %v", err)
    }

    // Comparar los valores de los mensajes
    if actual["message"] != expected["message"] {
        t.Errorf("handler devolvió respuesta inesperada: obtuvo %v esperaba %v", actual["message"], expected["message"])
    }
}


func TestNotFoundHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/not-found", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(notFoundHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf("handler devolvió status incorrecto: obtuvo %v esperaba %v", status, http.StatusNotFound)
    }

    // Cambiar la respuesta esperada según el mensaje real devuelto por el handler
    expected := `{"error": "404 - Resource not found"}`
    if rr.Body.String() != expected {
        t.Errorf("handler devolvió respuesta inesperada: obtuvo %v esperaba %v", rr.Body.String(), expected)
    }
}
