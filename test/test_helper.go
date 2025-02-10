package test

import (
	"chatross-api/internal/entity"
	"chatross-api/internal/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

func ClearUser() {
	err := db.Where("id is not null").Delete(&entity.User{}).Error
	if err != nil {
		log.Fatalf("Failed clear user data : %+v", err)
	}
}

func GetToken() string {
	w := httptest.NewRecorder()
	
	requestBody := model.LoginUserRequest{
		Username: "halludba",
		Password: "rahasia",
	}

	bodyJson, err := json.Marshal(requestBody)
	if err != nil {
		panic("Let's TestLogin bro!")
	}

	request := httptest.NewRequest(http.MethodPost, "/api/_login", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	app.ServeHTTP(w, request)
	bytes, err := io.ReadAll(w.Body)
	if err != nil {
		panic("Let's TestLogin bro!")
	}

	responseBody := new(model.WebResponse[model.TokenResponse])
	err = json.Unmarshal(bytes, responseBody)
	if err != nil {
		panic("Let's TestLogin bro!")
	}

	return responseBody.Data.AccessToken
}