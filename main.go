package main

import (
	"APIs/src/core/db"
	"APIs/src/core/middleware"
	"log"
	ds18b20s "APIs/src/ds18b20/application"
	ds18b20c "APIs/src/ds18b20/infraestructure/controllers"
	ds18b20r "APIs/src/ds18b20/infraestructure/repositories"
	ds18b20 "APIs/src/ds18b20/infraestructure/routes"
	max30102s "APIs/src/max30102/application"
	max30102r "APIs/src/max30102/infraestructure/repositories"
	max30102 "APIs/src/max30102/infraestructure/routes"
	mpu6050s "APIs/src/mpu6050/application"
	mpu6050c "APIs/src/mpu6050/infraestructure/controllers"
	mpu6050r "APIs/src/mpu6050/infraestructure/repositories"
	mpu6050 "APIs/src/mpu6050/infraestructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Cargar variables de entorno desde el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error cargando el archivo .env")
    }

    // Inicializar la conexión a la base de datos
    database, err := db.InitDB()
    if err != nil {
        log.Fatal("Error al inicializar la base de datos: ", err)
    }
    defer db.CloseDB(database)

    // Crear un enrutador Gin
    r := gin.Default()
    
    r.Use(middleware.MiddlewareCORS())

    // Inicialización de repositorios
    max30102Repo := max30102r.NewMax30102RepositoryMySQL(database)
    mpu6050Repo := mpu6050r.NewMPU6050RepositoryMySQL(database)
    ds18b20Repo := ds18b20r.NewDS18B20RepositoryMySQL(database)

    // Inicialización de servicios
    ds18b20Service := ds18b20s.NewDS18B20Service(ds18b20Repo)
    max30102Service := max30102s.NewMax30102Service(max30102Repo)
    mpu6050Service := mpu6050s.NewMpu6050Service(mpu6050Repo)

    // Inicialización de controladores
    ds18b20Ctrl := ds18b20c.NewDS18B20Controller(ds18b20Service)
    mpu6050Ctrl := mpu6050c.NewMpu6050Controller(mpu6050Service)

    // Configurar las rutas para cada módulo
    ds18b20.DS18B20Routes(r, ds18b20Ctrl)
    max30102.Max30102Routes(r, max30102Service)
    mpu6050.MPU6050Routes(r, mpu6050Ctrl)

    // Iniciar el servidor en el puerto 8080
    if err := r.Run(":8081"); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
