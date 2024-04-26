package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// kita buat map untuk menyimpan username dan password. Pada aplikasi sesungguhnya, ini menggunakan database
var users = map[string]string{
	"khairul": "12345",
	"dito":    "4321",
}

// map ini akan menyimpan users sessions.
var sessions = map[string]session{}

type session struct {
	username string
	expiry   time.Time
}

// kita akan menggunakan method ini untuk menentukan apakah session telah expired
func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get JSON body dan decode menjadi credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get expected password dari memory map
	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		// Password tidak sama dengan password yang diterima
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Membuat random session token menggunakan "github.com/google/uuid" library untuk generate UUIDs
	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(30 * time.Minute)

	// Set token pada session map, bersama dengan informasi dari session
	sessions[sessionToken] = session{
		username: creds.Username,
		expiry:   expiresAt,
	}

	// Terakhir, kita menetapkan client cookie dengan nama "session_token" sebagai session token yang telah kita generate
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt, // kita juga menetapkan expired 120 detik
	})

	w.Write([]byte(fmt.Sprintf("Login success with token %s", sessionToken)))
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Kita dapat memperoleh session token dari requests cookies, ini disertakan dalam setiap request
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// cookie tidak di set, return unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Maaf, anda harus login!"))
				return
			}
			// Untuk jenis error lainnya, return bad request status
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Maaf, anda harus login!"))
			return
		}
		sessionToken := c.Value

		// Kita mendapatkan userSession dari session map
		userSession, exists := sessions[sessionToken]

		if !exists {
			// Session token tidak ada pada session map, return unauthorized error
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Maaf, anda harus login!"))
			return
		}

		if userSession.isExpired() {
			delete(sessions, sessionToken)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Maaf, anda harus login!"))
			return
		}

		ctx := context.WithValue(r.Context(), "username", userSession.username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Admin(w http.ResponseWriter, r *http.Request) {
	username := fmt.Sprintf("%s", r.Context().Value("username"))
	w.Write([]byte(fmt.Sprintf("Welcome %s!", username)))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			// Jika cookie tidak di set, return unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// Untuk jenis error lainnya, return bad request status
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	sessionToken := c.Value

	// hapus users session dari session map
	delete(sessions, sessionToken)

	// Kita ubah nilai cookie dari user menjadi kosong dan tetapkan waktu expired  menjadi waktu saat ini
	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
	})

	w.Write([]byte("Logout success"))
}

func main() {
	http.HandleFunc("/login", Login)
	http.Handle("/admin", Authenticate(http.HandlerFunc(Admin)))
	http.HandleFunc("/logout", Logout)

	http.ListenAndServe(":8080", nil)
}
