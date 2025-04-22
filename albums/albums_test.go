package albums

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestGetAllAlbums(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums", main.getAlbums)

	req, err := http.NewRequest("GET", "/albums", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var rcvd []main.album
	json.Unmarshal(w.Body.Bytes(), &rcvd)

	if status := w.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: wanted %v got %v",
			status, http.StatusOK)
	}

	if !reflect.DeepEqual(rcvd, main.albums) {
		t.Fatalf("handler return unexpected body: got %v want %v",
			rcvd, main.albums)
	}
}

func TestGetSingleAlbum(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums/:id", main.getAlbumByID)

	req, err := http.NewRequest("GET", "/albums/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var rcvd main.album
	json.Unmarshal(w.Body.Bytes(), &rcvd)
	if status := w.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: wanted %v got %v",
			http.StatusOK, status)
	}
	if main.albums[1] != rcvd {
		t.Fatalf("handler return unexpected body: got %v want %v",
			rcvd, main.albums[1])
	}
}

func TestPostAlbum(t *testing.T) {
	r := SetUpRouter()
	r.POST("/albums", main.postAlbums)
	originalAlbumsLen := len(main.albums)

	var newAlbum = main.album{
		ID:     "4",
		Title:  "Next Level Foo",
		Artist: "The Bars",
		Price:  999.99,
	}

	jsonData, error := json.Marshal(newAlbum)
	if error != nil {
		t.Fatal(error)
	}

	req, err := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if len(main.albums) != originalAlbumsLen+1 {
		t.Fatal("Something went wrong: new album not added")
	}

	if main.albums[originalAlbumsLen] != newAlbum {
		t.Fatalf("expected new album to equal %v", newAlbum)
	}
}
