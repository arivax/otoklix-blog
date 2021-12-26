package logger

import (
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"go-otoklix-blog/infra/config"
)

var (
	host, err      = os.Hostname()
	app, err1      = os.Executable()
	logfile        = config.Get("LOG_FILE")
	standardFields = log.Fields{
		"host": host,
		"app":  filepath.Base(app),
	}
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.ErrorLevel)
	log.SetOutput(os.Stdout)
}

// Fatal ...
func Fatal(msg string, fields log.Fields) {
	/*file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)*/
	log.WithFields(standardFields).WithFields(fields).Fatal(msg)
}

// Error ...
func Error(msg string, fields log.Fields) {
	/*file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)*/
	log.WithFields(standardFields).WithFields(fields).Error(msg)
}

// Warn ...
func Warn(msg string, fields log.Fields) {
	/*file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)*/
	log.WithFields(standardFields).WithFields(fields).Warn(msg)
}

// Info ...
func Info(msg string, fields log.Fields) {
	/*file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)*/
	log.WithFields(standardFields).WithFields(fields).Info(msg)
}

// Debug ...
func Debug(msg string, fields log.Fields) {
	/*file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)*/
	log.WithFields(standardFields).WithFields(fields).Debug(msg)
}

// Trace ...
func Trace(msg string, fields log.Fields) {
	/*file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)*/
	log.WithFields(standardFields).WithFields(fields).Trace(msg)
}
