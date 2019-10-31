package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Movie movie
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{
		Title: "A", Year: 1990, Color: false, Actors: []string{"o1", "o2"},
	},
	{
		Title: "C", Year: 1991, Color: true, Actors: []string{"c1", "c2"},
	},
	{
		Title: "B", Year: 1992, Color: true, Actors: []string{"d1", "d2"},
	},
}

var jsonMovies = `[
    {
        "Title": "Casablanca",
        "released": 1942,
        "Actors": [
            "Humphrey Bogart",
            "Ingrid Bergman"
        ]
    },
    {
        "Title": "Cool Hand Luke",
        "released": 1967,
        "color": true,
        "Actors": [
            "Paul Newman"
        ]
    },
    {
        "Title": "Bullitt",
        "released": 1968,
        "color": true,
        "Actors": [
            "Steve McQueen",
            "Jacqueline Bisset"
        ]
    }
]`

func main() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))

	data, err = json.MarshalIndent(movies, "", "\t")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(data))

	var unMovies []Movie
	if err = json.Unmarshal([]byte(jsonMovies), &unMovies); err != nil {
		log.Fatalf("JSON unmarshaling failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(unMovies)
}
