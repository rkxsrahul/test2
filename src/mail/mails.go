package mail

import (
	"log"
	"os"
	"strings"

	"git.xenonstack.com/util/test-portal/config"
	"go.uber.org/zap"
)

func SendTestMail(claims map[string]interface{}) {

	//fetch user email from claims
	email, ok := claims["email"]
	zap.S().Info("email....", email)
	if !ok {
		return
	}
	//fetch user name from claims
	name, ok := claims["name"]
	zap.S().Info("name....", name)
	if !ok {
		return
	}
	//fetch test name from claims
	test, ok := claims["test_name"]
	zap.S().Info("test....", test)
	if !ok {
		return
	}

	// map saving name of user and verification code for email verification
	mapd := map[string]interface{}{
		"TestName": test,
		"Name":     name,
		"Url":      os.Getenv("HP_ACS_FRONT_ADDR"),
		"Website":  strings.TrimSuffix(strings.TrimPrefix(os.Getenv("HP_ACS_FRONT_ADDR"), "https://"), "/"),
	}
	log.Println(mapd)
	// saving subject as string
	sub := "XenonStack Career Portal Drive Notification"
	// saving template as string by parsing above map
	tmpl := EmailTemplate(config.MailPath+"/mail/templates/test.tmpl", mapd)

	//now sending mail
	SendMailV2(email.(string), sub, tmpl)
}
