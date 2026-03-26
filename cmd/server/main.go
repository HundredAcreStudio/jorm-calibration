package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/HundredAcreStudio/jorm-calibration/internal/handler"
	"github.com/HundredAcreStudio/jorm-calibration/internal/store"
)

func main() {
	db := store.NewMemoryStore()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", handler.ListUsers(db))
	mux.HandleFunc("POST /users", handler.CreateUser(db))
	mux.HandleFunc("GET /users/{id}", handler.GetUser(db))
	mux.HandleFunc("DELETE /users/{id}", handler.DeleteUser(db))

	addr := ":8080"
	fmt.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
