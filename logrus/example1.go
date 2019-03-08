package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})
	//log.SetFormatter(&log.TextFormatter{})
	//log.SetReportCaller(true)
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the info severity or above.
	//log.SetLevel(log.InfoLevel)
}

func foo() {
	log.WithFields(log.Fields{
		"prefix":      "sensor",
		"temperature": -4,
		"handler-1":   "error x",
	}).Info("Temperature changes")

	log.WithFields(log.Fields{
		"prefix":      "outro error1",
		"temperature": -10,
		"handler-1":   "error y",
	}).Warn("error warn")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatal("The ice breaks!")

	log.WithFields(log.Fields{
		"omg":    false,
		"number": 900,
	}).Error("Error here")

	log.Debug("Useful debugging information.")

}

func main() {
	foo()
}
