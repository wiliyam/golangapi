package guestModel

import (
	"context"
	"fmt"
	"golangapi/library/mongodb"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func InsertOne() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := mongodb.GetDb().Collection("guest").InsertOne(ctx, bson.D{})
}

func GetAll() {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cur, err := mongodb.GetDb().Collection("guest").Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(reflect.TypeOf(result))
		//	return result
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
