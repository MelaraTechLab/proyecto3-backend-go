package handlers

import (
    "encoding/json"
    "net/http"
    "proyecto3-backend-go/db"
)

// --- Estructuras ---
type Evento struct {
    Nombre string `json:"nombre"`
    Fecha  string `json:"fecha"`
    Lugar  string `json:"lugar"`
    Tipo   string `json:"tipo"`
}

type Conteo struct {
    Clave string `json:"clave"`
    Total int    `json:"total"`
}

// --- Reporte 1: Eventos por fecha y tipo ---
func Reporte1Handler(w http.ResponseWriter, r *http.Request) {
    desde := r.URL.Query().Get("desde")
    hasta := r.URL.Query().Get("hasta")
    tipo := r.URL.Query().Get("tipo")
    lugar := r.URL.Query().Get("lugar")

    query := `
    SELECT e.nombre, e.fecha, l.nombre, t.nombre
    FROM evento e
    JOIN lugar l ON e.lugar_id = l.id
    JOIN tipoevento t ON e.tipo_id = t.id
    WHERE ($1 = '' OR e.fecha >= $1::DATE)
      AND ($2 = '' OR e.fecha <= $2::DATE)
      AND ($3 = '' OR t.nombre = $3)
      AND ($4 = '' OR l.nombre = $4)
    ORDER BY e.fecha;
    `

    rows, err := db.DB.Query(query, desde, hasta, tipo, lugar)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var eventos []Evento
    for rows.Next() {
        var e Evento
        rows.Scan(&e.Nombre, &e.Fecha, &e.Lugar, &e.Tipo)
        eventos = append(eventos, e)
    }

    json.NewEncoder(w).Encode(eventos)
}


// --- Reporte 2: Total de asistentes por evento ---
func Reporte2Handler(w http.ResponseWriter, r *http.Request) {
    query := `
        SELECT e.nombre, COUNT(a.asistente_id) as total
        FROM evento e
        JOIN asistencia a ON e.id = a.evento_id
        GROUP BY e.nombre
        ORDER BY total DESC;
    `
    rows, err := db.DB.Query(query)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var resultados []Conteo
    for rows.Next() {
        var c Conteo
        rows.Scan(&c.Clave, &c.Total)
        resultados = append(resultados, c)
    }

    json.NewEncoder(w).Encode(resultados)
}

// --- Reporte 3: Eventos por lugar ---
func Reporte3Handler(w http.ResponseWriter, r *http.Request) {
    lugar := r.URL.Query().Get("lugar")

    query := `
        SELECT l.nombre, COUNT(*) as total
        FROM evento e
        JOIN lugar l ON e.lugar_id = l.id
        WHERE ($1 = '' OR l.nombre = $1)
        GROUP BY l.nombre;
    `
    rows, err := db.DB.Query(query, lugar)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var resultados []Conteo
    for rows.Next() {
        var c Conteo
        rows.Scan(&c.Clave, &c.Total)
        resultados = append(resultados, c)
    }

    json.NewEncoder(w).Encode(resultados)
}

// --- Reporte 4: Asistencias por género ---
func Reporte4Handler(w http.ResponseWriter, r *http.Request) {
    genero := r.URL.Query().Get("genero")

    query := `
        SELECT a.genero, COUNT(*) as total
        FROM asistencia s
        JOIN asistente a ON s.asistente_id = a.id
        WHERE ($1 = '' OR a.genero = $1)
        GROUP BY a.genero;
    `
    rows, err := db.DB.Query(query, genero)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var resultados []Conteo
    for rows.Next() {
        var c Conteo
        rows.Scan(&c.Clave, &c.Total)
        resultados = append(resultados, c)
    }

    json.NewEncoder(w).Encode(resultados)
}

// --- Reporte 5: Patrocinadores por evento ---
func Reporte5Handler(w http.ResponseWriter, r *http.Request) {
    evento := r.URL.Query().Get("evento")

    query := `
        SELECT e.nombre, COUNT(ep.patrocinador_id) as total
        FROM evento e
        JOIN evento_patrocinador ep ON e.id = ep.evento_id
        WHERE ($1 = '' OR e.nombre = $1)
        GROUP BY e.nombre;
    `
    rows, err := db.DB.Query(query, evento)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    defer rows.Close()

    var resultados []Conteo
    for rows.Next() {
        var c Conteo
        rows.Scan(&c.Clave, &c.Total)
        resultados = append(resultados, c)
    }

    json.NewEncoder(w).Encode(resultados)
}
