package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/naman1402/eigen-bootcamp-assignment-2/Execution_Service/utils" 
	"github.com/naman1402/eigen-bootcamp-assignment-2/Execution_Service/service" 
)

func main() {
	services.Init()
	router := gin.Default()
	router.POST("/task/execute", utils.ExecuteTask)
	log.Println("Server starting on :4003")

	if err := router.Run(":4003"); err != nil {
		log.Fatal(err)
	}
}