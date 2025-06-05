package mail

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"

	"github.com/DHIONECASTILHOBARBOSA/email-api/internal/model"
)

func SendEmail(req model.EmailRequest) error {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASS"), os.Getenv("SMTP_HOST"))

	to := []string{os.Getenv("RECEIVER_EMAIL")}
	from := os.Getenv("SMTP_USER")
	subject := fmt.Sprintf("Contato: %s", req.Subject)

	body := fmt.Sprintf(`
	<html>
	<body style="background-color:black;color:white;padding:20px;font-family:sans-serif;">
		<div style="text-align:center;margin-bottom:20px;">
			<img src="https://raw.githubusercontent.com/DhioneCastilhoBarbosa/hubinscar/dbae7b6ec92e75430765ddac74ec034f8714adde/src/assets/LOGO%%20BRANCA.png" alt="Logo" style="max-width:200px;">
		</div>
		<h2 style="color:white;">Nova mensagem do site</h2>
		<p><strong>Nome:</strong> %s %s</p>
		<p><strong>Email:</strong> %s</p>
		<p><strong>Telefone:</strong> %s</p>
		<p><strong>Mensagem:</strong><br/>%s</p>
	</body>
	</html>
`, req.FirstName, req.LastName, req.Email, req.Phone, req.Message)

	msg := strings.Join([]string{
		"From: " + from,
		"To: " + os.Getenv("RECEIVER_EMAIL"),
		"Subject: " + subject,
		"MIME-Version: 1.0",
		"Content-Type: text/html; charset=\"UTF-8\"",
		"",
		body,
	}, "\r\n")

	addr := fmt.Sprintf("%s:%s", os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))
	return smtp.SendMail(addr, auth, from, to, []byte(msg))
}
