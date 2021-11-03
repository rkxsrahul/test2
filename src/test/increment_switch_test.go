package test

import (
	"testing"

	"git.xenonstack.com/util/test-portal/config"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func BenchmarkBrowserSwitch(b *testing.B) {

	db, err := gorm.Open("postgres", config.DBConfig())
	if err != nil {
		zap.S().Error(err)
		return
	}
	// close db instance whenever whole work completed
	defer db.Close()
	db.DB().SetMaxIdleConns(50)
	// db.DB().SetConnMaxLifetime(1 * time.Hour)
	config.DB = db

	claims := map[string]interface{}{
		"drive": "5",
		"email": "harshit.agrawal@xenondigilabs.com",
	}
	for i := 0; i < b.N; i++ {
		BrowserSwitch(claims)
	}
}

func TestBrowserSwitch(t *testing.T) {
	db, err := gorm.Open("postgres", config.DBConfig())
	if err != nil {
		zap.S().Error(err)
		return
	}
	// close db instance whenever whole work completed
	defer db.Close()
	db.DB().SetMaxIdleConns(50)
	// db.DB().SetConnMaxLifetime(1 * time.Hour)
	config.DB = db

	// //if drive-id is not given
	claims := map[string]interface{}{
		// "drive": "5",
		"email": "harshit.agrawal@xenondigilabs.com",
	}

	_, err = BrowserSwitch(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//if email-id is not given
	claims = map[string]interface{}{
		"drive": "5",
		// "email": "harshit.agrawal@xenondigilabs.com",
	}

	_, err = BrowserSwitch(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//if everything is given
	claims = map[string]interface{}{
		"drive": "5",
		"email": "harshit.agrawal@xenondigilabs.com",
	}

	_, err = BrowserSwitch(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

}
