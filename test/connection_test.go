package test

import (
	"chatross-api/internal/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T){
	w := httptest.NewRecorder()

	request, err := http.NewRequest("GET", "/ping", nil)
	assert.Nil(t, err)

	expectResponse := model.WebResponse[string]{
		Status: 200,
		Data: "pong",
	}

	app.ServeHTTP(w, request)

	bytes, err := io.ReadAll(w.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[string])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)
	
	assert.Equal(t, expectResponse.Status, responseBody.Status)
	assert.Equal(t, expectResponse.Data, responseBody.Data)
}

