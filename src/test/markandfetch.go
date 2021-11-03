package test

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"git.xenonstack.com/util/test-portal/config"
	"git.xenonstack.com/util/test-portal/src/bodytypes"
	"git.xenonstack.com/util/test-portal/src/dbtypes"
	"go.uber.org/zap"
)

type Response struct {
	Error    error
	Question bodytypes.Question
}

//Checking index and pool
func checkIndex(index string, test string, pool string) error {
	indexInt, err := strconv.Atoi(index)
	if err != nil {
		indexInt = 0
	}
	//connecting the database
	db := config.DB
	// check index is less then zero or not
	if indexInt <= 0 {
		return errors.New("please enter a valid index")
	}

	// fetch total no. of questions from question table related to that question id
	pools := []dbtypes.TestPool{}

	db.Raw("select no_of_questions from test_pools where test_id=? and pool_id=?", test, pool).Scan(&pools)
	if len(pools) == 0 {
		return errors.New("No question are there in this pool")
	}

	//check index is less then total number of questions
	if pools[0].NoOfQuestions < indexInt {
		return errors.New("Please enter valid index (less then or equal to " + strconv.Itoa(pools[0].NoOfQuestions) + " )")
	}
	return nil
}

// Markandfetch Questions
func Markandfetch(claims map[string]interface{}, data bodytypes.New, ch chan Response) {

	res := Response{}
	//Fetch email from claims
	email, ok := claims["email"]
	if !ok {
		res.Error = errors.New("email is not set in token claims")
		res.Question = bodytypes.Question{}
		ch <- res
		return
	}
	//Fetch drive_id from claims
	driveId, ok := claims["drive"]
	if !ok {
		res.Error = errors.New("drive is not set in token claims")
		res.Question = bodytypes.Question{}
		ch <- res
		return
	}

	// go routine for update
	updateAnswerRoutine(email.(string), driveId.(string), data.CurrentIndex, data.CurrentPool, data.Marked, &res)
	if res.Error != nil {
		ch <- res
		return
	}

	// go routine for question fetch
	fetchQuestionRoutine(email.(string), driveId.(string), data.NextIndex, data.NextPool, claims["test"].(string), &res)
	ch <- res
	return
}

//Channel for Update Question
func updateAnswerRoutine(email, driveID, index, pool, marked string, ch *Response) {
	//connecting to the database
	db := config.DB
	if marked == "0" {
		ch.Error = nil
		ch.Question = bodytypes.Question{}
		return
	}
	//Marked Answer
	row := db.Exec("update answers set marked_id= ?, updated_at=? where drive_id = ? and email =?  and pool_id = ? and ques_index = ? ", marked, time.Now(), driveID, email, pool, index).RowsAffected
	if row == 0 {
		zap.S().Error("Answer is not Marked")
		ch.Error = errors.New("Answer is not Marked")
		ch.Question = bodytypes.Question{}
		return
	}
	ch.Error = nil
	ch.Question = bodytypes.Question{}
	return
}

//Channel for Fetch Question
func fetchQuestionRoutine(email, driveId string, index, pool, test string, ch *Response) {
	drive, _ := strconv.Atoi(driveId)
	//connecting to the database
	db := config.DB
	if pool == "" || index == "" {

		ch.Question = bodytypes.Question{}
		ch.Error = nil
		return
	}
	//Fetched Answer
	err := checkIndex(index, test, pool)
	if err != nil {
		ch.Question = bodytypes.Question{}
		ch.Error = err
		return
	}
	answers := []dbtypes.Answers{}
	db.Raw("select ques_id from answers where drive_id = ? and email=? and  pool_id =? and ques_index=? ", driveId, email, pool, index).Scan(&answers)
	if len(answers) != 0 {
		// fetch question by id
		question, err := fetchQuestionByID(answers[0].QuesId)
		if err != nil {
			ch.Question = bodytypes.Question{}
			ch.Error = err
			return
		}
		ch.Question = question
		ch.Error = err
		return
	}
	result, err := fetchQuestion(email, pool, index, drive)
	ch.Question = result
	ch.Error = err
	return
}

// fetchQuestionByID Fetch Question By ID
func fetchQuestionByID(quesID int) (bodytypes.Question, error) {
	//connecting the database
	db := config.DB.DB()
	result := bodytypes.Question{}
	row := db.QueryRow("SELECT q.title,q.type,q.image_url, json_agg(o.*) as options FROM questions q JOIN (SELECT cast (o.id as varchar),o.ques_id,o.type,o.value,o.image_url from options o) as o ON q.id = o.ques_id where q.id=" + strconv.Itoa(quesID) + " GROUP BY q.id,q.type,q.title,q.image_url")
	var title string
	var qtype string
	var image string
	var options json.RawMessage
	err := row.Scan(&title, &qtype, &image, &options)
	if err != nil {
		zap.S().Error(err)
		return bodytypes.Question{}, err
	}
	var v []bodytypes.Options
	err = json.Unmarshal(options, &v)
	if err != nil {
		zap.S().Error(err)
		return bodytypes.Question{}, err
	}
	result.ID = quesID
	result.Title = title
	result.Type = qtype
	result.ImageUrl = image
	result.Options = v
	return result, nil
}

//Fetch Questions
func fetchQuestion(email, pool, index string, drive int) (bodytypes.Question, error) {
	//connecting to the database
	db := config.DB
	//fetch question
	var ques []dbtypes.Questions
	db.Exec("select id from questions where id not in (select ques_id from answers where drive_id=? AND email=? AND pool_id=?) LIMIT 1", drive, email, pool).Find(&ques)

	if len(ques) == 0 {
		zap.S().Error("All questions had been assigned")
		return bodytypes.Question{}, errors.New("all questions had been assigned")
	}
	// fetch question by id
	result, err := fetchQuestionByID(ques[0].Id)
	if err != nil {
		return bodytypes.Question{}, err
	}

	var count int
	db.Raw("select count(email) from answers where drive_id=? and email=? and pool_id=? and ques_index = ?", drive, email, pool, index).Count(&count)
	if count == 0 {

		indexInt, _ := strconv.Atoi(index)
		//assign question in db
		db.Create(&dbtypes.Answers{
			Email:     email,
			DriveId:   drive,
			PoolId:    pool,
			QuesId:    ques[0].Id,
			AnswerId:  ques[0].AnswerId,
			QuesIndex: indexInt,
		})
	}

	return result, nil
}
