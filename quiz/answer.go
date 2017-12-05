package quiz

//Answer -- struct to hold an answer
type Answer struct {
	Answer string
	Correct	bool
}

//UserAnswer -- struct to hold individual answers made by the user
type UserAnswer struct {
	Answer	string
	Question *Question
	UserAnswer	*Answer
}
