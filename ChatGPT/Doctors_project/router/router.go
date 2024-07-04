package router

import(
	"github.com/kshirsagar-omkar/doctors_project/controllers"
    "github.com/gorilla/mux"
)

func DoctorRouter() *mux.Router{
	r := mux.NewRouter();

	//All Doctors Records
	r.HandleFunc("/api/doctors", controllers.GetAllRecords).Methods("GET");

	//insert a doctor Record
	r.HandleFunc("/api/doctor", controllers.InsertOneRecord).Methods("POST");

	//Update one record
	r.HandleFunc("/api/doctor/{doctorid}", controllers.UpdateOneRecord).Methods("PUT");

	//Delete one Record
	r.HandleFunc("/api/deletedoctor/{doctorid}", controllers.DeleteOneRecord).Methods("DELETE");

	//Delete many record
	r.HandleFunc("/api/deletedoctors", controllers.DeleteAllRecords).Methods("DELETE");

	return r;

}