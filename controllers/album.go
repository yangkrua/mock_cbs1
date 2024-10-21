package controllers

import (
	_ "mock_cbs1/docs"
	"mock_cbs1/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "The flood", Artist: "Of Mice and Men", Price: 19.45},
	{ID: "2", Title: "Infest", Artist: "Papa Roach", Price: 20.50},
	{ID: "3", Title: "Meteora", Artist: "Linkin Park", Price: 22.99},
}

// GetAlbums
//
//	@Summary		GetAlbums
//	@Description	GetAlbums
//	@Tags			GetAlbums
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	interface{}
//	@Failure		404		{object}	interface{}
//	@Failure		422		{object}	interface{}
//	@Failure		500		{object}	interface{}
//	@Router			/albums [get]
//
// @Security BearerAuth
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum models.Album

	err := c.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
