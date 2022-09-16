package file_manager

import (
	"github.com/nanih98/aws-sso/logger"
)

type FileProcessor struct {
	log *logger.CustomLogger
}

func NewFileProcessor(log *logger.CustomLogger) *FileProcessor {
	return &FileProcessor{log: log}
}
