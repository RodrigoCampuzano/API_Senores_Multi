package main

import (
    "log"
    "APIs/src/core/db"
    ds18b20 "APIs/src/ds18b20/infraestructure/routes"
    max30102 "APIs/src/max30102/infraestructure/routes"
    mpu6050 "APIs/src/mpu6050/infraestructure/routes"
    ds18b20c "APIs/src/ds18b20/infraestructure/controllers"
    max30102c "APIs/src/max30102/infraestructure/controllers"
    mpu6050c "APIs/src/mpu6050/infraestructure/controllers"
    ds18b20s "APIs/src/ds18b20/application"
    max30102s "APIs/src/max30102/application"
    mpu6050s "APIs/src/mpu6050/application"
    mpu6050r "APIs/src/mpu6050/infraestructure/repositories"
    max30102r "APIs/src/max30102/infraestructure/repositories"
    ds18b20r "APIs/src/ds18b20/infraestructure/repositories"
    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func main() {
    // Cargar variables de entorno desde el archivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Inicializar la conexión a la base de datos
    database, err := db.InitDB()
    defer db.CloseDB(database)

    // Crear un enrutador Gin
    r := gin.Default()

    max30102Repo := max30102r.NewMax30102RepositoryPostgres(database)
    mpu6050Repo := mpu6050r.NewMPU6050RepositoryPostgres(database)
    ds18b20Repo := ds18b20r.NewDS18B20RepositoryPostgres(database)

    ds18b20Service := ds18b20s.NewDS18B20Service(ds18b20Repo)
    max30102Service := max30102s.NewMax30102Service(max30102Repo)
    mpu6050Service := mpu6050s.NewMpu6050Service(mpu6050Repo)

    ds18b20Ctrl := ds18b20c.NewDS18B20Controller(ds18b20Service)
    max30102Ctrl := max30102c.NewMax30102Controller(max30102Service)
    mpu6050Ctrl := mpu6050c.NewMpu6050Controller(mpu6050Service)
    // Configurar las rutas para cada módulo
    ds18b20.DS18B20Routes(r, ds18b20Ctrl)
    max30102.Max30102hRoutes(r, max30102Ctrl)
    mpu6050.MPU6050Routes(r, mpu6050Ctrl)

    // Iniciar el servidor en el puerto 8080
    r.Run(":8080")
}