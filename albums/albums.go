package albums

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-server/docs"
	"net/http"
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

// @title Albums API
// @version 0.0.1
// @description.markdown
// @BasePath /api/v1
func Run() {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	rte := router.Group("/albums")
	{
		rte.GET("", getAlbums)
		rte.GET(":id", getAlbumByID)
		rte.POST("", postAlbums)
	}

	router.Run(":8080")
}

// postAlbums adds an album from JSON received in the request body.
// @Summary Add an album
// @Tags albums
// @Accept json
// @Produce json
// @Success 201
// @Router /albums  [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// return a list of all albums
// @Summary Return all albums
// @Tags albums
// @Produce json
// @Success 200
// @Router /albums [get]
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
// @Summary Get a single album
// @Tags albums
// @Produce json
// @Success 200
// @Router /albums/:id  [get]
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
