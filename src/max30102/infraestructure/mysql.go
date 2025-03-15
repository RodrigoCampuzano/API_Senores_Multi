package infraestructure

import (
	core "APIs/src/core/db"
	"APIs/src/max30102/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// 🔹 Guardar un producto
func (mysql *MySQL) Save(data *entities.Max30102) error {
	query := `INSERT INTO max30102 (bpm, spo2) VALUES (?, ?)`
	result, err := mysql.conn.DB.Exec(query, data.BPM, data.SpO2)
	if err != nil {
		log.Println("Error insertando producto:", err)
		return err
	}
	lastID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error obteniendo el ID del nuevo dato:", err)
		return err
	}
	fmt.Println("Nuevo dato almacenado con ID:", lastID)
	return nil
}

func (mysql *MySQL) GetAll() ([]*entities.Max30102, error) {
	query := `SELECT id, bpm, spo2 FROM max30102`
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		log.Println("Error consultando datos:", err)
		return nil, err
	}
	defer rows.Close()

	var products []*entities.Max30102
	for rows.Next() {
		var product entities.Max30102
		err := rows.Scan(&product.ID, &product.SpO2, &product.BPM)
		if err != nil {
			log.Println("Error leyendo fila:", err)
			continue
		}
		products = append(products, &product)
	}

	if err := rows.Err(); err != nil {
		log.Println("Error durante la iteración de las filas:", err)
		return nil, err
	}
	if len(products) == 0 {
		return []*entities.Max30102{}, nil
	}

	return products, nil
}
