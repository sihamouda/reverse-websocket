package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sihamouda/reverse-websocket/webserver/src/types"
	"go.mongodb.org/mongo-driver/bson"
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

func HandleConnection(channel chan *mongo.Client){ 
	client, ctx, cancel, err := connect()
    if err != nil{
        panic(err)
    }
    defer close(client, ctx, cancel)
    ping(client, ctx)
	channel <- client
}

func CreateWorker(client *mongo.Client, worker types.Worker) error{
	collection := client.Database("webserver").Collection("workers")

	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	
	_, err := collection.InsertOne(ctx, bson.M{"hostname": worker.Hostname})
	if err != nil { 
		return errors.New("db: worker could not be registred")
	}
	return nil
}

func ReadWorkers(client *mongo.Client) ([]bson.D, error){
	collection := client.Database("webserver").Collection("workers")

	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	
	cursor , err := collection.Find(ctx, bson.M{})
	if err != nil { 
		return nil ,errors.New("db: error while reading from db")
	}

	var results []bson.D

    if err := cursor.All(ctx, &results); err != nil {
        return nil ,errors.New("db: error while reading from db")
    }
	return results, nil
}