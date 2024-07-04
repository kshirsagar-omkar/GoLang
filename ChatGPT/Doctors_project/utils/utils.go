// package utils
// import (
// 	"fmt"
//     "context"
//     "log"
//     "go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
// )
// const connectionString = "mongodb+srv://omkarkshirsagar:1bUpJ1sXN5tFMial@mydb.pyncv4c.mongodb.net/?retryWrites=true&w=majority&appName=mydb";
// const dbName = "hospital";
// const collectionName = "doctor";


// var collection *mongo.Collection;


// func Init() {
// 	//client option
// 	clientOption := options.Client().ApplyURI(connectionString);

// 	//Connect to mongoDB
// 	client, err := mongo.Connect(context.TODO(), clientOption);
// 	CheckErr(err);
// 	fmt.Println("Connection Successfull\n");

// 	collection = client.Database(dbName).Collection(collectionName);

// }

// func GetDB() *mongo.Collection {
// 	return collection;
// }

// func CheckErr(err error){
// 	log.Fatal(err);
// }


// utils/database.go
package utils


import (
	"fmt"
    "context"
    "log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)


const connectionString = "mongodb+srv://omkarkshirsagar:1bUpJ1sXN5tFMial@mydb.pyncv4c.mongodb.net/?retryWrites=true&w=majority&appName=mydb";
const dbName = "hospital";

const colName = "doctors";


//MOST IMPORTANT as Connection
var collection *mongo.Collection





//Connect Wit MongoDB
//init works first tiem and only one time as singleton class of cpp
func ConnectDB() *mongo.Collection{
	
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

	return collection;
}



func CheckErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
