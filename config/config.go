package config

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	Port   string = os.Getenv("TEST_PORTAL_PORT")
	DbUser string = os.Getenv("DRIVE_PORTAL_DB_USER")
	DbPass string = os.Getenv("DRIVE_PORTAL_DB_PASS")
	DbHost string = os.Getenv("DRIVE_PORTAL_DB_HOST")
	DbPort string = os.Getenv("DRIVE_PORTAL_DB_PORT")
	DbType string = os.Getenv("DRIVE_PORTAL_DB_TYPE")
	DbName string = os.Getenv("DRIVE_PORTAL_DB_NAME")
	//common private key between test-portal and drive-portal
	TestPortalKey string = os.Getenv("TEST_PORTAL_PRIVATE_KEY")

	DisableMail string = os.Getenv("MAIL_DISABLE")

	// mail templates and images path
	MailPath string = os.Getenv("MAIL_TEMPLATES_PATH")
	DB       *gorm.DB
)

const (
	DBConnSSL string = "disable"
	//JWT Timeout in seconds
	JWTExpireTime time.Duration = time.Hour * 24
)

func init() {
	if MailPath == "" {
		MailPath = "./src"
	}
}

func DBConfig() string {
	if DbType == "postgres" {
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", DbHost, DbPort, DbUser, DbPass, DbName, DBConnSSL)
	}
	// creating db connection string
	str := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		DbUser,
		DbPass,
		DbHost,
		DbPort,
		DbName, DBConnSSL)

	return str
}
