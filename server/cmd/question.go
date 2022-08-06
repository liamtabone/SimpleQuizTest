package cmd

import (
	"context"
	"encoding/json"
	"io/ioutil"
	proto "simple_quiz/gen/proto"
	dataentity "simple_quiz/server/dataentity"

	"github.com/google/go-cmp/cmp"
)

// GetQuestions takes the user struct as a request and generates the next question for the user. The question is sent
// as a response in the Question structure
func (s *SimpleQuizServiceServer) GetQuestion(ctx context.Context, user *proto.User) (*proto.Question, error) {
	userInfo, err := getUserInfo(user.Username)
	if err != nil {
		return &proto.Question{QuestionId: -1}, err
	}

	question := createQuestionMsg(userInfo.NumOfQuestions)
	return &question, nil
}

func loadAllQuestions() {
	file, _ := ioutil.ReadFile(questionsJsonFileName)

	questionData = []dataentity.Question{}

	_ = json.Unmarshal([]byte(file), &questionData)
}

func getQuestionById(questionId int) dataentity.Question {
	if questionId < 0 {
		return dataentity.Question{}
	}

	if questionId < len(questionData) {
		return questionData[questionId]
	}

	return dataentity.Question{}
}

func createQuestionMsg(questionId int) proto.Question {
	question := getQuestionById(questionId)

	if cmp.Equal(question, dataentity.Question{}) {
		return proto.Question{QuestionId: -1}
	} else {
		return *question.GenerateQuestionMsgFromStruct()
	}
}
