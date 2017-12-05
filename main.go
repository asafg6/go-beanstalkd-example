package main

import (
	"time"
	"os"
	"fmt"
)

func ProducerMain(tubes []string) {
	var comments = []Comment{
		{UserName: "some_user", Text:"i love your cat", Date: time.Now()},
		{UserName: "some_other_user", Text:"i prefer dogs", Date: time.Now()},
		{UserName: "another_user", Text:"please close this thread", Date: time.Now()},
		{UserName: "admin", Text:"thread closed - not relevant", Date: time.Now()},
	}
	protocol := MakeJsonCommentProtocol()
	producer := MakeNewProducer("localhost:11300", protocol)
	producer.Connect()
	defer producer.Close()

	for _, tube := range tubes {
		producer.UseTube(tube)
		for _, comment := range comments {
			producer.PutComment(&comment)
		}
	}

}

func WorkerMain(commentsDir string, tube string) {
	protocol := MakeJsonCommentProtocol()
	os.Mkdir(commentsDir, 0777)
	processor := MakeNewCommentProcessor(commentsDir)
	worker := MakeNewWorker("localhost:11300", protocol, processor, tube)
	worker.Connect()
	defer worker.Close()
	for {
		worker.ProcessJob()
	}

}

func printUsage() {
	fmt.Println("Usage: example-app worker/producer")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}
	if os.Args[1] == "worker1" {
		WorkerMain("first_comments", "first")
	} else if os.Args[1] == "worker2" {
		WorkerMain("second_comments", "second")
	} else if os.Args[1] == "producer" {
		ProducerMain([]string{"first", "second"})
	} else {
		printUsage()
	}
}
