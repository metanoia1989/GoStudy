package main

import "net/http"

func main() {
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}
	server.ListenAndServe()
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
