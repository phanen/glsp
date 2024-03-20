package main

import (
	"bufio"
	"encoding/json"
	"glsp/lsp"
	"glsp/rpc"
	"log"
	"os"
)

func main() {
	logger := getLogger("/home/phan/b/glsp/log.txt")
	logger.Println("start logging...")

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

// handle method field... top down
func handleMessage(logger *log.Logger, method string, contents []byte) {
	logger.Print(string(contents))
	// logger.Printf("Recive msg with method: %s", method)

	switch method {
	case "initialize":
		var request lsp.InitiailizeRequest
		if err := json.Unmarshal(contents, &request); err != nil {
			logger.Printf("could not parse %s", err)
		}
		logger.Printf("connected to: %s %s", request.Params.ClientInfo.Name, request.Params.ClientInfo.Version)
	}
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("not a good file")
	}
	return log.New(logfile, "[glsp] ", log.Lshortfile)
}
