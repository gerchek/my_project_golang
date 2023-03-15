package main

import (
	"fmt"
	"my_project/internal/app"
	"my_project/internal/setup/constructor"
	"my_project/pkg/gormclient"
	"my_project/pkg/logging"
	customRedis "my_project/pkg/redis"
)

func main() {
	// log.Print("logger initializing...")
	logger := logging.Log()

	// project.NewProject(logger)

	logger.Info("initializing postgres config...")
	gormConfig := gormclient.NewGormConfig("newuser", "newpassword", "127.0.0.1", "5432", "my_project")

	logger.Info("connecting to postgres database...")
	client, err := gormclient.NewClient(gormConfig)
	if err != nil {
		logger.Fatal(err)
		fmt.Println(client)
	}

	logger.Info("initializing redis config...")
	redisConfig := customRedis.NewRedisConfig("localhost:6379")

	logger.Info("connecting to redis database...")
	redisClient := customRedis.NewRedisClient(redisConfig)

	logger.Info("setting up all repository, service, controller...")

	constructor.SetConstructor(client, redisClient, logger)

	logger.Info("initializing a new app...")
	app := app.NewApp(logger, redisClient)

	logger.Fatal(app.Listen(":3000"))

}
