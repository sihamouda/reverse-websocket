package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sihamouda/reverse-websocket/webserver/src/db"
	"github.com/sihamouda/reverse-websocket/webserver/src/types"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetWorkers(dbClient *mongo.Client) func (c *gin.Context){
    return func (c *gin.Context){ 

        workers , err := db.ReadWorkers(dbClient)
            
        if err != nil {
            return
        }

        println(workers)
        c.IndentedJSON(http.StatusOK, workers)
    }
}

func RegisterWorker(dbClient *mongo.Client) func(c *gin.Context){
    return func(c *gin.Context){
        var workerDTO types.Worker

        if err := c.BindJSON(&workerDTO); err != nil {
            return
        }

        err := db.CreateWorker(dbClient,workerDTO)
        
        if err != nil {
            return
        }

        c.IndentedJSON(http.StatusCreated,workerDTO)
}
}
