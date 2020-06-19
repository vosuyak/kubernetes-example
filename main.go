package main

import (
	"fmt"
	"net/http"
	"os"

	"context"
	"log"
	"time"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	LoadEnv()
}

// GetMessage - get request to display message to the GUI
func GetMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Webserver that has a few cool modern tools\nGoLang(of course), Docker, Docker-Compose and Kubernetes\nenv: "+os.Getenv("ENV"))
}

// Define the route
func routes() {
	http.HandleFunc("/", GetMessage)
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	initData()
}

func initData() {
	connecturi := fmt.Sprintf(
		"mongodb://%s:%s",
		os.Getenv("MONGO_DB_HOST"),
		os.Getenv("MONGO_DB_PORT"),
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(connecturi))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	quickstartDatabase := client.Database("quickstart")
	fmt.Println(quickstartDatabase)
}

func main() {
	fmt.Println("Connecting......", os.Getenv("MONGO_DB_HOST"))
	routes()
	http.ListenAndServe(":3000", nil)
}
