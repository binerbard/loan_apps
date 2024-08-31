package utilsmail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"strconv"
	"text/template"

	"loan_apps/config/setting"
)

func MailSender(){
	// Konfigurasi SMTP server
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpUsername := "sekaaranemia@gmail.com"
	smtpPassword := "xynhcnrxudysjzgv"

	// Pengaturan email pengirim
	from := "email@sekaar.com"
	fromName := "Sekar"

	// Pengaturan email penerima
	to := []string{"muhammadfahmyjob@gmail.com"}


	// Template Email
	t, _ := template.ParseFiles("assets/template/email.html")
	var message bytes.Buffer
	t.Execute(&message, struct {
	Name    string
	Message string
	}{
	Name:    "Puneet Singh",	
	Message: "This is a test message in a HTML template",
	})
	

	// Subject dan isi pesan
	subject := "Custom Subject Here"

	// Mengatur header email
	headers := make(map[string]string)
	headers["From"] = fromName + " <" + from + ">"
	headers["To"] = to[0] // Untuk multiple recipients, tambahkan semicolon (;) diantara alamat email
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""
	headers["Content-Transfer-Encoding"] = "base64"

	// Mengonversi header email ke format MIME
	var emailHeaders string
	for key, value := range headers {
		emailHeaders += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	emailContent := []byte(emailHeaders + "\r\n" + message.String())
	// Mengirim email menggunakan SMTP
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, emailContent)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}


func SendMail(assets string, data any, email string, subject string) {

	// Konfigurasi SMTP server
	smtpHost := setting.SMTPHost
	smtpPort := setting.SMTPPort
	smtpUsername := setting.SMTPEmail
	smtpPassword := setting.SMTPPassword

	// Pengaturan email pengirim
	from := setting.SMTPEmailSender
	fromName := setting.SMTPNameSender

	// Pengaturan email penerima
	to := []string{email}


	// Template Email
	t, _ := template.ParseFiles(assets)
	var message bytes.Buffer
	t.Execute(&message, data)
	

	// Mengatur header email
	headers := make(map[string]string)
	headers["From"] = fromName + " <" + from + ">"
	headers["To"] = to[0] // Untuk multiple recipients, tambahkan semicolon (;) diantara alamat email
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"utf-8\""
	headers["Content-Transfer-Encoding"] = "base64"

	// Mengonversi header email ke format MIME
	var emailHeaders string
	for key, value := range headers {
		emailHeaders += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	emailContent := []byte(emailHeaders + "\r\n" + message.String())
	// Mengirim email menggunakan SMTP
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)
	err := smtp.SendMail(smtpHost + ":" + strconv.Itoa(smtpPort), auth, from, to, emailContent)
	if err != nil {
		fmt.Println("Failed to send email:", err)
		return
	}

	fmt.Println("Email sent successfully!")

}