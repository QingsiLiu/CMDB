package testmail

import (
	"gopkg.in/gomail.v2"
)

func main() {

	smtpAddr := "smtp.qq.com"
	smtpPort := 465
	smtpUser := "1029806879@qq.com"
	smtpPassword := "calltvvvpfuxbfje"

	from := "1029806879@qq.com"
	tos := []string{"785154048@qq.com", "836549547@qq.com", "1194789138@qq.com"}
	subject := "测试找你们收邮件啦"
	content := `<span style="color: red; font-size: 20px">测试测试测试一下</span>`

	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", tos...)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(smtpAddr, smtpPort, smtpUser, smtpPassword)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
