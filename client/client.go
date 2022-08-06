package main

import (
	"context"
	"fmt"
	"log"
	proto "simple_quiz/gen/proto"
	"strconv"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.SimpleQuizServiceClient

func initClient() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Println(err)
	}

	client = proto.NewSimpleQuizServiceClient(conn)
}

func getResult(username string) {
	resp, err := client.GetResult(context.Background(), &proto.User{Username: username})

	if err == nil {
		fmt.Printf("You were better than %v%% of all quizzers \n", resp.Percentile)
	} else {
		fmt.Println(err)
	}
}

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Root",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	resultCmd = &cobra.Command{
		Use:   "result",
		Short: "Get result for username",
		Run: func(cmd *cobra.Command, args []string) {
			username := waitForUsername()
			getResult(username)
		},
	}

	startQuizCmd = &cobra.Command{
		Use:   "startQuiz",
		Short: "Start Quiz command",
		Run: func(cmd *cobra.Command, args []string) {
			var resp *proto.Question
			username := waitForUsername()
			_, err := client.NewUser(context.Background(), &proto.User{Username: username})

			if err == nil {
				resp, err = client.GetQuestion(context.Background(), &proto.User{Username: username})

				if err == nil {
					for resp.QuestionId >= 0 && err == nil {
						printQuestion(resp)
						userAnswer := waitForAnswer(len(resp.Choices))

						client.GetAnswer(context.Background(), &proto.Answer{Username: username, QuestionId: resp.QuestionId, AnswerId: userAnswer})

						resp, err = client.GetQuestion(context.Background(), &proto.User{Username: username})

						if err != nil {
							fmt.Println(err)
						}
					}

					if err == nil {
						getResult(username)
					}
				}
			}

			if err != nil {
				fmt.Println(err)
			}
		},
	}
)

func printQuestion(question *proto.Question) {
	fmt.Println()
	fmt.Println(question.Text)
	fmt.Println()
	for i := 0; i < len(question.Choices); i++ {
		fmt.Println(question.Choices[i].ChoiceId + ": " + question.Choices[i].Text)
	}
	fmt.Println()
}

func waitForUsername() string {
	var username string
	fmt.Print("Enter a username (Please note that white spaces are not allowed): ")
	fmt.Scanln(&username)
	return username
}

// Prompt the user to add an answer, continue asking until the
// user's guess is in the correct range
func waitForAnswer(numOfChoice int) string {
	var userAnswer string
	var userAnswerInt int

	for !(userAnswerInt > 0 && userAnswerInt <= numOfChoice) {
		fmt.Println()
		fmt.Print("Enter your answer ID: ")
		_, err := fmt.Scanln(&userAnswer)
		userAnswerInt, _ = strconv.Atoi(userAnswer)

		if err != nil {
			fmt.Println(err)
		}

		if userAnswerInt < 1 || userAnswerInt > numOfChoice {
			fmt.Println("Answer not in correct format or in the expected range of values, please enter another answer")
		}
	}

	return userAnswer
}

func init() {
	rootCmd.AddCommand(startQuizCmd)
	rootCmd.AddCommand(resultCmd)
}

func main() {
	initClient()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
