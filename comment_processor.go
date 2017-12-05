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
