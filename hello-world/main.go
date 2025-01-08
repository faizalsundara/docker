package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env is required")
	}

	instanceID := os.Getenv("INSTANCE_ID")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "The Request Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		text := "Hello World"
		if instanceID != "" {
			text += ". from " + instanceID
		}
		w.Write([]byte(text))
	})

	server := new(http.Server)
	server.Handler = mux
	server.Addr = ":" + port

	log.Println("web server is strting at", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
