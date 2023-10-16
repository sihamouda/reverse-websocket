package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)


func connect()(*mongo.Client,context.Context,context.CancelFunc, error) {
	uri := os.Getenv("MONGODB_URI")
	if uri == ""{
		log.Fatal("no database uri provided!")
		return nil,nil,nil, errors.New("env var: missing db uri")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	return client, ctx ,cancel , err
}

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc){
	defer cancel()
	
	defer func(){
		if err := client.Disconnect(ctx); err != nil{
			panic(err)
		}
	}()
}

func ping(client *mongo.Client, ctx context.Context) error{
 
    if err := client.Ping(ctx, readpref.Primary()); err != nil {
        return err
    }
    fmt.Println("connected to database successfully!")
    return nil
}

func HandleConnection()(*mongo.Client,context.Context){ 
	client, ctx, cancel, err := connect()
    if err != nil{
        panic(err)
    }
    defer close(client, ctx, cancel)
    ping(client, ctx)
	return client, ctx
}