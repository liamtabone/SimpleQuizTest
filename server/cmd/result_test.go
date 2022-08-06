package cmd

import (
	"math"
	"reflect"
	"testing"
)

func TestGetPercentageScore(t *testing.T) {

	inputX := 5
	inputY := 10
	expected := float32(50)

	result, err := getPercentageScore(inputX, inputY)
	if result != expected {
		t.Errorf("getPercentageScore(%v, %v) = %v, expected %v", inputX, inputY, &result, expected)
	}

	if err != nil {
		t.Errorf("getPercentageScore(%v, %v) got error '%v', did not expect error", inputX, inputY, &result)
	}

	inputX = 0
	inputY = 10
	expected = float32(0)

	result, err = getPercentageScore(inputX, inputY)
	if result != expected {
		t.Errorf("getPercentageScore(%v, %v) = %v, expected %v", inputX, inputY, &result, expected)
	}

	if err != nil {
		t.Errorf("getPercentageScore(%v, %v) got error '%v', did not expect error", inputX, inputY, &result)
	}

	inputX = 10
	inputY = 10
	expected = float32(100)

	result, err = getPercentageScore(inputX, inputY)
	if result != expected {
		t.Errorf("getPercentageScore(%v, %v) = %v, expected %v", inputX, inputY, &result, expected)
	}

	if err != nil {
		t.Errorf("getPercentageScore(%v, %v) got error '%v', did not expect error", inputX, inputY, &result)
	}

	inputX = 100
	inputY = 10
	expected = float32(math.NaN())

	result, err = getPercentageScore(inputX, inputY)
	if !math.IsNaN(float64(result)) {
		t.Errorf("getPercentageScore(%v, %v) = %v, expected %v", inputX, inputY, &result, expected)
	}

	if err == nil {
		t.Errorf("getPercentageScore(%v, %v) expected error", inputX, inputY)
	}

	inputX = 5
	inputY = 0
	expected = float32(math.NaN())

	result, err = getPercentageScore(inputX, inputY)
	if !math.IsNaN(float64(result)) {
		t.Errorf("getPercentageScore(%v, %v) = %v, expected %v", inputX, inputY, &result, expected)
	}

	if err == nil {
		t.Errorf("getPercentageScore(%v, %v) expected error", inputX, inputY)
	}

	inputX = -1
	inputY = 2
	expected = float32(math.NaN())

	result, err = getPercentageScore(inputX, inputY)
	if !math.IsNaN(float64(result)) {
		t.Errorf("getPercentageScore(%v, %v) = %v, expected %v", inputX, inputY, &result, expected)
	}

	if err == nil {
		t.Errorf("getPercentageScore(%v, %v) expected error", inputX, inputY)
	}

	inputX = 1
	inputY = -2
	expected = float32(math.NaN())

	result, err = getPercentageScore(inputX, inputY)
	if !math.IsNaN(float64(result)) {
		t.Errorf("getPercentageScore(%v, %v) = %v, expected %v", inputX, inputY, &result, expected)
	}

	if err == nil {
		t.Errorf("getPercentageScore(%v, %v) expected error", inputX, inputY)
	}
}

