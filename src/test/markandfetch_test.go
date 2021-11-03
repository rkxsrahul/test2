package test

import (
	"fmt"
	"log"
	"testing"

	"github.com/jinzhu/gorm"
	"go.uber.org/zap"

	"git.xenonstack.com/util/test-portal/config"
	"git.xenonstack.com/util/test-portal/src/bodytypes"
)

func BenchmarkMarkandfetch(b *testing.B) {
	fmt.Println("fmt", b.N)
	log.Println("log", b.N)
	b.StopTimer()
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

	var data bodytypes.New
	data.CurrentIndex = "3"
	data.CurrentPool = "cloud"
	data.Marked = "465210442541924353"
	data.NextIndex = "4"
	data.NextPool = "cloud"

	// claims := map[string]interface{}{
	// 	"drive": "5",
	// 	"email": "harshit.agrawal@xenondigilabs.com",
	// 	"test":  "cloud-engineer",
	// }

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		// Markandfetch(claims, data)

	}
}
func TestMarkandfetch(t *testing.T) {

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

	var data bodytypes.New
	data.CurrentIndex = "3"
	data.CurrentPool = "cloud"
	data.Marked = "465210442541924353"
	data.NextIndex = "4"
	data.NextPool = "cloud"

	//if drive-id is not given
	// claims := map[string]interface{}{
	// 	// "drive": "5",
	// 	"email": "harshit.agrawal@xenondigilabs.com",
	// 	"test":  "cloud-engineer",
	// }
	// _, err = Markandfetch(claims, data)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//if email-id  is not given
	// claims = map[string]interface{}{
	// 	"drive": "5",
	// 	// "email": "harshit.agrawal@xenondigilabs.com",
	// 	"test": "cloud-engineer",
	// }
	// _, err = Markandfetch(claims, data)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//How to give Right Answer
	data.CurrentIndex = "1"
	data.CurrentPool = "cloud"
	data.Marked = "465207395371057153"
	data.NextIndex = "2"
	data.NextPool = "cloud"
	//if email-id  is not given
	// claims = map[string]interface{}{
	// 	"drive": "14",
	// 	"email": "harshit.agrawal@xenondigilabs.com",
	// 	"test":  "cloud-engineer",
	// }
	// _, err = Markandfetch(claims, data)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//Testing for updateAnswerRoutine()
	// if CurrentIndex is zero
	data.CurrentIndex = "0"
	data.CurrentPool = "cloud"
	data.Marked = "46521044254192435"
	data.NextIndex = "60"
	data.NextPool = "cloud"
	//if email-id  is not given
	// claims = map[string]interface{}{
	// 	"drive": "5",
	// 	"email": "harshit.agrawal@xenondigilabs.com",
	// 	"test":  "cloud-engineer",
	// }
	// _, err = Markandfetch(claims, data)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//id data.NextPool is missing
	data.CurrentIndex = "4"
	data.CurrentPool = "cloud"
	data.Marked = "4652104425419243"
	data.NextIndex = "5"
	data.NextPool = ""
	//if email-id  is not given
	// claims = map[string]interface{}{
	// 	"drive": "5",
	// 	"email": "harshit.agrawal@xenondigilabs.com",
	// 	"test":  "cloud-engineer",
	// }
	// _, err = Markandfetch(claims, data)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//id nextindex is less then CurrentIndex is zero
	data.CurrentIndex = "1"
	data.CurrentPool = "cloud"
	data.Marked = "0"
	data.NextIndex = "2"
	data.NextPool = "cloud"
	//if email-id  is not given
	// claims = map[string]interface{}{
	// 	"drive": "15",
	// 	"email": "harshit.agrawal@xenondigilabs.com",
	// 	"test":  "cloud-engineer",
	// }
	// _, err = Markandfetch(claims, data)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//Marked is  zero
	data.CurrentIndex = "1"
	data.CurrentPool = "cloud"
	data.Marked = ""
	data.NextIndex = ""
	data.NextPool = "cloud"
	//if email-id  is not given
	// claims = map[string]interface{}{
	// 	"drive": "15",
	// 	"email": "harshit.agrawal@xenondigilabs.com",
	// 	"test":  "cloud-engineer",
	// }
	// _, err = Markandfetch(claims, data)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

}
