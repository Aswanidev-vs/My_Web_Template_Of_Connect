package handler

import (
	"log"
	"net/http"

	"github.com/Aswanidev-vs/Connect/db"
	tpl "github.com/Aswanidev-vs/Connect/templates"
	"golang.org/x/crypto/bcrypt"
)

type SignupRequest struct {
	Username   string
	Email      string
	Password   string
	RePassword string
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		err := tpl.Templates["signup"].Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	req := SignupRequest{
		Username:   r.FormValue("username"),
		Email:      r.FormValue("email"),
		Password:   r.FormValue("password"),
		RePassword: r.FormValue("repassword"),
	}

	// Basic validation
	if req.Username == "" || req.Email == "" || req.Password == "" || req.RePassword == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	if req.Password != req.RePassword {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	exists, err := db.CheckUser(req.Email)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	if exists {
		// http.Error(w, "Email already registered", http.StatusConflict)
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
    // <p>Email already registered</p>
    <script>
            window.location.href = '/signup';
    </script>
`))

		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err := db.NewUser(req.Username, req.Email, string(hashedPassword)); err != nil {
		log.Println("Error inserting user:", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// ✅ Redirect to login page after success
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
