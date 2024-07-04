package router

import(
	"github.com/gorilla/mux"
	"github.com/kshirsagar-omkar/mongoApi/controller"
)

func Router() *mux.Router{
	router := mux.NewRouter();

	//Get ALl Movies
	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET");

	//Create a Movie
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST");

	//Update movie as marked
	router.HandleFunc("/api/movie/{id}", controller.MarkMovieAsMarked).Methods("PUT");

	//Delete one movie
	router.HandleFunc("/api/deletemovie/{id}", controller.DeleteOneMovie).Methods("DELETE");

	//Delete many movies
	router.HandleFunc("/api/deletemovies", controller.DeleteAllMovies).Methods("DELETE");

	return router;
}