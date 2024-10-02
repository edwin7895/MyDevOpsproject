package main

import (
    "bytes"
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

    expected := `{"message":"Thank you, John Doe. Your message has been received."}`
    if rr.Body.String() != expected {
        t.Errorf("handler devolvió respuesta inesperada: obtuvo %v esperaba %v", rr.Body.String(), expected)
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

    expected := "Not found\n"
    if rr.Body.String() != expected {
        t.Errorf("handler devolvió respuesta inesperada: obtuvo %v esperaba %v", rr.Body.String(), expected)
    }
}
