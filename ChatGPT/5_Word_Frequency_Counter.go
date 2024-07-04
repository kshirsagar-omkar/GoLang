package main

import(
	"fmt"
	"strings"

)


/*
	Problem 5: Word Frequency Counter
	Write a function that counts the frequency of each word in a given string and returns a map with words as keys and their frequencies as values.

	Function Signature:
	func wordFrequency(text string) map[string]int

	Example:
	fmt.Println(wordFrequency("hello world hello"))  // Output: map[hello:2 world:1]
	fmt.Println(wordFrequency("go go golang"))       // Output: map[go:2 golang:1]
*/


func main(){

	fmt.Println("5: Word Frequency Counter");
	fmt.Println(wordFrequency("hello world hello"));
	fmt.Println(wordFrequency("go go golang lang go"))   
}

func wordFrequency(str string) map[string]int {
	var s = []string{}
	s = strings.Split(str," ");

	freq := make(map[string]int)

	for i := range s{
		freq[s[i]] += 1;
	}

	return freq;
}