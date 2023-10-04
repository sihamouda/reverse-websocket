package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "errors"
)

func healthcheck (c *gin.Context){
    c.IndentedJSON(http.StatusOK, gin.H{"status":"healthy."})
}

func main (){
    port := os.Getenv("PORT")
    var router = gin.Default()
    router.GET("/health",healthcheck)
    router.Run("0.0.0.0:"+port)
}

