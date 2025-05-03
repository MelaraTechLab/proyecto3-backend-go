# Proyecto 3 – Backend en Go

Este es el backend del Proyecto 3 de Bases de Datos 1: **Gestión de eventos culturales y asistencia**. La API fue desarrollada en Go y conecta a una base de datos PostgreSQL.

## 🔧 Tecnologías

- Lenguaje: Go
- Librerías: Gorilla Mux, lib/pq
- Base de datos: PostgreSQL

## ⚙️ Instalación

### 1. Clona este repositorio:

```bash
git clone https://github.com/MelaraTechLab/proyecto3-backend-go.git
cd proyecto3-backend-go
```

### 2. Instala dependencias:

```bash
go get github.com/gorilla/mux
go get github.com/lib/pq
```

### 3. Conecta a tu base de datos editando `db/database.go`:

```go
connStr := "host=localhost port=5432 user=postgres password=tu_clave dbname=eventos_culturales sslmode=disable"
```

### 4. Ejecuta el servidor:

```bash
go run main.go
```

## 📊 Endpoints disponibles

| Ruta        | Descripción                    | Filtros (por query params) |
| ----------- | ------------------------------ | -------------------------- |
| `/reporte1` | Eventos por fecha y tipo       | `desde`, `hasta`, `tipo`   |
| `/reporte2` | Total de asistentes por evento | —                          |
| `/reporte3` | Eventos por lugar              | `lugar`                    |
| `/reporte4` | Asistencias por género         | `genero`                   |
| `/reporte5` | Patrocinadores por evento      | `evento`                   |

## 🧪 Ejemplo de uso

```
GET http://localhost:3001/reporte1?desde=2025-05-01&hasta=2025-05-31&tipo=Festival
```

Devuelve una lista de eventos filtrados en formato JSON.

## ✍️ Autores

- Juan Cruz 23110
