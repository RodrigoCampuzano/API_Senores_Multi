package db

import (
    "database/sql"
    "fmt"
    "os"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func InitDB() (*sql.DB, error) {
    connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&tls=false",
        os.Getenv("DB_USER"), 
        os.Getenv("DB_PASS"),
        os.Getenv("DB_HOST"), 
        os.Getenv("DB_PORT"), 
        os.Getenv("DB_NAME"))

    // Intentar abrir la conexión con la base de datos
    db, err := sql.Open("mysql", connStr)
    if err != nil {
        log.Printf("Error al abrir la base de datos: %v", err)
        return nil, err
    }

    // Intentar hacer un ping a la base de datos para verificar la conexión
    err = db.Ping()
    if err != nil {
        log.Printf("Error al conectar a la base de datos: %v", err)
        return nil, err
    }

    fmt.Println("Conexión exitosa a la base de datos")
    return db, nil
}

func CloseDB(db *sql.DB) {
    if err := db.Close(); err != nil {
        log.Printf("Error al cerrar la base de datos: %v", err)
    }
}
