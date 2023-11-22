package common

import (
	"os"

	"enigmacamp.com/be-enigma-laundry/config"
	modelutil "enigmacamp.com/be-enigma-laundry/utils/model_util"
	"github.com/sirupsen/logrus"
)

type MyLogger interface {
	InitializeLogger() error // membuat sebuah file logger
	LogInfo(requestLog modelutil.RequestLog)
	LogWarn(requestLog modelutil.RequestLog)
	LogFatal(requestLog modelutil.RequestLog)
}

type myLogger struct {
	cfg config.LogFileConfig
	log *logrus.Logger
}

func (m *myLogger) InitializeLogger() error {
	file, err := os.OpenFile(m.cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	m.log = logrus.New()
	m.log.SetOutput(file)
	return nil
}

func (m *myLogger) LogInfo(requestLog modelutil.RequestLog) {
	m.log.Info(requestLog)
}

func (m *myLogger) LogWarn(requestLog modelutil.RequestLog) {
	m.log.Warn(requestLog)
}

func (m *myLogger) LogFatal(requestLog modelutil.RequestLog) {
	m.log.Fatal(requestLog)
}

func NewMyLogger(cfg config.LogFileConfig) MyLogger {
	return &myLogger{cfg: cfg}
}
