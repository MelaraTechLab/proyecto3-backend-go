package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"proyecto3-backend-go/db"
	"proyecto3-backend-go/handlers"
)

func main() {
	db.Connect()

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
r.Use(func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
})


	r.HandleFunc("/reporte1", handlers.Reporte1Handler).Methods("GET")
	r.HandleFunc("/reporte2", handlers.Reporte2Handler).Methods("GET")
	r.HandleFunc("/reporte3", handlers.Reporte3Handler).Methods("GET")
	r.HandleFunc("/reporte4", handlers.Reporte4Handler).Methods("GET")
	r.HandleFunc("/reporte5", handlers.Reporte5Handler).Methods("GET")

	log.Println("Servidor corriendo en http://localhost:3001")
	http.ListenAndServe(":3001", r)
}