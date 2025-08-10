package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/http2"
)

func main() {
	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from users")
	})

	port := 8080
	fmt.Printf("Server is running on port %d\n", port)

	cert := "cert.pem"
	key := "key.pem"
	
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler: nil,
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server , &http2.Server{})

	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatalf("Server failed to start: %v", err)
	// }

	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
