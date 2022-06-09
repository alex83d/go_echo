// Simple HTTP echo server written in Go.
package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// DefaultPort порт по умолчанию
const DefaultPort = "8080"
var port string


func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not find .env file:", err)
	}
	port = os.Getenv("SERVER_PORT")
	if port == "" {
		port = DefaultPort
	}
}

// EchoHandler 
func EchoHandler(w http.ResponseWriter, r *http.Request) {

    log.Println("echo" + " to client (" + r.RemoteAddr + ")")
	w.WriteHeader(200)

    w.Header().Set("Content-Type", "text/plain")

    b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(string(b))
    io.WriteString(w, string(b))
}

func main() {

    log.Println("Starting server, listening on port: " + port)

    http.HandleFunc("/", EchoHandler)
    log.Fatal(http.ListenAndServe(":" + port, nil))
}