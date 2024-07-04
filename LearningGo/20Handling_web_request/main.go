package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const Url = "https://zodogo.com/guimp.com/"

func main(){
	fmt.Println("Welcome to my webpage");

	response, err := http.Get(Url);
	checkErr(err);

	defer response.Body.Close();

	databyte, err := ioutil.ReadAll(response.Body);
	checkErr(err);

	content := string(databyte);

	fmt.Println(content)

}


func checkErr(err error) {
	if err != nil{
		panic(err);
	}
}