package test

import (
	"chatross-api/internal/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T){
	w := httptest.NewRecorder()
	
	exampleResponse := model.WebResponse[string]{
		Status: 200,
		Data: "pong",
	}

	resJson, err := json.Marshal(exampleResponse)
	assert.Nil(t, err)

	req, err := http.NewRequest("GET", "/ping", nil)
	assert.Nil(t, err)

	app.ServeHTTP(w, req)

	
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(resJson), w.Body.String())
}