package golanghttprouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterPatternNamedParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/category/:name/product/:id", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		category := params.ByName("name")
		id := params.ByName("id")
		text := "Category: " + category + ", Product ID: " + id
		fmt.Fprint(w, text)
	})

	request, _ := http.NewRequest("GET", "http://localhost:3000/category/elektronik/product/123", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	expectedBody := "Category: elektronik, Product ID: 123"
	assert.Equal(t, expectedBody, recorder.Body.String(), "Expected response body to match")
}

func TestRouterPatternCatchAllParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		fmt.Fprintf(w, "Images: %s", image)
	})

	request, _ := http.NewRequest("GET", "http://localhost:3000/images/small/profile.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	expectedBody := "Images: /small/profile.png"
	assert.Equal(t, expectedBody, recorder.Body.String(), "Expected response body to match")
}