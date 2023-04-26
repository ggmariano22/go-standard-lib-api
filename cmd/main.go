package main

import (
	"api-std-packages/config"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.ConfigServer()

	fmt.Println("*** Server up and running at port :8080 ***")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

