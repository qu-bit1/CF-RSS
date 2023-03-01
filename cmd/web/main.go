package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://localhost:27017/"

func main() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Send a ping to confirm a successful connection
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Println("Pinged the primary node of the cluster. You successfully connected to MongoDB!")
}

//func main() {
//	fmt.Println("CF-RSS")
//	environment := "development"
//
//	var logger *zap.Logger
//	var loggerErr error
//	// new Development function prints logger and error in human friendly format where newP prints in json
//	if environment == "development" {
//		if logger, loggerErr = zap.NewDevelopment(); loggerErr != nil {
//			log.Fatalln(loggerErr)
//		}
//	} else {
//		if logger, loggerErr = zap.NewProduction(); loggerErr != nil {
//			log.Fatalln(loggerErr)
//		}
//	}
//	defer func(logger *zap.Logger) {
//		err := logger.Sync()
//		if err != nil {
//
//		}
//	}(logger)
//	zap.ReplaceGlobals(logger)
//	//zap.S().Info("these are some info logs")
//	//zap.S().Error("these are some error logs")
//	obj := cfapi.NewCodeforcesClient()
//	recentActions, err := obj.RecentActions(1)
//	if err != nil {
//		fmt.Println("error occurred while importing data from cfapi")
//		return
//	}
//
//	data, err1 := json.MarshalIndent(recentActions, "", " ")
//	if err1 != nil {
//		fmt.Println("error occurred while marshal indenting")
//		return
//	}
//	zap.S().Info(string(data))
//}
