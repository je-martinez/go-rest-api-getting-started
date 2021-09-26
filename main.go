package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", putAlbums)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func putAlbums(c *gin.Context) {
	var modifiedAlbum album
	id := c.Param("id")
	if err := c.BindJSON(&modifiedAlbum); err != nil {
		return
	}
	_, index, err := _findByID(id)
	if err == "" {
		recordToModify := &albums[index]
		recordToModify.Title = modifiedAlbum.Title
		recordToModify.Artist = modifiedAlbum.Artist
		recordToModify.Price = modifiedAlbum.Price
		c.IndentedJSON(http.StatusOK, recordToModify)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})

	}
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	record, _, errr := _findByID(id)

	if errr == "" {
		c.IndentedJSON(http.StatusOK, record)
		return
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": errr})
	}
}

func _findByID(id string) (album, int, string) {
	for index, a := range albums {
		if a.ID == id {
			return a, index, ""
		}
	}
	return album{}, -1, "album not found"
}
