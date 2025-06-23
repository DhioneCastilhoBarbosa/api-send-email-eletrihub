package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/DHIONECASTILHOBARBOSA/email-api/internal/mail"
	"github.com/DHIONECASTILHOBARBOSA/email-api/internal/model"
	"github.com/gin-gonic/gin"
)

func SendEmailHandler(c *gin.Context) {
	var req model.EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := mail.SendEmail(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar e-mail: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "E-mail enviado com sucesso!"})
}

func SendEmailSupportHandler(c *gin.Context) {
	var req model.EmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	if err := mail.SendEmailSupport(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao enviar e-mail: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "E-mail de suporte enviado com sucesso!"})
}

func PostInstallerHandler(c *gin.Context) {
	var data model.InstallerData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(400, gin.H{"error": "Dados inválidos"})
		return
	}

	// Envia o e-mail
	if err := mail.SendEmailNewInstaller(data); err != nil {
		c.JSON(500, gin.H{"error": fmt.Sprintf("Erro ao enviar email: %v", err)})
		return
	}

	c.JSON(200, gin.H{"message": "Cadastro recebido com sucesso!"})
}

func PostWithdrawRequestHandler(c *gin.Context) {
	var data model.WithdrawMoneyRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	if os.Getenv("RECEIVER_EMAIL") == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Email de destino não configurado"})
		return
	}

	if err := mail.SendEmailWithdrawRequest(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Erro ao enviar e-mail: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Solicitação de saque enviada com sucesso!"})
}
