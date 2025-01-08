package main

import (
	"context"
	"encoding/json"
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
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getAllUserHandler(w, r)
		case "POST":
			createUserHandler(w, r)
		default:
			http.Error(w, "The Request Method Not Allowed", http.StatusBadRequest)
		}
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

func getAllUserHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := conn()
	if err != nil {
		writeError(w, err)
		return
	}
	defer conn.Close()

	qry, err := conn.QueryContext(context.Background(), "SELECT * FROM users")
	if err != nil {
		writeError(w, err)
		return
	}

	result := make([]User, 0)
	for qry.Next() {
		var user User
		err = qry.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Birth)
		if err != nil {
			writeError(w, err)
			return
		}
		result = append(result, user)
	}
	writeData(w, result)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	payload := new(User)

	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		writeError(w, err)
		return
	}

	conn, err := conn()
	if err != nil {
		writeError(w, err)
		return
	}
	defer conn.Close()

	stmt, err := conn.PrepareContext(context.Background(), "INSERT INTO users (first_name, last_name, birth) VALUES (?, ?, ?)")
	if err != nil {
		writeError(w, err)
		return
	}

	stmtRes, err := stmt.ExecContext(context.Background(), payload.FirstName, payload.LastName, payload.Birth)
	if err != nil {
		writeError(w, err)
		return
	}

	id, _ := stmtRes.LastInsertId()
	result := map[string]interface{}{"LastInsertId": id}
	writeData(w, result)
}
