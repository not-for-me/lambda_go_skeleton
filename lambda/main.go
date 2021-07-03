package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

func init() {
	log.Println("Init Func")
}

func handleRequest() (string, error) {
	log.Println("Body")
	return "Hello, Go Lambda", nil
}

func main() {
	lambda.Start(handleRequest)
}