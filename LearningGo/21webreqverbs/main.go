package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
)

func main(){
	fmt.Println("Title : Web Verb\n")

	//PerformGetRequest()
	//PerformPostJsonRequest()
	PerformPostFormRequest()

}



func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstName","omkar")
	data.Add("lastName","kshirsagar")
	data.Add("email","omkar@go.dev")

	response,err := http.PostForm(myurl,data)
	checkErr(err)

	content,err := ioutil.ReadAll(response.Body)
	checkErr(err)

	fmt.Println(string(content))
}


func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"CourseName":"go",
			"price":0,
			"platform":"youtube"
		}
	`)

	response,err := http.Post(myurl,"application/json",requestBody)
	checkErr(err)

	defer response.Body.Close()

	content,err := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformGetRequest(){
	const myurl = "http://localhost:8000/gets"

	response,err := http.Get(myurl)
	checkErr(err)

	defer response.Body.Close()

	fmt.Println("Statues code :",response.StatusCode,"\n")
	//if else (if status code is not 200:ok )
	
	//fmt.Println("Content Lenght is :",response.ContentLength,"\n")

	//read response of the body

	content,err := ioutil.ReadAll(response.Body)
	checkErr(err)

	var responseString strings.Builder
	byteCount,err := responseString.Write(content)
	checkErr(err)

	fmt.Println("Content Lenght is :",byteCount,"\n")

	fmt.Println(responseString.String())

	//fmt.Println(string(content),"\n")
}

func checkErr(err error)(){
	if err != nil {
		panic(err)
	}
}