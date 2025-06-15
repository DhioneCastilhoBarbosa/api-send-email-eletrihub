package mail

import (
	"fmt"
	"os"

	"strconv"

	"github.com/DHIONECASTILHOBARBOSA/email-api/internal/model"
	"gopkg.in/gomail.v2"
)

func SendEmail(req model.EmailRequest) error {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", os.Getenv("RECEIVER_EMAIL"))
	m.SetHeader("Subject", fmt.Sprintf("Motivo: %s", req.Subject))

	htmlBody := fmt.Sprintf(`
		<html>
		<body style="background-color:black;color:white;padding:20px;font-family:sans-serif;">
			<div style="text-align:left;margin-bottom:20px;">
				<img src="https://raw.githubusercontent.com/DhioneCastilhoBarbosa/hubinscar/dbae7b6ec92e75430765ddac74ec034f8714adde/src/assets/LOGO%%20BRANCA.png" alt="Logo" style="max-width:100px;">
			</div>
			<div style="color:black;background-color:white;padding:20px;border-radius:10px;">
				<h2>Recebemos uma nova solicitação de contato.</h2>
				<p><strong>Nome:</strong> %s %s</p>
				<p><strong>Email:</strong> %s</p>
				<p><strong>Telefone:</strong> %s</p>
				<p><strong>Mensagem:</strong><br/>%s</p>
			</div>
		</body>
		</html>
	`, req.FirstName, req.LastName, req.Email, req.Phone, req.Message)

	m.SetBody("text/html", htmlBody)
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %v", err)
	}
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port, // Porta SSL
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	d.SSL = false // Ativa SSL direto

	return d.DialAndSend(m)
}

func SendEmailSupport(req model.EmailRequest) error {
	m := gomail.NewMessage()

	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", os.Getenv("RECEIVER_EMAIL"))
	m.SetHeader("Subject", fmt.Sprintf("Motivo: %s", req.Subject))

	htmlBody := fmt.Sprintf(`
		<html>
		<body style="background-color:black;color:white;padding:20px;font-family:sans-serif;">
			<div style="text-align:left;margin-bottom:20px;margin-left:10px;">
				<img src="https://raw.githubusercontent.com/DhioneCastilhoBarbosa/hubinscar/dbae7b6ec92e75430765ddac74ec034f8714adde/src/assets/LOGO%%20BRANCA.png" alt="Logo" style="max-width:100px;">
			</div>
			<div style="color:black;background-color:white;padding:20px;border-radius:10px;">
				<h2>Recebemos uma nova solicitação de suporte.</h2>
				<p><strong>Nome:</strong> %s %s</p>
				<p><strong>Email:</strong> %s</p>
				<p><strong>Telefone:</strong> %s</p>
				<p><strong>Mensagem:</strong><br/>%s</p>
			</div>
		</body>
		</html>
	`, req.FirstName, req.LastName, req.Email, req.Phone, req.Message)

	m.SetBody("text/html", htmlBody)
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return fmt.Errorf("invalid SMTP_PORT: %v", err)
	}
	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port, // Porta SSL
		os.Getenv("SMTP_USER"),
		os.Getenv("SMTP_PASS"),
	)

	d.SSL = false // Ativa SSL direto

	return d.DialAndSend(m)
}
