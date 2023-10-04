package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "errors"
)

func healthcheck (c *gin.Context){
    c.IndentedJSON(http.StatusOK, gin.H{"status":"healthy."})
}

func main (){
    var router = gin.Default()
    router.GET("/health",healthcheck)
    router.Run("0.0.0.0:8080")
}

