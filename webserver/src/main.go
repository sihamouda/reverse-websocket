package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

func healthcheck (c *gin.Context){
    c.IndentedJSON(http.StatusOK, gin.H{"status":"healthy."})
}

type worker struct {
    // ID string `json:"id"`
    IP string `json:"ip"`
}

var workers = []worker{}

func getWorkers(c *gin.Context){
    c.IndentedJSON(http.StatusOK, workers)
}

func registerWorker(c *gin.Context){
    var workerDTO worker

    if err := c.BindJSON(&workerDTO); err != nil {
        return
    }

    workers = append(workers, workerDTO)
    c.IndentedJSON(http.StatusCreated,workerDTO)
}

func main() {
    var router = gin.Default()
    router.GET("/health",healthcheck)
    router.GET("/worker",getWorkers)
    router.POST("/worker",registerWorker)
    router.Run("0.0.0.0:8080")
}