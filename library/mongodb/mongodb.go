package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

//Init is Function to Initializing mongodb database
func Init() {

	// log.Print(client)
	color.Magenta("Initializing mongodb database...")

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URL")))

	if err != nil {
		color.Red("[mongodb]:error during making new client")
		log.Print(err)
		os.Exit(2)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		color.Red("[mongodb]: error during client connection")
		log.Print(err)
		os.Exit(2)

	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		color.Red("[mongodb]: error during ping mongodb server no connection found on url %s", os.Getenv("MONGODB_URL"))
		log.Print(err)
		os.Exit(2)

	}

	//link database

	client.Database(os.Getenv("DATABASE"))

	log.Print(client)

	color.Green("mongodb database connected successfully...")
}

//GetDb is Function to return mongo client
func GetDb() *mongo.Client {
	if client != nil {
		return client
	}
	Init()
	return client
}
