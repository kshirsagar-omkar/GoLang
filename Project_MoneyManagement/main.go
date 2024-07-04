package main

import (
	"fmt"
	"log"
	"net/http"



	"github.com/kshirsagar-omkar/money_management/router"
)

func main() {
	fmt.Println("Money Management\n")

	r := router.Router()

	fmt.Println("Server is getting Started...")

	log.Fatal(http.ListenAndServe(":4001", r))
}

