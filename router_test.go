package golanghttprouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, "Hello Http Router!")
	})

	request, _ := http.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	// if recorder.Code != http.StatusOK {
	// 	t.Errorf("Expected status code 200, got %d", recorder.Code)
	// }

	expectedBody := "Hello Http Router!"
	// if recorder.Body.String() != expectedBody {
	// 	t.Errorf("Expected body '%s', got '%s'", expectedBody, recorder.Body.String())
	// }

	assert.Equal(t, http.StatusOK, recorder.Code, "Expected status code to be 200")
	assert.Equal(t, expectedBody, recorder.Body.String(), "Expected response body to match")
}