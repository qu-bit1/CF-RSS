package main

import (
	"encoding/json"
	"fmt"
	"github.com/qu-bit1/project_new/pkg/cfapi"
	"go.uber.org/zap"
	"log"
)

func main() {
	fmt.Println("CF-RSS")
	environment := "development"

	var logger *zap.Logger
	var loggerErr error
	// new Development function prints logger and error in human friendly format where newP prints in json
	if environment == "development" {
		if logger, loggerErr = zap.NewDevelopment(); loggerErr != nil {
			log.Fatalln(loggerErr)
		}
	} else {
		if logger, loggerErr = zap.NewProduction(); loggerErr != nil {
			log.Fatalln(loggerErr)
		}
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)
	zap.ReplaceGlobals(logger)
	//zap.S().Info("these are some info logs")
	//zap.S().Error("these are some error logs")
	obj := cfapi.NewCodeforcesClient()
	recentActions, err := obj.RecentActions(1)
	if err != nil {
		fmt.Println("error occurred while importing data from cfapi")
		return
	}

	data, err1 := json.MarshalIndent(recentActions, "", " ")
	if err1 != nil {
		fmt.Println("error occurred while marshal indenting")
		return
	}
	zap.S().Info(string(data))
}
