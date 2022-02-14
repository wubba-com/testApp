package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHandler_createUser(t *testing.T) {
	inputData := []struct {
		inputJSON    string
		statusCode   int
		errorMessage string
	}{
		{
			inputJSON:    `{"display_name":"", "email":"test@gmail.com"}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "field must not be empty",
		},
		{
			inputJSON:    `{"display_name":"test", "email":"testgmail.com"}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "field email is not valid",
		},
		{
			inputJSON:    `{"display_name":"test", "email":""}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "field must not be empty",
		},
		{
			inputJSON:    `{"display_name":"test", "email":"test@gmail.com"}`,
			statusCode:   http.StatusCreated,
			errorMessage: "",
		},
	}

	for _, v := range inputData {
		clientTest := &http.Client{
			Timeout: 5 * time.Second,
		}

		r, err := http.NewRequest(http.MethodPost, "http://localhost:3333/api/v1/users", bytes.NewBufferString(v.inputJSON))
		r.Header.Set("Content-Type", "application/json")
		if err != nil {
			log.Fatalln(err)
		}
		w, err := clientTest.Do(r)
		if err != nil {
			log.Fatalf("client do: %s", err.Error())
		}

		assert.Equal(t, v.statusCode, w.StatusCode)
		if v.statusCode != http.StatusCreated {
			m := make(map[string]interface{})
			err := json.NewDecoder(w.Body).Decode(&m)
			if err != nil {
				log.Fatalf("Cannot convert to %s", err.Error())
			}
			assert.Equal(t, v.errorMessage, m["error"])
		}
	}
}

func TestHandler_getUser(t *testing.T) {
	client := http.Client{Timeout: 1 * time.Second}

	inputData := struct {
		inputJSON  string
		statusCode int
	}{
		inputJSON:  `{"display_name":"test", "email":"test@gmail.com"}`,
		statusCode: http.StatusOK,
	}

	id := seedUser()

	r, err := http.NewRequest(http.MethodGet, "http://localhost:3333/api/v1/users/"+id, nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")

	w, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	// декодируем ответ
	u := make(map[string]string)
	err = json.NewDecoder(w.Body).Decode(&u)

	defer w.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	// сравниваем
	assert.Equal(t, inputData.statusCode, w.StatusCode)
	assert.Equal(t, "test", fmt.Sprintf("%s", u["display_name"]))

	// Удаляем нового пользователя
	deleteUser(id)
}

func TestHandler_updateUser(t *testing.T) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	inputData := []struct {
		inputJSON    string
		statusCode   int
		errorMessage string
	}{
		{
			inputJSON:    `{"display_name":"", "email":"test@gmail.com"}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "field must not be empty",
		},
		{
			inputJSON:    `{"display_name":"test", "email":"testgmail.com"}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "field email is not valid",
		},
		{
			inputJSON:    `{"display_name":"test", "email":""}`,
			statusCode:   http.StatusBadRequest,
			errorMessage: "field must not be empty",
		},
	}

	id := seedUser()
	for _, v := range inputData {
		r, err := http.NewRequest(http.MethodPatch, "http://localhost:3333/api/v1/users/"+id, bytes.NewBufferString(v.inputJSON))
		if err != nil {
			log.Fatal(err)
		}
		r.Header.Set("Content-Type", "application/json")

		w, err := client.Do(r)
		if err != nil {
			log.Fatal(err)
		}
		assert.Equal(t, v.statusCode, w.StatusCode)

		if v.statusCode != http.StatusNoContent {
			m := make(map[string]interface{})
			err := json.NewDecoder(w.Body).Decode(&m)
			if err != nil {
				log.Fatalf("Cannot convert to %s", err.Error())
			}
			assert.Equal(t, v.errorMessage, m["error"])
		}
	}
	deleteUser(id)
}

func TestHandler_deleteUser(t *testing.T) {
	client := http.Client{
		Timeout: 1 * time.Second,
	}
	id := seedUser()
	desiredResult := struct {
		ID         string
		statusCode int
	}{
		ID:         id,
		statusCode: http.StatusNoContent,
	}

	r, err := http.NewRequest(http.MethodDelete, "http://localhost:3333/api/v1/users/"+id, nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")

	w, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, desiredResult.statusCode, w.StatusCode)
}

func seedUser() string {
	client := http.Client{Timeout: 1 * time.Second}

	inputData := struct {
		inputJSON    string
		statusCode   int
		errorMessage string
	}{
		inputJSON:    `{"display_name":"test", "email":"test@gmail.com"}`,
		statusCode:   http.StatusOK,
		errorMessage: "",
	}
	r, err := http.NewRequest(http.MethodPost, "http://localhost:3333/api/v1/users/", bytes.NewBufferString(inputData.inputJSON))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")

	w, err := client.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	// декодируем ответ
	res := make(map[string]float64)
	err = json.NewDecoder(w.Body).Decode(&res)
	if err != nil {
		log.Println(res)
		log.Fatal(err)
	}

	// Получаем нового пользователя
	id := fmt.Sprintf("%d", int(res["user_id"]))

	return id
}

func deleteUser(id string) {
	r, err := http.NewRequest(http.MethodDelete, "http://localhost:3333/api/v1/users/"+id, nil)
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")
}
