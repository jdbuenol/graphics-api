package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/drew138/graphics-api/routes"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/acme/autocert"
)

// https://medium.com/rungo/secure-https-servers-in-go-a783008b36da
func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache("cert-cache"),
		HostPolicy: autocert.HostWhitelist("drew-graphics.site"),
	}

	server := &http.Server{
		Addr:    ":443",
		Handler: r,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
		},
	}
	go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	log.Println("Server started, running on port 443.")
	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatal("Server failed to start: ", err.Error())
	}
}
