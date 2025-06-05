# 📧 Email API - Go + SMTP

API simples e robusta escrita em Go que permite o envio de e-mails via SMTP com TLS, utilizando a infraestrutura da Hostinger. A API aceita dados de um formulário (nome, sobrenome, e-mail, telefone, assunto e mensagem) e envia para um e-mail definido, com template HTML estilizado em **fundo preto e letras brancas**, incluindo sua logo no topo.

---

## 🚀 Tecnologias utilizadas

- [Go (Golang)](https://golang.org)
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [GoDotEnv](https://github.com/joho/godotenv) - Carregamento de variáveis de ambiente
- [SMTP](https://pkg.go.dev/net/smtp) - Envio de e-mails
- [CORS Middleware](https://github.com/gin-contrib/cors) - Suporte a requisições frontend

---

## 📁 Estrutura do Projeto

```
email-api/
├── cmd/
│   └── server/
│       └── main.go           # Inicialização do servidor
├── internal/
│   ├── config/
│   │   └── config.go         # Carrega variáveis do .env
│   ├── handler/
│   │   └── email_handler.go  # Controlador da rota /send-email
│   ├── mail/
│   │   └── smtp.go           # Lógica de envio de e-mail via SMTP
│   └── model/
│       └── email.go          # Estrutura do payload recebido
├── .env                      # Variáveis sensíveis (NÃO versionar)
├── go.mod
└── README.md
```

---

## ⚙️ Como rodar o projeto

### Pré-requisitos

- Go 1.18+
- Conta de e-mail com SMTP ativo (ex: Hostinger)

### 1. Clone o projeto

```bash
git clone https://github.com/SeuUsuario/email-api.git
cd email-api
```

### 2. Configure o `.env`

Crie um arquivo `.env` com:

```
SMTP_HOST=smtp.hostinger.com
SMTP_PORT=587
SMTP_USER=seu@email.com
SMTP_PASS=sua_senha
RECEIVER_EMAIL=seu@email.com
```

### 3. Instale as dependências

```bash
go mod tidy
```

### 4. Rode a API

```bash
go run cmd/server/main.go
```

A API estará disponível em `http://localhost:8080`.

---

## 📬 Exemplo de Requisição

### Endpoint: `POST /send-email`

```json
{
  "first_name": "João",
  "last_name": "Silva",
  "email": "joao@email.com",
  "phone": "11999999999",
  "subject": "Contato",
  "message": "Gostaria de saber mais sobre o serviço."
}
```

---

## ✅ Resultado

- O e-mail é enviado com fundo preto, letras brancas e sua logo no topo (via URL do GitHub)
- O conteúdo é entregue no e-mail definido em `RECEIVER_EMAIL`

---

## 🔒 Segurança

- Nunca versionar seu `.env`
- Usar App Passwords se possível para SMTP (ex: Hostinger, Gmail, etc.)

---

## 👨‍💻 Autor

Dhione Castilho Barbosa  


## 📄 Licença

MIT License