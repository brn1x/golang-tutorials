package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type album struct {
	ID     int64   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var albumsTest = []album{
	{ID: 1, Title: "My Everything", Artist: "Ariana Grande", Price: 56.99},
	{ID: 2, Title: "Sounds good feels good", Artist: "5 Seconds of Summer", Price: 17.99},
	{ID: 3, Title: "Midnight Memories", Artist: "One Direction", Price: 39.99},
}

func main() {
	databaseConnection()
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.GET("/albums/artist/:name", getAlbumsByArtistName)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func databaseConnection() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DBHOST"),
		DBName:               "recordings",
		AllowNativePasswords: true,
	}

	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connection with database stablished")
}

func getAlbums(context *gin.Context) {
	var albums []album
	rows, err := db.Query("SELECT id, title, artist, price FROM album")

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error on database query"})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var alb album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error formating response"})
			return
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {

	}

	context.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(context *gin.Context) {
	var newAlbum album

	if err := context.BindJSON(&newAlbum); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error on formating data into Album"})
		return
	}

	result, err := db.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error on saving data to database"})
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error on getting last inserted id on database"})
		return
	}

	newAlbum.ID = id

	context.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(context *gin.Context) {
	var album album

	id, err := strconv.ParseInt(context.Param("id"), 0, 32)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID is not a number"})
		return
	}

	row := db.QueryRow("SELECT id, title, artist, price FROM album WHERE id = ?", id)

	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
			return
		}

		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error on searching for data"})
	}

	context.IndentedJSON(http.StatusOK, album)
}

func getAlbumsByArtistName(context *gin.Context) {
	var albums []album

	artistName := context.Param("name")

	rows, err := db.Query("SELECT * FROM album WHERE artist LIKE ?", "%"+artistName+"%")

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error on database query"})
		return
	}

	rows.Close()

	for rows.Next() {
		var alb album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error on query"})
		}

		fmt.Println("TESTEEEEE -> ", alb.Title)
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		if err == sql.ErrNoRows {
			context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
			return
		}
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error on query"})
		return
	}

	context.IndentedJSON(http.StatusOK, albums)
}
