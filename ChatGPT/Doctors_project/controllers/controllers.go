package controllers

import(
	"github.com/kshirsagar-omkar/doctors_project/utils"
	"github.com/kshirsagar-omkar/doctors_project/models"
	//"github.com/kshirsagar-omkar/doctors_project/model"

	"fmt"
	"context"
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
    //"log"

    //"go.mongodb.org/mongo-driver/mongo/options"


)




//MOST IMPORTANT as Connection
var collection *mongo.Collection;

func init(){
	collection = utils.ConnectDB();
}







//insert one doctor
func insertOneRecord(doctor model.Doctor){
	inserted,err := collection.InsertOne(context.Background(), doctor);
	utils.CheckErr(err);

	fmt.Println("Inserted One Doctor In DataBase With ID :",inserted.InsertedID);
}

//update one record
func updateOneRecord(doctorId string, doctor model.Doctor){
	id,err := primitive.ObjectIDFromHex(doctorId);

	filter := bson.M{"_doctorid":id}; //grab doctorid from db

	update := bson.M{
		"$set":bson.M{
			"doctorid" 				: doctorId,
			"doctorname" 			: doctor.DoctorName,
			"doctorspecialist"		: doctor.DoctorSpecialist,
			"doctorqualification"	: doctor.DoctorQualification,
			"doctorsalary"			: doctor.DoctorSalary,
		},
	}

	result,err := collection.UpdateOne(context.Background(), filter, update);
	utils.CheckErr(err);

	fmt.Println("Modified Count :",result.ModifiedCount);
}


//Deleting Record
func deleteOneRecord(doctorId string){

	id,err := primitive.ObjectIDFromHex(doctorId);

	filter := bson.M{"_doctorid":id};

	deleteCount, err := collection.DeleteOne(context.Background(), filter);
	utils.CheckErr(err);

	fmt.Println("Record Of Docter Having Id :",deleteCount," is Deleted");
}

func deleteAllRecords() string{
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil);
	utils.CheckErr(err);
	return string(deleteResult.DeletedCount);
}


func getAllRecords() []primitive.M{
	cursor, err := collection.Find(context.Background(), bson.D{{}});
	utils.CheckErr(err);

	var doctors []primitive.M;

	for cursor.Next(context.Background()){
		var doctor bson.M;
		err := cursor.Decode(&doctor);
		utils.CheckErr(err);
		doctors = append(doctors, doctor);
	}
	defer cursor.Close(context.Background());
	return doctors;
}







//DATABASE

//get all movies from db
func GetAllRecords(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","GET");

	allRecords := getAllRecords();

	json.NewEncoder(w).Encode(allRecords);
}


//Create one Doctor Record
func InsertOneRecord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","POST");

	var doctor model.Doctor;

	_ = json.NewDecoder(r.Body).Decode(&doctor);

	//Check existance or handle this one propper

	insertOneRecord(doctor);

	json.NewEncoder(w).Encode(doctor);
}

//update One Record
func UpdateOneRecord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","PUT");

	params := mux.Vars(r);
	var doctor model.Doctor;
	_ = json.NewDecoder(r.Body).Decode(&doctor);

	updateOneRecord(params["doctorid"], doctor);

	json.NewEncoder(w).Encode(doctor);
}

//Delete Record/Records
func DeleteOneRecord(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","DELETE");

	params := mux.Vars(r);

	deleteOneRecord(params["doctorid"]);

	json.NewEncoder(w).Encode("Record Deleted Having doctorID : "+params["doctorid"]);
}

func DeleteAllRecords(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json");
	w.Header().Set("Allow-Control-Allow-Methods","DELETE");

	deleteCount := deleteAllRecords();

	json.NewEncoder(w).Encode("Total "+deleteCount+" Records Have Been Deleted");
}