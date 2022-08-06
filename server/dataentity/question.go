package dataentity

import "simple_quiz/gen/proto"

type Choice struct {
	ChoiceId string `json:"choiceId"`
	Text     string `json:"text"`
}

type Question struct {
	Question  string   `json:"question"`
	Choices   []Choice `json:"choices"`
	CorrectId string   `json:"correctId"`
}

func (questionStruct Question) GenerateQuestionMsgFromStruct() *proto.Question {
	var choices []*proto.Choice

	for i := 0; i < len(questionStruct.Choices); i++ {
		choices = append(choices, &proto.Choice{
			ChoiceId: questionStruct.Choices[i].ChoiceId,
			Text:     questionStruct.Choices[i].Text,
		})
	}

	return &proto.Question{
		Text:    questionStruct.Question,
		Choices: choices,
	}
}
