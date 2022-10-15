package utils

import "github.com/sirupsen/logrus"

//for local testing to see the logs lambda
var (
	Logger = NewLogger() // Global Variable for Logging across all packages
)

func NewLogger()*logrus.Logger{
	log := logrus.New()

	return log
}