package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"os"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", handlePage)

	// os.Setenv("PORT", ":8000")
	port := os.Getenv("PORT")

	srv := http.Server{
		Handler:      m,
		Addr:         port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// this blocks forever, until the server
	// has an unrecoverable error
	fmt.Println("server started on ", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.WriteHeader(200)
	const page = `<html>
<head></head>
<body>
	<p> Hello from Docker! I'm a Go server. </p>
</body>
</html>
`
	w.Write([]byte(page))
}