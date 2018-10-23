package main

import (
	"ctrl"
	"log"
	"net/http"
	"os"
)

func main() {
	// log everything to /tmp/server.log file
	logfile, err := os.OpenFile("/tmp/server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	// initialise controllers
	ctrl.Setup()

	// start the server
	log.Fatal(http.ListenAndServe(":8288", nil))
}
