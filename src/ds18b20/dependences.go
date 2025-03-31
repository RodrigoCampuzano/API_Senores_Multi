package infraestructure

import (
	"APIs/src/ds18b20/application"
	"APIs/src/ds18b20/domain/repositories"
	"APIs/src/ds18b20/infraestructure/controllers"
	"APIs/src/ds18b20/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

// InitTemperatura se encarga de inicializar las dependencias y rutas del módulo ds18b20.
func InitTemperatura(r *gin.Engine) {
    // Inicializa la conexión a la base de datos (comentado por ahora).
    // ps := NewMySQL()

    // Para pruebas, podemos pasar nil o un stub de ps

    // Carga las variables de entorno (comentado por ahora).
    // if err := godotenv.Load(); err != nil {
    // 	log.Fatalf("Error al cargar el archivo .env: %v", err)
    // }
	ds18b20Repo := repositories.NewDS18B20RepositoryStub()
    // Inicializa los casos de uso del módulo ds18b20.
    ds18b20Service := application.NewDS18B20Service(ds18b20Repo)
    //respuestaService := application.NewRespuestaService(nil)
	
    // Conexión al broker RabbitMQ (comentado por ahora).
    // host := os.Getenv("BROKER_HOST")
    // user := os.Getenv("BROKER_USER")
    // pass := os.Getenv("BROKER_PASS")
    // connURL := fmt.Sprintf("amqp://%s:%s@%s/", user, pass, host)
    // conn, err := amqp.Dial(connURL)
    // if err != nil {
    // 	log.Fatal("Error al conectar con RabbitMQ:", err)
    // }
    // publisher, err := broker.NewRabbitMQPublisher(conn, "Q1")
    // if err != nil {
    // 	log.Fatal("Error al crear el publicador de RabbitMQ:", err)
    // }

    // Por ahora, dejamos publisher en nil.

    // Inicializa el controlador del módulo ds18b20.
    ds18b20Controller := controllers.NewDS18B20Controller(ds18b20Service)

    // Registra las rutas del módulo ds18b20.
    routes.DS18B20Routes(r, ds18b20Controller)
}