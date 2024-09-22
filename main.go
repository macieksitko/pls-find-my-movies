package main

import (
	"database/sql"
	"fmt"
	"log"
	"pls-find-my/models"
	"pls-find-my/views"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var db *sql.DB

func initDatabase(dbPath string) error {
	var err error
	db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := initDatabase("./movies.sqlite")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := gin.Default()
	r.Static("/static", "./static")

	r.GET("/", handleIndex)
	r.GET("/search", handleSearch)
	r.GET("/streamings/:id", handleStreamings)
	r.Run("localhost:8080")
}

func handleIndex(c *gin.Context) {
	component := views.Index()
	component.Render(c.Request.Context(), c.Writer)
}

func searchMovies(query string) ([]models.Movie, error) {
	time.Sleep(2 * time.Second)

	rows, err := db.Query("SELECT id, title, description FROM movies WHERE title LIKE ?", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie

		if err := rows.Scan(&m.ID, &m.Title, &m.Description); err != nil {
			return nil, err
		}

		movies = append(movies, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}

func handleSearch(c *gin.Context) {
	query := c.Query("q")
	results, err := searchMovies(query)
	if err != nil {
		c.String(500, "Error searching movies: %v", err)
		return
	}

	component := views.SearchResults(results)
	component.Render(c.Request.Context(), c.Writer)
}

func handleStreamings(c *gin.Context) {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "Invalid movie ID")
		return
	}

	fmt.Printf("Movie ID: %d\n", movieID)

	results, err := getMovieStreamings(movieID)
	if err != nil {
		c.String(500, "Error getting movie streaming details: %v", err)
		return
	}

	component := views.StreamingResults(results)
	component.Render(c.Request.Context(), c.Writer)
}

func getMovieStreamings(movieID int) ([]models.MovieStreaming, error) {
	time.Sleep(2 * time.Second)

	query := `
        SELECT ms.movie_id, ms.streaming_id, ms.movie_url, s.service
        FROM movie_streamings ms
        JOIN streamings s ON ms.streaming_id = s.id
        WHERE ms.movie_id = ?
    `
	rows, err := db.Query(query, movieID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []models.MovieStreaming
	for rows.Next() {
		var ms models.MovieStreaming
		err := rows.Scan(&ms.MovieID, &ms.StreamingID, &ms.MovieURL, &ms.Streaming.Service)
		if err != nil {
			return nil, err
		}
		results = append(results, ms)
	}

	return results, nil
}
