package router

import (
    "net/http"
    "path/filepath"

    "github.com/gorilla/mux"
    "github.com/kshirsagar-omkar/money_management/controller"
)

func serveHTML(w http.ResponseWriter, r *http.Request) {
    path, _ := filepath.Abs("./static/index.html")
    http.ServeFile(w, r, path)
}

func Router() *mux.Router {
    router := mux.NewRouter()

    // Create a record
    router.HandleFunc("/api/add", controller.CreateRecord).Methods("POST", "OPTIONS")

    // Get All Records
    router.HandleFunc("/api/records", controller.GetAllRecords).Methods("GET", "OPTIONS")

    // Delete One Record
    router.HandleFunc("/api/record/delete/{id}", controller.DeleteOneRecord).Methods("DELETE", "OPTIONS")

    // Update One Record
    router.HandleFunc("/api/record/update/{id}", controller.UpdateOneRecord).Methods("PUT", "OPTIONS")

    // Serve static files from the "static" folder
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    // Serve the HTML file
    router.HandleFunc("/", serveHTML).Methods("GET")

    return router
}