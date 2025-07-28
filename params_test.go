package golanghttprouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/user/:name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		name := params.ByName("name")
		fmt.Fprintf(w, "Hello %s!", name)
	})

	request, _ := http.NewRequest("GET", "http://localhost:3000/user/Lev", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	expectedBody := "Hello Lev!"
	assert.Equal(t, expectedBody, recorder.Body.String(), "Expected response body to match")
}