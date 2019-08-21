package utils

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"

	"bx.com/user-service/config"
	"github.com/sirupsen/logrus"
)

func GetEmailConfig() config.EmailSource {
	cfg := config.Parse("../config/app.yaml")
	return cfg.EmailSource
}

func SendToMail(to, subject, body, mailtype string) error {
	config := GetEmailConfig()
	hp := strings.Split(config.Host, ":")
	auth := smtp.PlainAuth("", config.User, config.Password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + config.User + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(config.Host, auth, config.User, send_to, msg)
	if err != nil {
		logrus.Infof("Email sent got error: %s", err.Error())
		return err
	}
	return nil
}

func ReadTemplate(fileName string) (string, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func RegisterEmail(baseUrl, email, code string) error {
	subject := "New Account Register Activation Email"

	body, err := ReadTemplate("../utils/regist_email.htm")
	if err != nil {
		return err
	}
	body = strings.Replace(body, "$User", email, -1)

	activateUrl := fmt.Sprintf("%s%s", baseUrl, code)
	body = strings.Replace(body, "$ActivateLink", activateUrl, -1)

	err = SendToMail(email, subject, body, "html")
	if err != nil {
		logrus.Errorf(err.Error())
		return err
	} else {
		logrus.Infof("Eail sent to %s success!", email)
		return nil
	}
}

func VerifyCodeEmail(email, user, code string) error {
	subject := "New Verification Code Email"

	body, err := ReadTemplate("../utils/resetpwd_email.htm")
	if err != nil {
		return err
	}
	body = strings.Replace(body, "$User", user, -1)
	body = strings.Replace(body, "$VerifyCode", code, -1)

	err = SendToMail(email, subject, body, "html")
	if err != nil {
		logrus.Errorf(err.Error())
		return err
	} else {
		logrus.Infof("Eail sent to %s success!", email)
		return nil
	}
}
