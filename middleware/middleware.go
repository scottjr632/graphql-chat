package middleware

import (
	"log"
	"net/http"
	"time"
)

func Auth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("auth")
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			return
		}
		if token.String() == "authenticate" {
			next.ServeHTTP(w, r)
		} else {
			w.WriteHeader(401)
			return
		}
	}
}

func CORS(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, access-control-allow-origin")

		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	}
}

func LoggingHandler(next http.Handler) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		next.ServeHTTP(w, r)
		t2 := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
	}

	return fn
}
