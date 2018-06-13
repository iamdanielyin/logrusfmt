package logrusfmt

import (
	log "github.com/sirupsen/logrus"
	"fmt"
)

type SimpleTextFormatter struct {
}

func (f *SimpleTextFormatter) Format(entry *log.Entry) ([]byte, error) {
	output := fmt.Sprintf("%s %v\n", entry.Time.Format("2006-01-02 03:04:05"), entry.Message)
	return []byte(output), nil
}
