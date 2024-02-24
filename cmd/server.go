package main

import (
	"log"

	"go.uber.org/zap/zapcore"
	"honnef.co/go/tools/config"
)

func run() {
	conf, err := config.Load(configFile)

}
