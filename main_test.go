package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
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

	fmt.Println("http://127.0.0.1" + url)

	requestBody, _ := json.Marshal(body)
	request, _ := http.NewRequest(method, "http://127.0.0.1"+url, bytes.NewBuffer(requestBody))

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)

	return writer
}
