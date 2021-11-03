package test

import (
	"testing"
	"time"

	"git.xenonstack.com/util/test-portal/config"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func BenchmarkSubmitTest(b *testing.B) {

	claims := map[string]interface{}{
		"drive":     "5",
		"email":     "harshit.agrawal@xenondigilabs.com",
		"test":      "cloud-engineer",
		"questions": "50",
		"startTime": "1568619010000",
	}
	for i := 0; i < b.N; i++ {
		SubmitTest(claims)

	}
}
func TestSubmitTest(t *testing.T) {

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

	time := float64(time.Now().Unix())
	questions := float64(50)
	//drive-id is not given
	claims := map[string]interface{}{
		// "drive": "5",
		"email": "harshitl@xenondigilabs.com",
		// "test":      "cloud-engineer",
		"questions": questions,
		"startTime": time,
	}
	_, err = SubmitTest(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed : ")
	}

	//if email is not given
	claims = map[string]interface{}{
		"drive": "5",
		// "email": "harshitl@xenondigilabs.com",
		// "test":      "cloud-engineer",
		"questions": questions,
		"startTime": time,
	}

	_, err = SubmitTest(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//if test-id is not given
	claims = map[string]interface{}{
		"drive": "5",
		"email": "harshitl@xenondigilabs.com",
		// "test":      "cloud-engineer",
		"questions": questions,
		// "startTime": time,
	}

	_, err = SubmitTest(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Log("test case Passed")
	}

	//if number of question is not given
	claims = map[string]interface{}{
		"drive": "5",
		"email": "harshitl@xenondigilabs.com",
		// "test":  "cloud-engineer",
		// "questions": "50",
		"startTime": time,
	}
	_, err = SubmitTest(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Logf("test case Passed")
	}

	//if starting time is not given
	claims = map[string]interface{}{
		"drive": "5",
		"email": "harshitl@xenondigilabs.com",
		// "test":      "cloud-engineer",
		"questions": questions,
		// "startTime": time,
	}
	_, err = SubmitTest(claims)
	if err == nil {
		t.Errorf("Expected Error but got nil")
	} else {
		t.Logf("test case Passed")
	}

	//if everything is  time is not given
	claims = map[string]interface{}{
		"drive": "5",
		"email": "harshitl@xenondigilabs.com",
		// "test":      "cloud-engineer",
		"questions": questions,
		"startTime": time,
	}
	_, err = SubmitTest(claims)
	if err == nil {
		t.Logf("Expected Error but got nil")
	} else {
		t.Logf("test case Passed")
	}

}
