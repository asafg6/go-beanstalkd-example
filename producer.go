//Copyright 2017 Asaf Gur
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"time"
	"fmt"
)

type Producer struct {
	PapaBeanstalk
	protocol CommentProtocol
}

func (producer *Producer) UseTube(tubeName string) error {
	return producer.serverConnection.Use(tubeName)
}

func (producer *Producer) PutComment(comment *Comment) error {
	body, err := producer.protocol.Encode(comment)
	if err != nil {
		return err
	}
	priority := uint32(10)
	delay := 0 * time.Second
	time_to_run := 20 * time.Second
	jobId, err := producer.serverConnection.Put(body, priority, delay, time_to_run)
	if err != nil {
		return err
	}
	fmt.Println("inserted Job id: ", jobId)
	return nil
}

func MakeNewProducer(serverAdress string, protocol CommentProtocol) *Producer {
	producer := Producer{protocol: protocol}
	producer.ServerAddress = serverAdress
	return &producer
}