package cmd

import (
	dataentity "simple_quiz/server/dataentity"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInitUser(t *testing.T) {

	input := initUser("User1")
	expected := dataentity.UserInfo{
		Username:            "User1",
		PercentageScore:     0,
		NumOfCorrectAnswers: 0,
		NumOfQuestions:      0,
	}

	if !cmp.Equal(input, expected) {
		t.Errorf("User structs do not match")
	}
}

func TestGetUserInfo(t *testing.T) {

	setup()

	// Test for user that exists
	input := "User1"
	expected := dataentity.UserInfo{
		Username:            "User1",
		PercentageScore:     0,
		NumOfCorrectAnswers: 0,
		NumOfQuestions:      0,
	}

	result, err := getUserInfo(input)
	if *result != expected {
		t.Errorf("getUserInfo(%v) = %v, expected %v", input, result, expected)
	}
	if err != nil {
		t.Errorf("getUserInfo(%v) error is %v, no error expected", input, err)
	}

	// Test for user that does not exist
	input = "NoUser"
	expected = dataentity.UserInfo{}

	result, err = getUserInfo(input)
	if *result != expected {
		t.Errorf("getUserInfo(%v) = %v, expected %v", input, result, expected)
	}
	if err == nil {
		t.Errorf("getUserInfo(%v) error is %v, expected error", input, err)
	}

	// Test for empty input
	input = ""
	expected = dataentity.UserInfo{}

	result, err = getUserInfo(input)
	if *result != expected {
		t.Errorf("getUserInfo(%v) = %v, expected %v", input, result, expected)
	}
	if err == nil {
		t.Errorf("getUserInfo(%v) error is %v, expected error", input, err)
	}

	teardown()
}
