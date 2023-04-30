package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
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
		response := makeRequest("POST", "/create", testUser)
		assert.Equal(t, http.StatusCreated, response.StatusCode)
	}
}

func TestUserUpdate(t *testing.T) {
	user := h.User{
		Name: "Anton",
		Age:  29,
	}

	response := makeRequest("PUT", "/1", user)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestUserAttach(t *testing.T) {
	attach := h.Attach{
		SourceId: 1,
		TargetId: 2,
	}

	response := makeRequest("POST", "/make_friends", attach)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestUserGetFriends(t *testing.T) {
	response := makeRequest("GET", "/friends/1", nil)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func TestUserDelete(t *testing.T) {
	deleteId := h.Delete{
		Id: 2,
	}

	response := makeRequest("DELETE", "/user", deleteId)
	assert.Equal(t, http.StatusOK, response.StatusCode)
}

func makeRequest(method, url string, body interface{}) *http.Response {
	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	request, err := http.NewRequest(method, "http://127.0.0.1"+url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// close response body
	err = response.Body.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(string(responseBody))

	return response
}
