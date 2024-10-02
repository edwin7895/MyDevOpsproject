package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

// Estructura para manejar los datos del formulario
type ContactForm struct {
    Name    string `json:"name"`
    Email   string `json:"email"`
    Message string `json:"message"`
}

// Middleware para habilitar CORS
func enableCors(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func main() {
    // Crear un router
    mux := http.NewServeMux()

    // Registrar rutas
    mux.HandleFunc("/", HomeHandler)
    mux.HandleFunc("/api/contact", contactHandler)

    // Registrar un manejador de rutas no encontradas (404)
    mux.HandleFunc("/404", notFoundHandler)

    // Asignar el manejador 404 en el router
    mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/" && r.URL.Path != "/api/contact" {
            notFoundHandler(w, r)
            return
        }
        mux.ServeHTTP(w, r)
    }))

    fmt.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}

// Manejador de la ruta principal "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
    enableCors(w)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Welcome to the Backend API!"}`))
}

// Manejador para la ruta "/api/contact"
func contactHandler(w http.ResponseWriter, r *http.Request) {
    enableCors(w)

    // Manejar solicitud preflight OPTIONS
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var contactForm ContactForm

    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = json.Unmarshal(body, &contactForm)
    if err != nil {
        http.Error(w, "Could not parse JSON", http.StatusBadRequest)
        return
    }

    fmt.Printf("Received contact form submission: %+v\n", contactForm)

    response := map[string]string{
        "message": "Thank you, " + contactForm.Name + ". Your message has been received.",
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// Manejador para rutas no encontradas
func notFoundHandler(w http.ResponseWriter, r *h
