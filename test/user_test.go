package test

import (
	"chatross-api/internal/model"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {
	ClearUser()
	w := httptest.NewRecorder()

	requestBody := model.RegisterUserRequest{
		Username: "halludba",
		Password: "rahasia",
	}
	
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/user", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	app.ServeHTTP(w, request)


	bytes, err := io.ReadAll(w.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)
	

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, requestBody.Username, responseBody.Data.Username)
	assert.NotNil(t, responseBody.Data.CreateAt)
	assert.NotNil(t, responseBody.Data.UpdateAt)

}

func TestLogin(t *testing.T) {
	TestRegister(t)
	w := httptest.NewRecorder()

	requestBody := model.LoginUserRequest{
		Username: "halludba",
		Password: "rahasia",
	}
	
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/_login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	app.ServeHTTP(w, request)

	bytes, err := io.ReadAll(w.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.TokenResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)
	

	assert.Equal(t, 200, w.Code)
	assert.NotNil(t, responseBody.Data.AccessToken)
	assert.NotNil(t, responseBody.Data.RefreshToken)

}

func TestVerify(t *testing.T) {
	token := GetToken()
	w := httptest.NewRecorder()

	expectResponse := model.WebResponse[model.UserResponse]{
		Status: 200,
		Data: model.UserResponse{
			Username: "halludba",
		},
	}

	request := httptest.NewRequest(http.MethodGet, "/api/_verify", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", token)

	app.ServeHTTP(w, request)

	bytes, err := io.ReadAll(w.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.UserResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, expectResponse.Status, w.Code)
	assert.Equal(t, expectResponse.Data.Username, responseBody.Data.Username)
	assert.NotNil(t, responseBody.Data.CreateAt)
	assert.NotNil(t, responseBody.Data.UpdateAt)
}
