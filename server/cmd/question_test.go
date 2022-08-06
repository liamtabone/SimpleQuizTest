package cmd

import (
	"simple_quiz/server/dataentity"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetQuestionById(t *testing.T) {

	setup()

	input := 0
	expected := "Test Question 1"

	result := getQuestionById(input)
	if result.Question != expected {
		t.Errorf("getQuestionById(%v).Question = %v, expected %v", input, result, expected)
	}

	input = 1
	expected = "Test Question 2"

	result = getQuestionById(input)
	if result.Question != expected {
		t.Errorf("getQuestionById(%v).Question = %v, expected %v", input, result, expected)
	}

	input = -1
	expectedStruct := dataentity.Question{}

	result = getQuestionById(input)
	if !cmp.Equal(result, expectedStruct) {
		t.Errorf("getQuestionById(%v) = %v, expected %v", input, result, expected)
	}

	input = 3
	expectedStruct = dataentity.Question{}
	result = getQuestionById(input)
	if !cmp.Equal(result, expectedStruct) {
		t.Errorf("getQuestionById(%v) = %v, expected %v", input, result, expected)
	}

	teardown()
}

func TestCreateQuestionMsg(t *testing.T) {

	input := 0
	expectedQuestionText := "Test Question 1"
	expectedQuestionId := int32(0)
	expectedChoiceId0 := "1"
	expectedChoiceIdText0 := "Choice 1"

	result := createQuestionMsg(input)
	if result.Text != expectedQuestionText {
		t.Errorf("createQuestionMsg(%v).Text = %v, expected %v", input, &result, expectedQuestionText)
	}
	if result.QuestionId != expectedQuestionId {
		t.Errorf("createQuestionMsg(%v).QuestionId = %v, expected %v", input, &result, expectedQuestionId)
	}
	if result.Choices[0].ChoiceId != expectedChoiceId0 {
		t.Errorf("createQuestionMsg(%v).Choices[0].ChoiceId = %v, expected %v", input, &result, expectedChoiceId0)
	}
	if result.Choices[0].Text != expectedChoiceIdText0 {
		t.Errorf("createQuestionMsg(%v).Choices[0].Text = %v, expected %v", input, &result, expectedChoiceIdText0)
	}

	input = 3
	expectedQuestionId = int32(-1)

	result = createQuestionMsg(input)
	if result.QuestionId != expectedQuestionId {
		t.Errorf("createQuestionMsg(%v).QuestionId = %v, expected %v", input, &result, expectedQuestionId)
	}

	input = -1
	expectedQuestionId = int32(-1)

	result = createQuestionMsg(input)
	if result.QuestionId != expectedQuestionId {
		t.Errorf("createQuestionMsg(%v).QuestionId = %v, expected %v", input, &result, expectedQuestionId)
	}
}
