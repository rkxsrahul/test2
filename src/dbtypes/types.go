package dbtypes

import (
	"time"
)

// question structure
type Questions struct {
	Id       int
	Title    string
	Type     string
	PoolId   string
	AnswerId int
	ImageUrl string
	Options  []Options
}

//options structure to store options values
type Options struct {
	Id        string
	Type      string
	Value     string
	ImageUrl  string
	IsCorrect bool
	QuesId    int
}

// test structure to store test pools
type TestPool struct {
	Id            int
	PoolId        string
	TestId        string
	NoOfQuestions int
}

// usersession structure to store user session token
type UserSession struct {
	Email   string
	DriveId int
	Token   string
	Expire  int64
	Browser int
}

// answer structure to store submitted answers and assigned questions
type Answers struct {
	Email     string
	DriveId   int
	PoolId    string
	QuesId    int
	MarkedId  int
	AnswerId  int
	QuesIndex int
	Time      int64
	CreatedAt time.Time
	UpdatedAt time.Time
}
