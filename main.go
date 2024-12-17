package main

import (
	"ExperimentalLsp/lsp"
	"ExperimentalLsp/rpc"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/prm/projects/go_projects/ExperimentalLsp/log.txt")
	logger.Println("Hey Logger Started!")
	fmt.Print("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, method, contents)
	}
}

func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Printf("Received msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitializeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, We couldn't parse this: %s", err)
		}
		logger.Printf("Connected to: %s %s",
			request.Params.ClientInfo.Name,
			request.Params.ClientInfo.Version)

		// lets reply
		msg := lsp.NewInitializeResponse(request.ID)
		reply := rpc.EncodeMessage(msg)
		writer := os.Stdout
		writer.Write([]byte(reply))
		logger.Println("Sent Reply")

	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey you did not give me a good file")
	}

	return log.New(logfile, "[ExperimentalLsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