func TestCalcPercentile(t *testing.T) {

	inputMarks := []float32{80, 40, 60, 90, 50}
	inputScore := float32(90)
	expected := float32(100)

	result, err := calcPercentile(inputMarks, inputScore)
	if result != expected {
		t.Errorf("calcPercentile(%v, %v) = %v, expected %v", inputMarks, inputScore, &result, expected)
	}

	if err != nil {
		t.Errorf("calcPercentile(%v, %v) got error '%v', did not expect error", inputMarks, inputScore, err)
	}

	inputMarks = []float32{80, 40, 60, 90, 50}
	inputScore = float32(40)
	expected = float32(0)

	result, err = calcPercentile(inputMarks, inputScore)
	if result != expected {
		t.Errorf("calcPercentile(%v, %v) = %v, expected %v", inputMarks, inputScore, &result, expected)
	}

	if err != nil {
		t.Errorf("calcPercentile(%v, %v) got error '%v', did not expect error", inputMarks, inputScore, err)
	}

	inputMarks = []float32{80, 40, 60, 90, 50}
	inputScore = float32(60)
	expected = float32(50)

	result, err = calcPercentile(inputMarks, inputScore)
	if result != expected {
		t.Errorf("calcPercentile(%v, %v) = %v, expected %v", inputMarks, inputScore, &result, expected)
	}

	if err != nil {
		t.Errorf("calcPercentile(%v, %v) got error '%v', did not expect error", inputMarks, inputScore, err)
	}

	inputMarks = []float32{80, 40, 60, 90, 100}
	inputScore = float32(100)
	expected = float32(100)

	result, err = calcPercentile(inputMarks, inputScore)
	if result != expected {
		t.Errorf("calcPercentile(%v, %v) = %v, expected %v", inputMarks, inputScore, &result, expected)
	}

	if err != nil {
		t.Errorf("calcPercentile(%v, %v) got error '%v', did not expect error", inputMarks, inputScore, err)
	}

	inputMarks = []float32{80, 40, 60, 90, 50}
	inputScore = float32(101)
	expected = float32(math.NaN())

	result, err = calcPercentile(inputMarks, inputScore)
	if !math.IsNaN(float64(result)) {
		t.Errorf("calcPercentile(%v, %v) = %v, expected %v", inputMarks, inputScore, &result, expected)
	}
	if err == nil {
		t.Errorf("calcPercentile(%v, %v) expected error", inputMarks, inputScore)
	}

	inputMarks = []float32{80, 40, 60, 90, 50}
	inputScore = float32(-1)
	expected = float32(math.NaN())

	result, err = calcPercentile(inputMarks, inputScore)
	if !math.IsNaN(float64(result)) {
		t.Errorf("calcPercentile(%v, %v) = %v, expected %v", inputMarks, inputScore, &result, expected)
	}

	if err == nil {
		t.Errorf("calcPercentile(%v, %v) expected error", inputMarks, inputScore)
	}

	inputMarks = []float32{}
	inputScore = float32(50)
	expected = float32(math.NaN())

	result, err = calcPercentile(inputMarks, inputScore)
	if !math.IsNaN(float64(result)) {
		t.Errorf("calcPercentile(%v, %v) = %v, expected %v", inputMarks, inputScore, &result, expected)
	}

	if err == nil {
		t.Errorf("calcPercentile(%v, %v) expected error", inputMarks, inputScore)
	}
}

func TestSortFloat32(t *testing.T) {

	input := []float32{80, 40, 60, 90, 50}
	expected := []float32{40, 50, 60, 80, 90}

	result := sortFloat32(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sortFloat32(%v) = %v, expected %v", input, &result, expected)
	}

	input = []float32{}
	expected = []float32{}

	result = sortFloat32(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("sortFloat32(%v) = %v, expected %v", input, &result, expected)
	}
}

func TestGetMarksInSlice(t *testing.T) {

	setup()

	input := userInfoData
	expected := []float32{0, 50, 100, 70, 30}

	result := getMarksInSlice(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getMarksInSlice(%v) = %v, expected %v", input, &result, expected)
	}

	teardown()
}

func TestGetUserPercentile(t *testing.T) {

	setup()

	input := "User1"
	expected := float32(0)

	result, err := getUserPercentile(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getUserPercentile(%v) = %v, expected %v", input, &result, expected)
	}
	if err != nil {
		t.Errorf("getUserPercentile(%v) got error '%v', did not expect error", input, err)
	}

	input = "User3"
	expected = float32(100)

	result, err = getUserPercentile(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getUserPercentile(%v) = %v, expected %v", input, &result, expected)
	}
	if err != nil {
		t.Errorf("getUserPercentile(%v) got error '%v', did not expect error", input, expected)
	}

	input = "User2"
	expected = float32(50)

	result, err = getUserPercentile(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("getUserPercentile(%v) = %v, expected %v", input, &result, expected)
	}
	if err != nil {
		t.Errorf("getUserPercentile(%v) got error '%v', did not expect error", input, err)
	}

	input = "UnknownUser"
	expected = float32(math.NaN())

	result, err = getUserPercentile(input)
	if !math.IsNaN(float64(result)) {
		t.Errorf("getUserPercentile(%v) = %v, expected %v", input, &result, expected)
	}
	if err == nil {
		t.Errorf("getUserPercentile(%v) expected error", input)
	}

	teardown()
}
