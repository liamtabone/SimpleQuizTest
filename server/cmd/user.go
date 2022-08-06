package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	proto "simple_quiz/gen/proto"
	"simple_quiz/server/dataentity"

	"github.com/google/go-cmp/cmp"
)

// NewUser creates new user from the request, initializes all fields and saves the user in the JSON file
func (s *SimpleQuizServiceServer) NewUser(ctx context.Context, user *proto.User) (*proto.Empty, error) {
	err := createNewUser(user.Username)
	return &proto.Empty{}, err
}

func getUserInfo(username string) (*dataentity.UserInfo, error) {
	if username == "" {
		return &dataentity.UserInfo{}, errors.New("cannot get user info for empty username")
	}

	for i := 0; i < len(userInfoData); i++ {
		if userInfoData[i].Username == username {
			return &userInfoData[i], nil
		}
	}

	return &dataentity.UserInfo{}, errors.New("user '" + username + "' does not exist")
}

func initUser(username string) dataentity.UserInfo {
	return dataentity.UserInfo{
		Username:            username,
		PercentageScore:     0,
		NumOfCorrectAnswers: 0,
		NumOfQuestions:      0,
	}
}

// Create a new user based on parameter username, return an error if the user already exists
func createNewUser(username string) error {
	userInfo, _ := getUserInfo(username)
	if (!cmp.Equal(userInfo, &dataentity.UserInfo{})) {
		return errors.New("user '" + username + "' already exists")
	}

	userInfoData = append(userInfoData, initUser(username))
	err := writeUserInfoToJson(userInfoJsonFileName)

	if err != nil {
		return err
	}

	return nil
}

func writeUserInfoToJson(fileName string) error {
	file, err := json.MarshalIndent(userInfoData, "", " ")

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, file, 0644)

	if err != nil {
		return err
	}

	return nil
}
