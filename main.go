package main

import (
	"log"
	"net/http"

	"github.com/Aswanidev-vs/Connect/config"
	"github.com/Aswanidev-vs/Connect/dashboard"
	"github.com/Aswanidev-vs/Connect/handler"
	tpl "github.com/Aswanidev-vs/Connect/templates"
)

func main() {
	config.LoadEnv()
	// Initialize database

	if err := config.InitDB(); err != nil {
		log.Fatal("Database init error:", err)
	}
	defer config.DB.Close()
	// Load templates
	tpl.LoadTemplates()

	// Serve static files (CSS, JS, images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Register routes
	http.HandleFunc("/", handler.IndexHandler) // homepage
	http.HandleFunc("/signup", handler.SignupHandler)
	http.HandleFunc("/login", handler.LoginHandler) // if you add login
	http.HandleFunc("/dashboard", handler.JWTMiddleware(dashboard.DashboardHandler))
	log.Println("Server started at http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))

}
