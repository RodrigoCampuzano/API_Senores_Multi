package main

import (
	"APIs/src/core/middleware"
	infrastructure "APIs/src/ds18b20"
	infraestructure"APIs/src/max30102/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.MiddlewareCORS())
	infraestructure.Init(r)
	infrastructure.InitTemperatura(r)
	if err := r.Run(":8081"); err != nil {
		panic(err)
	}

}
