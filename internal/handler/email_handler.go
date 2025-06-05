package handler

import (
	"net/http"

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
