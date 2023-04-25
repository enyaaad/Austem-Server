package main

import (
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
)

var users = map[string]string{
	"asd": "123",
}

func isAuthorised(username, password string) bool {
	pass, ok := users[username]
	if !ok {
		return false
	}

	return password == pass
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200/")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Authorization")
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "No basic auth present"}`))
		return
	}

	if !isAuthorised(username, password) {
		w.Header().Add("WWW-Authenticate", `Basic realm="Give username and password"`)
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"message": "Invalid username or password"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "welcome to golang world!"}`))
	return
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			log.Print("preflight detected: ", r.Header)
			w.Header().Add("Connection", "keep-alive")
			w.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200/")
			w.Header().Add("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PUT")
			w.Header().Add("Access-Control-Allow-Headers", "content-type, authorization")
			w.Header().Add("Access-Control-Max-Age", "86400")
		} else {
			h.ServeHTTP(w, r)
		}
	}
}
func main() {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200/", "http://localhost:8080/example"},
	})
	handler := http.HandlerFunc(greeting)
	err := http.ListenAndServe(":8080", corsHandler(c.Handler(handler)))
	if err != nil {
		return
	}
	fmt.Println("Starting Server at port :8080")
}
