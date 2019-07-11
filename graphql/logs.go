package graphql

import (
	"fmt"
	"log"
)

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
