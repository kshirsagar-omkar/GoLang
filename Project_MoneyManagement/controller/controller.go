package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"time"
	//"strconv"

	"github.com/gorilla/mux"
	"github.com/kshirsagar-omkar/money_management/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"


)

const connectionString = "mongodb+srv://omkarkshirsagar:1bUpJ1sXN5tFMial@mydb.pyncv4c.mongodb.net/?retryWrites=true&w=majority&appName=mydb"
const dbName = "MoneyManagement"
const colName = "testing"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}









func insertOneRecord(record model.MoneyManagement) bool {

	//record.Time = time.Now().Format("2006-01-02- 15:04:05");

	if record.Tag !="" && record.ForWhat != "" && record.Money != "" && record.Time != "" {
		inserted, err := collection.InsertOne(context.Background(), record)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single record: ", inserted.InsertedID)
		return true;
	}
	return false;
}


func updateOneRecord(id string, record model.MoneyManagement) bool {

	objID, err := primitive.ObjectIDFromHex(id);
	if err != nil{
		fmt.Println("Error Converting Id to objID");
		log.Fatal(err);
	}

	filter := bson.M{"_id":objID};

	update := bson.M{
	    "$set": bson.M{
	        "tag":     record.Tag,
	        "forwhat": record.ForWhat,
	        "money":   record.Money,
	        "time":    record.Time,
	    },
	}

	if record.Tag !="" && record.ForWhat != "" && record.Money != "" && record.Time != "" {

		result,err := collection.UpdateOne(context.Background(), filter, update);
		if err != nil{
			log.Fatal(err);
		} 
		fmt.Println("Modified Count :",result.ModifiedCount);
		return true;
	}
	return false;
}



func getAllRecords() []primitive.M {

	// // opts := options.Find()
	// // opts.SetSort(bson.D{{"time", 1}}) // Sort by 'time' field in ascending order

	 opts := options.Find()
	 opts.SetSort(bson.D{{"time", -1}}) // Sort by 'time' field in descending order



	// Define options for the find operation with a descending sort on the _id field
	// opts := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}})


	cursor, err := collection.Find(context.Background(), bson.D{{}}, opts);
	if err != nil{
		log.Fatal(err);
	}

	var records []primitive.M;

	for cursor.Next(context.Background()){
		var record bson.M;
		if err = cursor.Decode(&record); err != nil{
			log.Fatal(err);
		}

		records = append(records, record);
	}


	defer cursor.Close(context.Background());
	return records;
}



func deleteOneRecord(recordID string) {
	// Convert movie ID suitable to db
	id, err := primitive.ObjectIDFromHex(recordID)
	if err != nil {
		log.Fatalf("Error converting ID to ObjectID: %v", err)
	}

	// ID is used to filter and delete the record
	filter := bson.M{"_id": id}

	delCnt, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatalf("Error deleting record: %v", err)
	}

	fmt.Println("Delete count of record is:", delCnt.DeletedCount)
}



















func DeleteOneRecord(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle preflight OPTIONS request
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Ensure the request method is DELETE
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Get the ID parameter from the request URL
    params := mux.Vars(r)
    deleteOneRecord(params["id"]);
    json.NewEncoder(w).Encode("Record With id "+params["id"]+" Removed Succesfully");
}



func CreateRecord(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle preflight OPTIONS request
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Ensure the request method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Process POST request
    var record model.MoneyManagement
    if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    if insertOneRecord(record){
	    // Respond with the created record
	    w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	    json.NewEncoder(w).Encode(record)
	    return;
	}else{
		w.WriteHeader(http.StatusNotFound) // HTTP 404 not found
	    //json.NewEncoder(w).Encode("Record Not Added")
	}

}



func UpdateOneRecord(w http.ResponseWriter, r *http.Request) {
    // Set CORS headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    // Handle preflight OPTIONS request
    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // Ensure the request method is POST
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Process POST request
    var record model.MoneyManagement
    if err := json.NewDecoder(r.Body).Decode(&record); err != nil {
        http.Error(w, "Bad request", http.StatusBadRequest)
        return
    }

    params := mux.Vars(r);
    if updateOneRecord(params["id"], record){
	    // Respond with the created record
	    w.WriteHeader(http.StatusCreated) // HTTP 201 Created
	    json.NewEncoder(w).Encode(record)
	    return;
	}else{
		w.WriteHeader(http.StatusNotFound) // HTTP 404 not found
	    //json.NewEncoder(w).Encode("Record Not Added")
	}

}



func GetAllRecords(w http.ResponseWriter, r *http.Request){
							   //application/json"
	//w.Header().Set("Content-Type","application/json");


	// Set CORS headers
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	allRecords := getAllRecords();

	json.NewEncoder(w).Encode(allRecords);
}





