//router.go
package app

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine { 
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	return router
} 