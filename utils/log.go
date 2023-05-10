package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Logger struct {
	TimeStamp string            `json:"timestamp"`
	Severity  string            `json:"severity"`
	Message   string            `json:"message"`
	Labels    map[string]string `json:"labels"`
}

const (
	DDMMYYYYhhmmss = "2006-01-02 15:04:05"
)

func init() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)
}

func (l Logger) String() string {
	l.Message = strings.ReplaceAll(l.Message, "\"", "")
	if len(l.Labels) == 0 {
		return fmt.Sprintf("{\"timestamp\": \"%v\", \"severity\": \"%v\", \"message\": \"%v\"}", l.TimeStamp, l.Severity, l.Message)
	} else {
		lb, err := json.Marshal(l.Labels)
		if err != nil {
			return fmt.Sprintf("\"timestamp\": \"%v\", {\"severity\": \"%v\", \"message\": \"%v\"}", l.TimeStamp, l.Severity, l.Message)
		}
		return fmt.Sprintf("\"timestamp\": \"%v\", {\"severity\": \"%v\",\"message\": \"%v\", \"labels\": %v}", l.TimeStamp, l.Severity, l.Message, string(lb))
	}
}

func LogInfo(body string, t ...interface{}) {
	log.Println(Logger{TimeStamp: time.Now().Format(DDMMYYYYhhmmss), Severity: "INFO", Message: fmt.Sprintf(body, t...)}.String())
}

func LogError(body string, t ...interface{}) {
	log.Fatalln(Logger{TimeStamp: time.Now().Format(DDMMYYYYhhmmss), Severity: "Error", Message: fmt.Sprintf(body, t...)})
}
