package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"
	"github.com/vanphuoc3012/bazel-example/projects/go_hello_world"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")
	w.Write([]byte(go_hello_world.HelloWorld("Phuoc") + " Bazoku is cool!!"))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return port
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", YourHandler)
	// Bind to a port and pass our router in
	port := getPort()
	log.Println("running program's operating system target: " + runtime.GOOS)
	log.Println("running program's architecture target: " + runtime.GOARCH)
	log.Println("Going to listen on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
