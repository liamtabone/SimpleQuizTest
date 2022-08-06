package dataentity

type UserInfo struct {
	Username            string  `json:"username"`
	PercentageScore     float32 `json:"percentageScore"`
	NumOfCorrectAnswers int     `json:"numOfCorrectAnswers"`
	NumOfQuestions      int     `json:"numOfQuestions"`
}
