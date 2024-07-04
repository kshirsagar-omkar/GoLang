package main

import (
	"fmt"
	"net/http"
	"log"

	//go get -u github.com/gorilla/mux
	//go mod tidy
	"github.com/gorilla/mux"
)

func main(){
	fmt.Println("Mod in Golang");
	greeter();

	r := mux.NewRouter();
    r.HandleFunc("/", serveHome).Methods("GET");

    // Run This As Server
    log.Fatal(http.ListenAndServe(":4000",r));
}


func greeter(){
	fmt.Println("Hey There Mod Users");
}

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome To Golang Series on YT</h1>"));
}