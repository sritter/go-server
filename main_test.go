package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
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
	r.GET("/albums", getAlbums)

	req, err := http.NewRequest("GET", "/albums", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var rcvd []album
	json.Unmarshal(w.Body.Bytes(), &rcvd)

	if status := w.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: wanted %v got %v",
			status, http.StatusOK)
	}

	if !reflect.DeepEqual(rcvd, albums) {
		t.Fatalf("handler return unexpected body: got %v want %v",
			rcvd, albums)
	}
}

func TestGetSingleAlbum(t *testing.T) {
	r := SetUpRouter()
	r.GET("/albums/:id", getAlbumByID)

	req, err := http.NewRequest("GET", "/albums/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var rcvd album
	json.Unmarshal(w.Body.Bytes(), &rcvd)
	if status := w.Code; status != http.StatusOK {
		t.Fatalf("handler returned wrong status code: wanted %v got %v",
			http.StatusOK, status)
	}
	if albums[1] != rcvd {
		t.Fatalf("handler return unexpected body: got %v want %v",
			rcvd, albums[1])
	}
}

func TestPostAlbum(t *testing.T) {
	r := SetUpRouter()
	r.POST("/albums", postAlbums)
	originalAlbumsLen := len(albums)

	var newAlbum = album{
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

	if len(albums) != originalAlbumsLen+1 {
		t.Fatal("Something went wrong: new album not added")
	}

	if albums[originalAlbumsLen] != newAlbum {
		t.Fatalf("expected new album to equal %v", newAlbum)
	}
}
