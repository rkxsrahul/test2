package health

import (
	"git.xenonstack.com/util/test-portal/config"
)

func Healthz() error {
	// connecting to db
	db := config.DB
	err := db.Exec("SELECT 1").Error
	return err
}
