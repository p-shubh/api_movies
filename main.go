package main

import (
	"encoding/csv"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run(":8086") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(r *gin.Engine) {

	// 1. GET /movies/year/:year ---> All movies released in the provided year
	// 2. GET /movies/rating/:rating --> all movies which have rotten tomatoes value higher than provided rating
	// 3. GET /movies/genre/:genre ---> All movies which have provided genre

	// r.GET("/city/:name", Dummy)
	r.GET("/movies/year/:year", dummyYear)
	r.GET("/movies/rating:rating", dummyRating)
	r.GET("/movies/genre:genre", dummyGenre)
}

func getYearMovies(year string, records [][]string) []string { //year vejenge aur movies ka total naam ayega

	var movies []string
	for i := 0; i < len(records); i++ {
		if records[i][7] == year {
			// movies = records[i][0]
			movies = append(movies, records[i][0])
			// movies = append(int[],movies)
		}
	}
	return movies
}

func getRatingMovies(rating string, records [][]string) []string { //year vejenge aur movies ka total naam ayega

	var ratings []string
	for i := 0; i < len(records); i++ {
		if records[i][5] == rating {
			// movies = records[i][0]
			ratings = append(ratings, records[i][0])
			// movies = append(int[],movies)
		}
	}
	return ratings
}

func getGenreMovies(Genre string, records [][]string) []string { //year vejenge aur movies ka total naam ayega

	var genre []string
	for i := 0; i < len(records); i++ {
		if records[i][7] == Genre {
			// movies = records[i][0]
			genre = append(genre, records[i][0])
			// movies = append(int[],movies)
		}
	}
	return genre
}

// Dummy function
// func Dummy(c *gin.Context) {
// 	name, ok := c.Params.Get("name")
// 	if ok == false {
// 		res := gin.H{
// 			"error": "name is missing",
// 		}
// 		c.JSON(http.StatusOK, res)
// 		return
// res := gin.H{
// 	"name": name,
// 	"city": "hajipur",
// }
// c.JSON(http.StatusOK, res)
// 	}

func dummyYear(c *gin.Context) {

	year, ok := c.Params.Get("year")
	records := readCsvFile("./movies.csv")
	movieName := getYearMovies(year, records)

	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res := gin.H{
		"year":   year,
		"movies": movieName,
	}
	c.JSON(http.StatusOK, res)
}

func dummyRating(c *gin.Context) {

	rating, ok := c.Params.Get("Rating")
	records := readCsvFile("./movie.csv")
	movieName := getRatingMovies(rating, records)

	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res := gin.H{
		"rating": rating,
		"movies": movieName,
	}
	c.JSON(http.StatusOK, res)
}

func dummyGenre(c *gin.Context) {

	Genre, ok := c.Params.Get("Genre")
	records := readCsvFile("./movie.csv")
	movieName := getGenreMovies(Genre, records)

	if ok == false {
		res := gin.H{
			"error": "name is missing",
		}
		c.JSON(http.StatusOK, res)
		return
	}
	res := gin.H{
		"Genre":  Genre,
		"movies": movieName,
	}
	c.JSON(http.StatusOK, res)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
