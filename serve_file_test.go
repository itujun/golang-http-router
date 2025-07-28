package golanghttprouter

import (
	"embed"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/static/*filepath", http.FS(directory))

	request, _ := http.NewRequest("GET", "http://localhost:3000/static/hello.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	expectedBody := "Hello HttpRouter"
	assert.Equal(t, expectedBody, recorder.Body.String(), "Expected response body to match")
}

func TestServeFile2(t *testing.T) {
	router := httprouter.New()
	directory, _ := fs.Sub(resources, "resources")
	router.ServeFiles("/files/*filepath", http.FS(directory))

	request, _ := http.NewRequest("GET", "http://localhost:3000/files/bye.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, "bye", recorder.Body.String(), "Expected response body to match")
}