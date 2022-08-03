package main

import (
	"log"
	"my_project/internal/project"
	"my_project/pkg/logging"
)

func main() {
	log.Print("logger initializing...")
	logger := logging.Log()
	logger.Info("main function")
	project.NewProject(logger)

}
