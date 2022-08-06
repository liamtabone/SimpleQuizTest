package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math"
	proto "simple_quiz/gen/proto"
	"simple_quiz/server/dataentity"
	"sort"
)

// GetResult expects a user in the request message, calculates the user percentile (i.e. how many questions
// the user answered correctly compared to other users) and returns the percentile in the response message
func (s *SimpleQuizServiceServer) GetResult(ctx context.Context, user *proto.User) (*proto.Result, error) {
	percentile, err := getUserPercentile(user.Username)
	return &proto.Result{Percentile: percentile}, err
}

func loadAllMarks() {
	file, _ := ioutil.ReadFile("data/userInfo.json")

	userInfoData = []dataentity.UserInfo{}

	_ = json.Unmarshal([]byte(file), &userInfoData)
}

func getUserPercentile(username string) (float32, error) {
	userInfo, err := getUserInfo(username)
	if err != nil {
		return float32(math.NaN()), err
	}

	return calcPercentile(getMarksInSlice(userInfoData), userInfo.PercentageScore)
}

func getPercentageScore(x int, y int) (percentage float32, err error) {

	if y == 0 {
		return float32(math.NaN()), errors.New("cannot divide by 0")
	}

	if x > y {
		return float32(math.NaN()), errors.New("numerator cannot be bigger than the denominator")
	}

	if x < 0 || y < 0 {
		return float32(math.NaN()), errors.New("negative numbers are not accepted")
	}

	return float32(float32(x) / float32(y) * 100), nil
}

func getMarksInSlice(userMark []dataentity.UserInfo) []float32 {
	var mark []float32

	for i := 0; i < len(userMark); i++ {
		mark = append(mark, userMark[i].PercentageScore)
	}

	return mark
}

// sort function for an array of type float32
func sortFloat32(float32Values []float32) []float32 {
	float32AsFloat64Values := make([]float64, len(float32Values))

	for i, val := range float32Values {
		float32AsFloat64Values[i] = float64(val)
	}
	sort.Float64s(float32AsFloat64Values)

	for i, val := range float32AsFloat64Values {
		float32Values[i] = float32(val)
	}

	return float32Values
}

// calculate the percentile of parameter userPercentage compared to the array parameter allMarks
func calcPercentile(allMarks []float32, userPercentange float32) (percentile float32, err error) {
	length := len(allMarks)
	if length == 0 {
		return float32(math.NaN()), errors.New("error on mark array")
	}

	if length == 1 {
		return float32(math.NaN()), errors.New("no marks available to compare with")
	}

	if userPercentange < 0 || userPercentange > 100 {
		return float32(math.NaN()), errors.New("user percentage out of bounds")
	}

	allMarks = sortFloat32(allMarks)

	i := 0
	for i < len(allMarks) && allMarks[i] < userPercentange {
		i++
	}

	result := float32(i*100) / float32(len(allMarks)-1)
	return result, nil

}
