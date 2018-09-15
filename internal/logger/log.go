package logger

import (
	"io/ioutil"
	"log"
	"os"
)

var logger *log.Logger
var output = ioutil.Discard

func SetUpLogger(logFile string, logOutput bool) {
	if logOutput {
		outfile, err := os.OpenFile(logFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		if err != nil {
			log.Fatalln("Couldn't open logger file: " + err.Error())
		}
		output = outfile
	}

	logger = log.New(output, "", log.LstdFlags)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Println(v ...interface{}) {
	logger.Println(v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}
