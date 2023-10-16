package main

import (
	"net/http"
	"os"

	"github.com/sihamouda/reverse-websocket/webserver/src/api"
	"github.com/sihamouda/reverse-websocket/webserver/src/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func healthcheck (c *gin.Context){
    c.IndentedJSON(http.StatusOK, gin.H{"status":"healthy"})
}

func main() {
    dbCh := make(chan *mongo.Client , 1)

    go db.HandleConnection(dbCh)

    dbClient := <- dbCh

    port := os.Getenv("PORT")
    var router = gin.Default()
    router.GET("/health",healthcheck)
    router.GET("/worker",api.GetWorkers(dbClient))
    router.POST("/worker",api.RegisterWorker(dbClient))
    router.Run("0.0.0.0:"+port)
}