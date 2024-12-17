package main

import (
	"ExperimentalLsp/analysis"
	"ExperimentalLsp/lsp"
	"ExperimentalLsp/rpc"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/prm/projects/go_projects/ExperimentalLsp/log.txt")
	logger.Println("Hey Logger Started!")
	fmt.Print("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	state := analysis.NewState()

	writer := os.Stdout
	for scanner.Scan() {
		msg := scanner.Bytes()
		method, contents, err := rpc.DecodeMessage(msg)
		if err != nil {
			logger.Printf("Got an error: %s", err)
			continue
		}
		handleMessage(logger, writer, state, method, contents)
	}
}

func handleMessage(logger *log.Logger, writer io.Writer, state analysis.State, method string, contents []byte) {
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
		writeResponce(writer, msg)
		logger.Println("Sent Reply")

	case "textDocument/didOpen":
		var request lsp.DidOpenTextDocumentNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("Hey, We couldn't parse this: %s", err)
			return
		}
		logger.Printf("Opened: %s", request.Params.TextDocument.URI)
		state.OpenDocument(request.Params.TextDocument.URI, request.Params.TextDocument.Text)

	case "textDocument/didChange":
		var request lsp.TextDocumentDidChageNotification
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("TextDocument/didChange: %s", err)
			return
		}
		logger.Printf("Changed: %s", request.Params.TextDocument.URI)
		for _, change := range request.Params.ContentChanges {
			state.UpdateDocument(request.Params.TextDocument.URI, change.Text)
		}
	case "textDocument/hover":
		var request lsp.HoverRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("TextDocument/hover: %s", err)
			return
		}

		// crate responce
		responce := state.Hover(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		// write back
		writeResponce(writer, responce)

	case "textDocument/codeAction":
		var request lsp.CodeActionRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("TextDocument/codeAction: %s", err)
			return
		}

		// crate responce
		responce := state.TextDocumentCodeAction(request.ID, request.Params.TextDocument.URI, request.Params.Position)
		// write back
		writeResponce(writer, responce)
	}
}

func writeResponce(writer io.Writer, msg any) {
	reply := rpc.EncodeMessage(msg)
	writer.Write([]byte(reply))
}
func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("hey you did not give me a good file")
	}

	return log.New(logfile, "[ExperimentalLsp]", log.Ldate|log.Ltime|log.Lshortfile)
}
