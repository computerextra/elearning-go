package routes

import (
	"computerextra/elaerning-go/templates"

	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	tokens = make(map[string]time.Time)
	mu     sync.Mutex
)

func generateToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: Datenbank verbindung einbauen.
	// TODO: Passwort Hash generieren
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == "user" && password == "pass" {
			token, err := generateToken()
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			mu.Lock()
			tokens[token] = time.Now().Add(30 * time.Minute)
			mu.Unlock()
			fmt.Fprintf(w, "Login successful. Your token: %s", token)
			return
		}
	}
	w.Header().Add("Content-Type", "text/html")
	templates.Login().Render(r.Context(), w)
	http.ServeFile(w, r, "login.html")
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		mu.Lock()
		expiry, exists := tokens[token]
		mu.Unlock()
		if !exists || time.Now().After(expiry) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Erneuern der Session
		mu.Lock()
		tokens[token] = time.Now().Add(30 * time.Minute)
		mu.Unlock()
		next.ServeHTTP(w, r)
	})
}
