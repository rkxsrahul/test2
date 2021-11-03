package test

import (
	"errors"
	"net/http"
	"strconv"

	"git.xenonstack.com/util/test-portal/config"
	"go.uber.org/zap"
)

// BrowserSwitch is a function
func BrowserSwitch(claims map[string]interface{}) (int, error) {
	db := config.DB

	//fetch email from claims
	email, ok := claims["email"]
	zap.S().Info("email....", email)
	if !ok {
		return http.StatusInternalServerError, errors.New("email is not set in token claims")
	}
	//fetch drive_id from claims
	drive, ok := claims["drive"]
	if !ok {
		return http.StatusInternalServerError, errors.New("drive is not set in token claims")
	}
	driveID, err := strconv.Atoi(drive.(string))
	if err != nil {
		return http.StatusInternalServerError, errors.New("please set valid drive id only int")
	}
	zap.S().Info("drive....", driveID)

	//Updating browser count by 1
	rows := db.Exec("update user_sessions set browser = browser + 1 where drive_id=? AND email=?", driveID, email.(string)).RowsAffected
	zap.S().Info("rows updated.....", rows)
	return http.StatusAccepted, nil
}
