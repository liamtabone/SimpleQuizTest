package cmd

import (
	"context"
	proto "simple_quiz/gen/proto"
)

// Get question id and the client's answer as a request, check if the answer is correct and update the UserInfo.Json file
// with the number of questions answered by the user and the updated score
func (s *SimpleQuizServiceServer) GetAnswer(ctx context.Context, userAnswer *proto.Answer) (*proto.Empty, error) {
	submitAnswer(userAnswer.Username, int(userAnswer.QuestionId), userAnswer.AnswerId)
	return &proto.Empty{}, nil
}

func submitAnswer(username string, questionId int, answerId string) error {
	userInfo, err := getUserInfo(username)
	if err != nil {
		return err
	}

	if answerId == getQuestionById(userInfo.NumOfQuestions).CorrectId {
		userInfo.NumOfCorrectAnswers++
	}
	userInfo.NumOfQuestions++
	userInfo.PercentageScore, err = getPercentageScore(userInfo.NumOfCorrectAnswers, userInfo.NumOfQuestions)

	if err != nil {
		return err
	}

	writeUserInfoToJson(userInfoJsonFileName)

	return nil
}
