package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}
type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
 w.Header().Set("Content-Type","application/json")
 json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	id := params["id"]
	for index, movie := range movies {
		if movie.ID == id {
			// Remove the movie by slicing the array
			movies = append(movies[:index], movies[index+1:]...)

			// Respond with a success message
			fmt.Fprintf(w, "Movie with ID %s has been deleted\n", id)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)



}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' from the URL parameters

	params := mux.Vars(r)
	id := params["id"]

	// Loop through the movies slice to find the movie with the specified ID
	for _, movie := range movies {
		if movie.ID == id {
			// If movie is found, return it as JSON
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

	// If movie not found, return an error response with a 404 status
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Movie with ID %s not found", id)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Get the 'id' from the URL parameters
	params := mux.Vars(r)
	id := params["id"]

	// Find the movie in the slice
	for index, movie := range movies {
		if movie.ID == id {
			// Movie found, we need to update it

			// Create a temporary movie struct to hold the new data
			var updatedMovie Movie
			// Decode the JSON body into the updatedMovie struct
			if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid input data: %v", err)
				return
			}

			// Update the movie in the slice
			movies[index] = updatedMovie

			// Return the updated movie as JSON
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}

	// If the movie with the given id was not found, return 404
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Movie with ID %s not found", id)
}
func createMovie(w http.ResponseWriter, r *http.Request) {
	// Create a new movie struct to hold the data from the request body
	var newMovie Movie
	
	// Decode the JSON request body into the newMovie struct
	if err := json.NewDecoder(r.Body).Decode(&newMovie); err != nil {
		// If there's an error decoding, return a 400 Bad Request
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid input data: %v", err)
		return
	}

	// Assign a new ID to the movie (this is a simple example; you could implement a better ID generator)
	newMovie.ID = fmt.Sprintf("%d", len(movies)+1)

	// Add the new movie to the movies slice
	movies = append(movies, newMovie)

	// Return the newly created movie as JSON with a 201 Created status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMovie)
}

func main(){

	r:=mux.NewRouter()

	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "123456789",
		Title: "The Shawshank Redemption",
		Director: &Director{
			Firstname: "Frank",
			Lastname:  "Darabont",
		},
	})
	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "987654321",
		Title: "The Godfather",
		Director: &Director{
			Firstname: "Francis",
			Lastname:  "Coppola",
		},
	})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

fmt.Printf("Starting Server at port 8000\n")
log.Fatal(http.ListenAndServe(":8000",r))


}