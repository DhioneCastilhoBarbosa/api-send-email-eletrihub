package main

import (
	"github.com/DHIONECASTILHOBARBOSA/email-api/internal/config"
	"github.com/DHIONECASTILHOBARBOSA/email-api/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	r := gin.Default()
	r.Use(cors.Default()) // Habilita CORS para todos os domínios

	r.POST("/send-email", handler.SendEmailHandler)

	r.Run(":8080")
}
