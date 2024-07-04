package main
import(
	"fmt"
	"log"
	"net/http"
	"github.com/kshirsagar-omkar/mongoApi/router"
)


func main(){
	fmt.Println("MongoDB API");

	r := router.Router();

	fmt.Println("Server is getting Started...");

	log.Fatal(http.ListenAndServe(":4003",r));

	fmt.Println("Server Listining at Port :4003");
}