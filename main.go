package main

import (
	"fmt"
	"net/http"
	"os"
)

// GetMessage - get request to display message to the GUI
func GetMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go Webserver that has a few cool modern tools\nGoLang(of course), Docker, Docker-Compose and Kubernetes\nenv: "+os.Getenv("ENV"))
}

// Define the route
func routes() {
	http.HandleFunc("/", GetMessage)
}

func main() {
	fmt.Println("Connecting......")
	routes()
	http.ListenAndServe(":3000", nil)
}
