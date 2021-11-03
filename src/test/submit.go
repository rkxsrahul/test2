package test

import (
	"errors"
	"log"
	"time"

	"git.xenonstack.com/util/test-portal/config"
	"git.xenonstack.com/util/test-portal/src/bodytypes"
	"git.xenonstack.com/util/test-portal/src/mail"
	"go.uber.org/zap"
)

func SubmitTest(claims map[string]interface{}) (bodytypes.Overview, error) {
	// start span using parent span context
	now := time.Now().Unix()
	//create database connection
	db := config.DB
	log.Println(claims)
	//fetch total no. of questions from claims
	totalQuestions, ok := claims["questions"]
	log.Println(totalQuestions)
	if !ok {
		return bodytypes.Overview{}, errors.New("question is not set in token claim")
	}

	//fetch email from claims
	email, ok := claims["email"]
	if !ok {
		return bodytypes.Overview{}, errors.New("email si not set in token claim")
	}
	//fetch drive_id from claims
	driveId, ok := claims["drive"]
	if !ok {
		return bodytypes.Overview{}, errors.New("drive is not set in token claim")
	}
	//fetch drive_id from claims
	startTime, ok := claims["startTime"]
	if !ok {
		return bodytypes.Overview{}, errors.New("drive is not set in token claim")
	}
	expt := int64(startTime.(float64))
	//update data in user session table
	row := db.Exec("update user_sessions set expire=?, time_taken=? where email=? AND drive_id=? AND expire>=?", now, now-expt, email, driveId, now).RowsAffected
	zap.S().Info("rows updated.....", row)
	if row != 0 {
		// send mail for notifying your test had been submitted
		go mail.SendTestMail(claims)
	}

	//sql query for answered
	var answered int
	db.Raw("select COUNT(email) from answers where email=? AND drive_id=? AND marked_id>0", email, driveId).Count(&answered)

	//return the overview of the test
	return bodytypes.Overview{
		TimeTaken:  now - expt,
		Total:      int(totalQuestions.(float64)),
		Attempted:  answered,
		Unanswered: int(totalQuestions.(float64)) - answered,
	}, nil
}
