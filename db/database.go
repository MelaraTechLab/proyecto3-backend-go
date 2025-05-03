package db

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    var err error

    connStr := "host=localhost port=5432 user=postgres password=tu_clave dbname=eventos_culturales sslmode=disable"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Error al conectar a la base de datos:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("No se pudo hacer ping a la base de datos:", err)
    }

    log.Println("Conectado a PostgreSQL exitosamente.")
}
