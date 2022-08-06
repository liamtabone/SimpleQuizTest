package cmd

import (
	"fmt"
	"log"
	"net"
	proto "simple_quiz/gen/proto"
	dataentity "simple_quiz/server/dataentity"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type SimpleQuizServiceServer struct {
	proto.UnimplementedSimpleQuizServiceServer
}

var (
	questionsJsonFileName string
	userInfoJsonFileName  string
	questionData          []dataentity.Question
	userInfoData          []dataentity.UserInfo
)

var rootCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize quiz",
	Long:  `Initialize quiz by loading all questions to memory from JSON.`,

	Run: func(cmd *cobra.Command, args []string) {
		questionsJsonFileName = "data/questions.json"
		userInfoJsonFileName = "data/userInfo.json"
		loadAllQuestions()
		loadAllMarks()

		lis, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalln(err)
		}

		grpcServer := grpc.NewServer()

		proto.RegisterSimpleQuizServiceServer(grpcServer, &SimpleQuizServiceServer{})

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Println(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
