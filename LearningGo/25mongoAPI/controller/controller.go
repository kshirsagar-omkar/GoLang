package controller


// import(
// 	"fmt"
// 	"log"
// 	"context"

// 	//"github.com/kshirsagar-omkar/mongoApi"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypto/options"
// )

import(
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kshirsagar-omkar/mongoApi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




const connectionString = "mongodb+srv://omkarkshirsagar:1bUpJ1sXN5tFMial@mydb.pyncv4c.mongodb.net/?retryWrites=true&w=majority&appName=mydb";
const dbName = "netflix";

const colName = "watchlist";


//MOST IMPORTANT as Connection
var collection *mongo.Collection





//Connect Wit MongoDB
//init works first tiem and only one time as singleton class of cpp
func init(){
	
	//Client Option
	clientOption := options.Client().ApplyURI(connectionString);

	//Connection to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption);
	if err != nil{
		log.Fatal(err);
	}
	fmt.Println("Connect To DataBase Succesfully\n");

	collection = client.Database(dbName).Collection(colName);
	//Collection instance / object
	fmt.Println("Collection Instance/Object is Ready");
}





//MONGODB Helpers - files

//Inser 1 Record

func insertOneMovie(movie model.Netflix){
	inserted, err := collection.InsertOne(context.Background(), movie);
	if err != nil{
		log.Fatal(err);
	}

	fmt.Println("Inserted 1 Movie in DataBase With ID :",inserted.InsertedID);
}   



//Updat 1 Record
func updateOneMovie(movieId string){
	//conver movieId as acceptible for MongoDB
	id,err := primitive.ObjectIDFromHex(movieId);
	if err != nil{
		log.Fatal(err);
	} 

	//id is used to run through all record

	filter := bson.M{"_id":id};		//Grab Id/Filter For you
	update := bson.M{"$set":bson.M{"watched":true}};


	result,err := collection.UpdateOne(context.Background(), filter, update);
	if err != nil{
		log.Fatal(err);
	} 

	fmt.Println("Modified Count :",result.ModifiedCount);
}



//delete 1 or many Record
//Delete 1 Movie
func deleteOneMovie(movieId string){

	id, err := primitive.ObjectIDFromHex(movieId);
	if err != nil{
		log.Fatal(err);
	}

	filter := bson.M{"_id":id};

	//returns only int value
	deleteCount, err := collection.DeleteOne(context.Background(), filter);
	if err != nil{
		log.Fatal(err);
	}

	fmt.Println("Delete Count of Movie is :",deleteCount);
}

//Delete all movies
func deleteAllMovie() int64{

	//filter := bson.D{{}}; 		//{} manje sagle records select kr
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil);			// both r ok () or (nil)
	if err != nil{
		log.Fatal(err);
	}
	fmt.Println("Delete Count of Movie is :",deleteResult.DeletedCount);
	return deleteResult.DeletedCount;
}



//Read Operation //Get All Movies From Database
func getAllMovies() []primitive.M {
	cursor, err := collection.Find(context.Background(), bson.D{{}});
	if err != nil{
		log.Fatal(err);
	}

	var movies []primitive.M;

	for cursor.Next(context.Background()){
		var movie bson.M;
		if err = cursor.Decode(&movie); err != nil{
			log.Fatal(err);
		}

		movies = append(movies, movie);
	}
	defer cursor.Close(context.Background());
	return movies;
}





//Actual Controller  - file
//Uper case first letter

//Get All Movies From DB
func GetAllMovies(w http.ResponseWriter, r *http.Request){
							   //application/json"
	w.Header().Set("Content-Type","application/x-www-form-urlencoded");

	allMovies := getAllMovies();

	json.NewEncoder(w).Encode(allMovies);
}


//Add Movie to DB
func CreateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded");

	//Accept only POST Method
	w.Header().Set("Allow-Control-Allow-Methods","POST");

	var movie model.Netflix;
	_ = json.NewDecoder(r.Body).Decode(&movie);

	insertOneMovie(movie);

	json.NewEncoder(w).Encode(movie);
}

//Mark Movie As Watched	//Update
func MarkMovieAsMarked(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded");

	//Accept only POST Method
	w.Header().Set("Allow-Control-Allow-Methods","PUT");

	params := mux.Vars(r);

	updateOneMovie(params["id"]);

	json.NewEncoder(w).Encode("Movie With id "+params["id"]+" Marked As Watched");
}



//DELETE 1
func DeleteOneMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded");

	//Accept only POST Method
	w.Header().Set("Allow-Control-Allow-Methods","DELETE");

	params := mux.Vars(r);

	deleteOneMovie(params["id"]);

	json.NewEncoder(w).Encode("Movie With id "+params["id"]+" Removed Succesfully");
}

//DELETE MANY
func DeleteAllMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/x-www-form-urlencoded");

	//Accept only POST Method
	w.Header().Set("Allow-Control-Allow-Methods","DELETE");

	deleteCount := string(deleteAllMovie());

	json.NewEncoder(w).Encode(deleteCount +"Movies Removed Succesfully");
}