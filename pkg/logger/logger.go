package logger

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/Sirupsen/logrus"
)

type ILogger interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	NewErrorf(format string, args ...interface{}) error
	NewError(message string) error
	LogNewError(message string) error
	Info(args ...interface{})
	Debug(args ...interface{})
	Error(args ...interface{})
	Panic(args ...interface{})
	Initialise()
}

type RealLogger struct {
	log *logrus.Logger
}

func (al *RealLogger) Infof(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		formatWithDetails := concatLogDetails(line, filepath.Base(file), format)
		al.log.Infof(formatWithDetails, args)
	} else {
		al.log.Infof(format, args)
	}

}

func (al *RealLogger) Debugf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		formatWithDetails := concatLogDetails(line, filepath.Base(file), format)
		al.log.Debugf(formatWithDetails, args)
	} else {
		al.log.Debugf(format, args)
	}
}

func (al *RealLogger) Errorf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		formatWithDetails := concatLogDetails(line, filepath.Base(file), format)
		al.log.Errorf(formatWithDetails, args)
	} else {
		al.log.Errorf(format, args)
	}
}

func (al *RealLogger) Panicf(format string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		formatWithDetails := concatLogDetails(line, filepath.Base(file), format)
		al.log.Panicf(formatWithDetails, args)
	} else {
		al.log.Panicf(format, args)
	}
}

func (al *RealLogger) NewErrorf(format string, args ...interface{}) error {
	return fmt.Errorf(format, args)
}

func (al *RealLogger) NewError(message string) error {
	return errors.New(message)
}

func (al *RealLogger) LogNewError(message string) error {
	al.log.Error(message)
	return errors.New(message)
}

func (al *RealLogger) Info(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Info(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Info(args)
	}

}

func (al *RealLogger) Debug(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Debug(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Debug(args)
	}
}

func (al *RealLogger) Error(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Error(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Error(args)
	}
}

func (al *RealLogger) Panic(args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		al.log.Panic(filepath.Base(file), "(", line, ") ", args)
	} else {
		al.log.Panic(args)
	}
}

func (al *RealLogger) Initialise() {
	al.log = logrus.New()
	al.log.Formatter = new(logrus.TextFormatter)
	al.log.Out = os.Stdout
}

func concatLogDetails(lineNumber int, filepath, format string) string {
	var buffer bytes.Buffer
	buffer.WriteString(filepath)
	buffer.WriteString("(")
	buffer.WriteString(strconv.Itoa(lineNumber))
	buffer.WriteString(")")
	buffer.WriteString(format)
	return buffer.String()
}
