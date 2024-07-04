package main


import(
	"fmt"
	"net/http"
	"log"
	"github.com/kshirsagar-omkar/doctors_project/router"
)


func main(){
	fmt.Println("Doctors_Project\n");
	fmt.Println("Server Is Getting Started...");

	r := router.DoctorRouter();

	log.Fatal(http.ListenAndServe(":4008",r));
}