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
	r.POST("/send-email-support", handler.SendEmailSupportHandler)
	r.POST("/send-email-new-installer", handler.PostInstallerHandler)
	r.POST("/withdrawal-request", handler.PostWithdrawRequestHandler)
	r.Run(":8093")
}
