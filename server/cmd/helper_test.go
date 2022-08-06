package cmd

import dataentity "simple_quiz/server/dataentity"

func setup() {
	userInfoData = append(userInfoData, dataentity.UserInfo{
		Username:            "User1",
		PercentageScore:     0,
		NumOfCorrectAnswers: 0,
		NumOfQuestions:      0,
	})
	userInfoData = append(userInfoData, dataentity.UserInfo{
		Username:            "User2",
		PercentageScore:     50,
		NumOfCorrectAnswers: 5,
		NumOfQuestions:      10,
	})
	userInfoData = append(userInfoData, dataentity.UserInfo{
		Username:            "User3",
		PercentageScore:     100,
		NumOfCorrectAnswers: 10,
		NumOfQuestions:      10,
	})
	userInfoData = append(userInfoData, dataentity.UserInfo{
		Username:            "User4",
		PercentageScore:     70,
		NumOfCorrectAnswers: 7,
		NumOfQuestions:      10,
	})
	userInfoData = append(userInfoData, dataentity.UserInfo{
		Username:            "User5",
		PercentageScore:     30,
		NumOfCorrectAnswers: 3,
		NumOfQuestions:      10,
	})

	questionData = append(questionData, dataentity.Question{
		Question: "Test Question 1",
		Choices: []dataentity.Choice{
			{
				ChoiceId: "1",
				Text:     "Choice 1",
			},
			{
				ChoiceId: "1",
				Text:     "Choice 2",
			},
		},
		CorrectId: "1",
	})

	questionData = append(questionData, dataentity.Question{
		Question: "Test Question 2",
		Choices: []dataentity.Choice{
			{
				ChoiceId: "1",
				Text:     "Choice 1",
			},
			{
				ChoiceId: "1",
				Text:     "Choice 2",
			},
		},
		CorrectId: "1",
	})
}

func teardown() {
	userInfoData = []dataentity.UserInfo{}
}
