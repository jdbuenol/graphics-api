package main

import (
	"github.com/gin-gonic/gin"

	router "github.com/drew138/graphics-api/api/routes"
)

func main() {
	// r := mux.NewRouter()
	// routes.RegisterRoutes(r)
	//
	// certManager := autocert.Manager{
	// 	Prompt:     autocert.AcceptTOS,
	// 	Cache:      autocert.DirCache("cert-cache"),
	// 	HostPolicy: autocert.HostWhitelist("drew-graphics.site", "www.drew-graphics.site"),
	// }
	//
	// server := &http.Server{
	// 	Addr:    ":443",
	// 	Handler: r,
	// 	TLSConfig: &tls.Config{
	// 		GetCertificate: certManager.GetCertificate,
	// 	},
	// }
	// go http.ListenAndServe(":80", certManager.HTTPHandler(nil))
	// log.Println("Server started, running on port 443.")
	// if err := server.ListenAndServeTLS("", ""); err != nil {
	// 	log.Fatal("Server failed to start: ", err.Error())
	// }

	eng := gin.Default()

	router := router.NewRouter(eng)
	router.MapRoutes()

	if err := eng.Run(); err != nil {
		panic(err)
	}

}
