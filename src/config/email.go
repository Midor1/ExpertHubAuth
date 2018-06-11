package config

import (
	"strings"
	"net/smtp"
)

func SendMail(to, username, subject, emailContent string) error {
	user := C.User.EmailAddress
	password := C.User.EmailPassword
	host := C.User.EmailHost
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	body := "<table align='center' border='0' cellpadding='0' cellspacing='0'style=' width:88%; margin-top:20px; margin-bottom:20px; background:#fafafa; border:1px solid #ddd;'><tbody>" +
		"<tr>" +
		"<td width='24'>&nbsp;</td>" +
		"</tr><tr>" +
		"<td width='24'>&nbsp;</td>" +
		"<td colspan='2' style='color:#858585; font-family:Arial, Helvetica, sans-serif; font-size:14px; line-height:20px; padding-top:24px;'>" + username + " 你好：</td>" +
		"<td width='24'>&nbsp;</td>" +
		"</tr><tr>" +
		"<td width='24'>&nbsp;</td>" +
		"<td colspan='2' style='color:#858585; font-family:Arial, Helvetica, sans-serif; font-size:14px; line-height:20px; padding-top:18px;'>" + emailContent + "</td>" +
		"<td width='24'>&nbsp;</td>" +
		"</tr>" +
		"<tr>" +
		"<td style='padding-top:18px; padding-bottom:32px; border-bottom:1px solid #e1e1e1;' width='24'>&nbsp;</td>" +
		"<td colspan='2' style='color:#858585; font-family:Arial, Helvetica, sans-serif; font-size:14px; line-height:20px; padding-top:18px; padding-bottom:32px; border-bottom:1px solid #e1e1e1;'>谢谢使用！<br> </td>" +
		"<td style='padding-top:18px; padding-bottom:32px; border-bottom:1px solid #e1e1e1;' width='24'>&nbsp;</td>" +
		"</tr>" +
		"</tbody></table>"
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: [ExpertHub邮件OA]" + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	sendTo = append(sendTo, user)
	err := smtp.SendMail(host, auth, user, sendTo, msg)
	return err
}
