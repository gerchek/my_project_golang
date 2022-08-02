package logging

import (
	"fmt"
	"io"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

const logFile = "logs/project.log"

var log *logrus.Logger

func init() {
	log = logrus.New()
	// log to console and file
	log.SetReportCaller(true)

	log.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			filename := f.File
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})

	f, err := os.OpenFile("logs/project.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	wrt := io.MultiWriter(os.Stdout, f)

	log.SetOutput(wrt)

}

func Log() *logrus.Logger {
	return log
}
