package infra

import (
	"bytes"
	"fmt"
	"log"
	"pantori/internal/domains/notifiers/core"
	"text/template"

	"gopkg.in/gomail.v2"
)

const emailTemplate = `
    <!DOCTYPE html>
    <html>
   <head>
    <title>Alerta de Vencimento</title>
    <style>
        table {
            width: 35%;
			margin: 20px auto;
            border-collapse: collapse;
        }
        th, td {
            border: 1px solid black;
            padding: 8px;
            text-align: left;
        }
        th {
            background-color: #f2f2f2;
        }
		h3 {
            text-align: center;
        }
        .container {
            padding-bottom: 20px;
        }
    </style>
</head>
    <body>
        <!-- Your exported HTML content here -->
        <div style="text-align: center;">
            <img src="cid:email_header.png" alt="Header Image" style="width: 100%; max-width: 600px;">
            <h2>Alerta de Vencimento</h2>
            <p>{{.Username}}, os seguintes produtos vão vencer ou já venceram:</p>

			{{ if .Expired }}
			<h3>Itens Vencidos :(( -> É melhor jogar logo fora</h3>
			<table>
				<tr>
					<th>Item</th>
					<th>Data de Validade</th>
				</tr>
				{{ range .Expired }}
				<tr>
					<td>{{ .Name }}</td>
					<td>{{ .Expire }}</td>
				</tr>
				{{ end }}
			</table>
			<hr>
			{{ end }}

			{{ if .ExpiresToday }}
			<h3>Itens que vencem hoje!! Nhamnham</h3>
			<table>
				<tr>
					<th>Item</th>
					<th>Data de Validade</th>
				</tr>
				{{ range .ExpiresToday }}
				<tr>
					<td>{{ .Name }}</td>
					<td>{{ .Expire }}</td>
				</tr>
				{{ end }}
			</table>
			<hr>
			{{ end }}

			{{ if .ExpiresSoon }}
			<h3>Itens que vencem em breve --- hora de planejar usar ou doar!!</h3>
			<table>
				<tr>
					<th>Item</th>
					<th>Data de Validade</th>
				</tr>
				{{ range .ExpiresSoon }}
				<tr>
					<td>{{ .Name }}</td>
					<td>{{ .Expire }}</td>
				</tr>
				{{ end }}
			</table>
			<hr>
			{{ end }}

            <p>Utilize ou descarte corretamente.</p>
            <p>Obrigado!</p>
        </div>
    </body>
    </html>
    `

type EmailAuth struct {
	Email    string
	Password string
}

type email struct {
	auth     EmailAuth
	smtpHost string
	smtpPort string
}

func NewEmailProvider(auth EmailAuth) *email {
	return &email{
		auth:     auth,
		smtpHost: "smtp.gmail.com",
		smtpPort: "587",
	}
}

func (e *email) SendEmail(user core.User, report core.Report) error {
	report.Username = user.Name

	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		log.Fatalf("Failed to parse template: %s", err)
	}

	var bodyContent string
	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, report)
	if err != nil {
		log.Fatalf("Failed to execute template: %s", err)
	}
	bodyContent = buffer.String()

	// Create a new email
	m := gomail.NewMessage()
	m.SetHeader("From", e.auth.Email)
	m.SetHeader("To", user.Email)
	m.SetHeader("Subject", "Alerta de Vencimento")
	m.SetBody("text/html", bodyContent)
	m.Embed("/go/bin/email_header.png")

	// SMTP server configuration
	d := gomail.NewDialer("smtp.gmail.com", 587, e.auth.Email, e.auth.Password)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		log.Fatalf("Failed to send email: %s", err)
	}
	fmt.Println("Email sent successfully!")
	return nil
}
