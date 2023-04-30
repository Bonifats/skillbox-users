package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	h "students/pkg/http"
)

func TestUserCreate(t *testing.T) {
	for _, testUser := range []*h.User{
		{
			Name: "Bogdan",
			Age:  30,
		},
		{
			Name: "Sergey",
			Age:  27,
		},
	} {
		writer := makeRequest("POST", "/create", testUser)
		assert.Equal(t, http.StatusCreated, writer.Code)
	}
}

func TestUserUpdate(t *testing.T) {
	user := h.User{
		Name: "Anton",
		Age:  29,
	}

	writer := makeRequest("PUT", "/1", user)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestUserAttach(t *testing.T) {
	attach := h.Attach{
		SourceId: 1,
		TargetId: 2,
	}

	writer := makeRequest("POST", "/make_friends", attach)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestUserGetFriends(t *testing.T) {
	writer := makeRequest("GET", "/friends/1", nil)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func TestUserDelete(t *testing.T) {
	deleteId := h.Delete{
		Id: 2,
	}

	writer := makeRequest("DELETE", "/user", deleteId)
	assert.Equal(t, http.StatusOK, writer.Code)
}

func makeRequest(method, url string, body interface{}) *httptest.ResponseRecorder {
	router := chi.NewRouter()

	fmt.Println("localhost" + url)

	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(requestBody)

	request, err := http.NewRequest(method, "http://localhost"+url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(request.Header)

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	return writer
}
