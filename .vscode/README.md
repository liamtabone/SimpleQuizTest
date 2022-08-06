Generate proto files (note that these have been generated and included in the Zip file) :
    export PATH="$PATH:$(go env GOPATH)/bin"
    protoc --proto_path=proto proto/*.proto --go_out=gen
    protoc --proto_path=proto proto/*.proto --go-grpc_out=ge

Build server:
    cd server
    go build *.go

Build client (from project root folder)):
    cd client
    go build *.go

Run server (from project root folder)):
    cd server
    ./server

Start quiz from client (from project root folder)):
    (Open new terminals)
    cd client
    ./client startQuiz

Get result from client (from project root folder)):
    cd client
    ./client result