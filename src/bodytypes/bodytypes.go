package bodytypes

// New Request for NewQuestion
type New struct {
	CurrentIndex string `json:"index"`
	CurrentPool  string `json:"pool"`
	Marked       string `json:"marked"`
	NextPool     string `json:"next_pool"`
	NextIndex    string `json:"next_index"`
}

// Question details structure
type Question struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Type     string    `json:"type"`
	Options  []Options `json:"options"`
	ImageUrl string    `json:"image_url"`
}

// Options stucture details of the Question
type Options struct {
	ID        string `json:"id"`
	Value     string `json:"value"`
	Type      string `json:"type"`
	ImageUrl  string `json:"image_url"`
	IsCorrect bool   `json:"-"`
}

// Overview is a strcuture defining test overview after test completion
type Overview struct {
	Attempted  int   `json:"attempted"`
	Unanswered int   `json:"unanswered"`
	Total      int   `json:"total"`
	TimeTaken  int64 `json:"timetaken"`
}
