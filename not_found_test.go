package golanghttprouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterNotFound(t *testing.T) {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// http.Error(w, "Not Found", http.StatusNotFound)
		fmt.Fprint(w, "Gak Ketemu!")
	})

	request, _ := http.NewRequest("GET", "http://localhost:3000/test", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, "Gak Ketemu!", recorder.Body.String(), "Expected response body to match")
}