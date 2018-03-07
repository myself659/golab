package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

func caller() {
	call()
}
func call() {
	if pc, file, line, ok := runtime.Caller(1); ok {
		fName := runtime.FuncForPC(pc).Name()
		fmt.Println(fName, file, line)
	}
}

func main() {
	log := logrus.New()

	log.Out = os.Stdout
	log.Level = logrus.DebugLevel
	log.Debugln("hello")
	log.WithField("key", "value").Debugf("%d", 40)
	caller()

}
