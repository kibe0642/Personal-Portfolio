package main

import (
    "fmt"
    "log"
    "net/http"
    "kibet-kevin-portfolio/internal/handler"
    "kibet-kevin-portfolio/db"
)

func main() {
    // Initialize database connection
    db.InitDB()

    // Serve static assets (CSS, JS, images)
    http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

    // Define route handlers
    http.HandleFunc("/", serveTemplate("index.html"))
    http.HandleFunc("/about", serveTemplate("about.html"))
    http.HandleFunc("/projects", serveTemplate("projects.html"))
    http.HandleFunc("/contact", handler.ContactHandler)

    // Start the server
    fmt.Println("Server running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("Error starting server:", err)
    }
}

// serveTemplate is used to render HTML templates
func serveTemplate(page string) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tmpl, err := template.ParseFiles("templates/" + page)
        if err != nil {
            http.Error(w, "Unable to load page", http.StatusInternalServerError)
            return
        }
        tmpl.Execute(w, nil)
    }
}
