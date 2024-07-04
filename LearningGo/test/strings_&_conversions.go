package main
import(
	"fmt"
	"strings"
	"strconv"
)

func main(){

	{
		fmt.Println("Problem 1: String to Integer Conversion");
		strAge := "29";
		age,err := strconv.Atoi(strAge);
		if err != nil{
			fmt.Println(err,"\n");
		} else {
			fmt.Println(age,"\n");
		}
	}

	{
		fmt.Println("Problem 2: Integer to String Conversion");
		var p2 int64 = -123;
		s := strconv.FormatInt(p2,10);
		fmt.Println(s,"\n");
	}


	{
		fmt.Println("Problem 3: String to Float Conversion");
		s := "3.14";
		sToFloat,err := strconv.ParseFloat(s,64);
		if err != nil{
			fmt.Println(err,"\n");
		} else {
			fmt.Println(sToFloat,"\n"); 
		}
		
	}

	{
		fmt.Println("Problem 4: Float to String Conversion");
		var value = 3.14;
		s := strconv.FormatFloat(value, 'f', 2, 64);
		fmt.Println(s,"\n");
	}

	{
		fmt.Println("Problem 5: String Contains Substring");
		s1 := "Hello World";
		fmt.Println(strings.Contains(s1,"World"),"\n");
	}

	{
		fmt.Println("Problem 6: String Has Prefix \nWrite a function that checks if a string starts with a specified prefix. Use strings.HasPrefix.");
		fmt.Println("hello world", ":world\n",strings.HasPrefix("hello world", "world"));
		fmt.Println("hello world", ":hello\n",strings.HasPrefix("hello world", "hello"));
		fmt.Println();
	}

	{
		fmt.Println("Problem 6: String Has Suffix \nWrite a function that checks if a string ends with a specified suffix. Use strings.HasSuffix.");
		fmt.Println("hello world", ":world\n",strings.HasSuffix("hello world", "world"));
		fmt.Println("hello world", ":hello\n",strings.HasSuffix("hello world", "hello"));
		fmt.Println();
	}

	{
		fmt.Println("Problem 8: Repeat String\nWrite a function that repeats a string a specified number of times. Use strings.Repeat.");
		fmt.Println(strings.Repeat("Hello ",3));
		fmt.Println();
	}

	/*
		Problem 9: Replace All Substrings

		Write a function that replaces all occurrences of a substring within a string with another substring. Use strings.ReplaceAll.

		Function Signature:

		go

		func replaceAll(s, old, new string) string

		Example:

		go

		fmt.Println(replaceAll("hello world", "world", "golang"))  // Output: "hello golang"
		fmt.Println(replaceAll("foo bar foo", "foo", "baz"))       // Output: "baz bar baz"

		Problem 10: Trim Prefix

		Write a function that removes a specified prefix from a string. If the string does not start with the prefix, it returns the string unchanged. Use strings.TrimPrefix.

		Function Signature:

		go

		func trimPrefix(s, prefix string) string

		Example:

		go

		fmt.Println(trimPrefix("hello world", "hello"))  // Output: " world"
		fmt.Println(trimPrefix("hello world", "world"))  // Output: "hello world"

	*/





}