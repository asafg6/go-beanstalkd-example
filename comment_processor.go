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
	"os"
	"fmt"
	"time"
)

type DiskCommentProcessor struct {
	dir string
}

func (processor *DiskCommentProcessor) DoProcess(comment *Comment) error {
	filePath := fmt.Sprintf("%s/%s_%s.txt", processor.dir, comment.Date.Format(time.RFC3339), comment.UserName)
	commentFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer commentFile.Close()
	commentFile.WriteString(comment.Text)
	return nil
}

func MakeNewCommentProcessor(dir string) *DiskCommentProcessor {
	return &DiskCommentProcessor{dir: dir}
}
