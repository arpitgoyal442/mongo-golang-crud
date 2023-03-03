package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/mongo-golang-hitesh/router"

)

func main(){

	fmt.Println("Server getting started.....")
	log.Fatal(http.ListenAndServe("localhost:8080",router.MyRouter()))
	fmt.Println("Server is Running")
}
