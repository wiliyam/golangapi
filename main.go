package main

import (
	"golangapi/dotenv"
	"golangapi/library/mongodb"
	"golangapi/web"

	"github.com/fatih/color"
)

func main() {
	//	port := dotenv.GetVariableValue("SERVER_PORT")
	color.Magenta("Initializing server...")

	//init env variales
	dotenv.Init()

	//init dabase
	mongodb.Init()

	//init server

	web.InitServer()

}
