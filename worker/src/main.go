package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	// "errors"
)

func healthcheck (c *gin.Context){
    c.IndentedJSON(http.StatusOK, gin.H{"status":"healthy."})
}

func registerWorker(){
    var webServerDomain = os.Getenv("WEBSERVER")
    url := fmt.Sprintf("http://%s/worker", webServerDomain)

    hostname , err := os.Hostname() 
    if err != nil {
        log.Println(err)
        return
    }
    jsonBodyString := fmt.Sprintf(`{"hostname": "%s"}`,hostname)
    jsonBody := []byte(jsonBodyString)
    bodyReader := bytes.NewReader(jsonBody)
    
    res , err := http.Post(url ,"application/json" ,bodyReader)
    if err != nil {
        log.Println(err)
        return
    }
    defer res.Body.Close()
    log.Println("Registrated to Webserver with status code"+ res.Status)
    
}

func main (){
    registerWorker()
    port := os.Getenv("PORT")
    var router = gin.Default()
    router.GET("/health",healthcheck)
    router.Run("0.0.0.0:"+port)
}

