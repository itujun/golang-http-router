package golanghttprouter

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		// http.Error(w, fmt.Sprintf("Panic occurred: %v", err), http.StatusInternalServerError)
		fmt.Fprint(w, "Panic: ", err)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		panic("Ups!")
	})

	request, _ := http.NewRequest("GET", "http://localhost:3000/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	assert.Equal(t, "Panic: Ups!", recorder.Body.String(), "Expected response body to match")
}