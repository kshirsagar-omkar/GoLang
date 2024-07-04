package main

import(
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"time"
	"log"

	"github.com/gorilla/mux"
)


// Model for course - file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"courseprice"`
	Author *Author `json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}



//Creating Fake DB
var courses []Course;





//Middlewere, Helper - file
//Check is empty or not
func (c *Course) IsEmpty() bool {
	//return c.CourseId=="" && c.CourseName=="";
	return c.CourseName=="";
}


//Sending a API json response for all course - file
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API home Page</h1>"));
}     


//sending the API response for all courses
func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get All Courses");

	//How To Set Headers there are to many go through documentation
								//Key Value Pairs
	w.Header().Set("Content-Type","application/json");

	//Throw Entire Slice of course up here
	json.NewEncoder(w).Encode(courses);
}


//Get One Course Based on Request ID
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Courses");

	//How To Set Headers there are to many go through documentation
								//Key Value Pairs
	w.Header().Set("Content-Type","application/json");

	//Grab Id From Request
	//It is key value pairs  //Probabily in future we can get more data to grab
	
	params := mux.Vars(r);
	//All The Variables Are Comming In This AND This Will Extracted From Request That is Passed On

	fmt.Printf("Type Of params is : %T\n",params);
	fmt.Println(params);


	//Loop Through the course and find matching id and return it
	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course);
			return;
		}
	}
	json.NewEncoder(w).Encode("No Course Found Of Given CourseId : "+params["id"]);
	return
}


//Add a Create course Controller
func createOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Get One Courses");

	//How To Set Headers there are to many go through documentation
								//Key Value Pairs
	w.Header().Set("Content-Type","application/json");

	//what if body is empty //nobadu sends a data
	if r.Body == nil{
		json.NewEncoder(w).Encode("Please Send Some Data!!");
	}

	// what if data is in this - {} form
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course);
	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No Data Inside JSON!!");	
		return;
	}

	//Checking if same courseName is existing 
	for _,test_course := range courses{
		if  course.CourseName == test_course.CourseName {
			json.NewEncoder(w).Encode("Dupblicate Is Present in DataBase");
			return;
		}
	} 

	//Generate Unique Id,strin
	//Append course into courses

	rand.Seed(time.Now().UnixNano());
	course.CourseId = strconv.Itoa(rand.Intn(100));

	courses = append(courses, course);

	json.NewEncoder(w).Encode(course);
	return
}


//Updade a Course Controller
func updateOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Update One Courses");

	//How To Set Headers there are to many go through documentation
								//Key Value Pairs
	w.Header().Set("Content-Type","application/json");

	//First Grab The Id From Request
	params := mux.Vars(r);

	//Loop through it and find the course
	//Remove the Course
	//Update the new course on request id
	for i,course := range courses{
		if course.CourseId == params["id"]{
			//This deletes the course
			courses = append(courses[:i],courses[i+1:]...);

			//get new course detials from request
			var course Course;
			_ = json.NewDecoder(r.Body).Decode(&course);

			//TODO : if data is empty or {} handle this

			course.CourseId = params["id"];
			//Update course Dtials
			courses = append(courses, course);
			json.NewEncoder(w).Encode(course);
			return;
		}
	}
	//if Id not found
	json.NewEncoder(w).Encode("Id Nod Found In DataBase!!");
}


//Delete a Course Controller
func deleteOneCourse(w http.ResponseWriter, r *http.Request){

	fmt.Println("Delete One Courses");

	//How To Set Headers there are to many go through documentation
								//Key Value Pairs
	w.Header().Set("Content-Type","application/json");

	//First Grab The Id From Request
	params := mux.Vars(r);

	//Loop through it and find the course
	//Remove the Course
	for i,course := range courses{
		if course.CourseId == params["id"]{
			//This deletes the course
			courses = append(courses[:i],courses[i+1:]...);
			json.NewEncoder(w).Encode("Course Has Be Removed Succesfully");
			return;
		}
	}
	//if Id not found
	json.NewEncoder(w).Encode("Id Nod Found In DataBase!!");
}



func main(){
	fmt.Println("Building API");

	r := mux.NewRouter();

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	//Routing
	r.HandleFunc("/",serveHome).Methods("GET");
	r.HandleFunc("/courses",getAllCourses).Methods("GET");
					//params["id"] same name should both side
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET");
	r.HandleFunc("/course",createOneCourse).Methods("POST");
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT");
	r.HandleFunc("/course/{id}",deleteOneCourse).Methods("DELETE");



	//Listen to a port
	log.Fatal(http.ListenAndServe(":4001",r));


}

