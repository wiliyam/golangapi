package main

import (
	"context"
	"fmt"
	"golangapi/dotenv"
	"golangapi/library/mongodb"
	"log"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
)

type Dummy struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Title  string        `json:"title"`
	Title2 string        `json:"title2"`
}

func main() {
	//	port := dotenv.GetVariableValue("SERVER_PORT")
	color.Magenta("Initializing server...")

	//init env variales
	dotenv.Init()

	//init dabase
	mongodb.Init()

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// result, err := mongodb.GetDb().Collection("dummy").InsertOne(ctx, bson.D{
	// 	{"title", "dummmfsddfsd"},
	// 	{"tag", bson.A{"ccfsdfsdfdf", "nsdsd"}},
	// })

	resultData := []Dummy{}

	cur, err := mongodb.GetDb().Collection("dummy").Find(ctx, bson.D{})

	for cur.Next(ctx) {
		var dummy Dummy
		err := cur.Decode(&dummy)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(dummy)
		resultData = append(resultData, dummy)
		//	return result
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(resultData)

}
