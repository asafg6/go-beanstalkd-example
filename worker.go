package main

import (
	"fmt"
	"time"
	"github.com/iwanbk/gobeanstalk"
	"os"
)

type CommentWorker struct {
	PapaBeanstalk
	protocol CommentProtocol
	processor CommentProcessor
	tube string
}


func (worker *CommentWorker) ProcessJob() {
	fmt.Println("reserving job")
	job, err := worker.serverConnection.Reserve()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("got job Id: ", job.ID)
	comment, err := worker.protocol.Decode(job.Body)
	if err != nil {
		worker.handleError(job, err)
		return
	}
	processError := worker.processor.DoProcess(comment)
	if processError != nil {
		worker.handleError(job, err)
		return
	}
	fmt.Println("processed job id: ", job.ID)
	worker.serverConnection.Delete(job.ID)
}

func (worker *CommentWorker) watch() error {
	watching, err := worker.serverConnection.Watch(worker.tube)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("watching ", watching, " tubes")
	return nil
}

func (worker *CommentWorker) Connect() {
	worker.PapaBeanstalk.Connect()
	worker.watch()
}

func (worker *CommentWorker) handleError(job *gobeanstalk.Job, err error) {
	fmt.Println(err)
	priority := uint32(5)
	delay := 0 * time.Second
	worker.serverConnection.Release(job.ID ,priority, delay) // hey I can't handle this
}

func MakeNewWorker(serverAddress string, protocol CommentProtocol, processor CommentProcessor,
								tubeToListenOn string) *CommentWorker {
	worker := CommentWorker{protocol: protocol, processor: processor, tube: tubeToListenOn}
	worker.ServerAddress = serverAddress
	return &worker
}