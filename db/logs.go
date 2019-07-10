package db

import (
	"fmt"
	"log"
	"os"
	"time"
)

const logfile = "development.log"

var logFileName string

func OpenLogFile() {
	t := time.Now()
	oldFileName := logFileName

	logFileName = t.Format("2006-01-02") + "_" + logfile

	if len(oldFileName) == 0 || oldFileName != logFileName {
		fmt.Printf("Logging to file %v\n", logFileName)

		lf, err := os.OpenFile("./logs/"+logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatalf("OpenLogfile: os.OpenFile: %s", err)
		}

		log.SetOutput(lf)
	}
}

func LogFatalf(format string, v ...interface{}) {
	format = "[ERR|FATAL]|" + format + "\n"
	fmt.Printf(format, v...)
	log.Fatalf(format, v...)
}

func Logf(format string, v ...interface{}) {
	format = format + "\n"
	fmt.Printf(format, v...)
	log.Printf(format, v...)
}

func LogEF(format string, v ...interface{}) {
	fmt.Printf("[0;31m"+format+"[39m\n", v...)
	log.Printf("[ERR]|"+format+"\n", v...)
}
