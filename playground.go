package playground

import (
	"context"
	"fmt"
	"golangapi/dotenv"
	"golangapi/library/mongodb"
	"log"
	"time"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Dummy struct {
	ID     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title  string             `json:"title"`
	Title2 string             `json:"title2"`
}

type GuestModel struct {
	ID         primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	DeviceId   string             `json:"deviceId"`
	UserGender string             `json:"userGender"`
}

func getData() []GuestModel {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// result, err := mongodb.GetDb().Collection("dummy").InsertOne(ctx, bson.D{
	// 	{"title", "dummmfsddfsd"},
	// 	{"tag", bson.A{"ccfsdfsdfdf", "nsdsd"}},
	// })

	resultData := []GuestModel{}

	cur, err := mongodb.GetDb().Collection("guest").Find(ctx, bson.D{})

	for cur.Next(ctx) {
		var dummy GuestModel
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

	return resultData
}

func main1() {
	//	port := dotenv.GetVariableValue("SERVER_PORT")
	color.Magenta("Initializing server...")

	//init env variales
	dotenv.Init()

	//init dabase
	mongodb.Init()

	result := getData()

	fmt.Println("--result->", result)

}
