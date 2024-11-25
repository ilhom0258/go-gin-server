//server.go
package app

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"example.com/web-service-gin/models"
)


// albums slice to seed record album data.
var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) { 
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil { 
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	albums = append(albums, newAlbum)

	c.IndentedJSON(http.StatusCreated, newAlbum);
}

func getAlbumByID(c *gin.Context) {
	ID := c.Param("id");
	for _, v := range albums {
		if v.ID == ID {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	
	c.AbortWithStatus(http.StatusNotFound)
}
