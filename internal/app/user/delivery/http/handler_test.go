package http

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestHandler_createUser(t *testing.T) {
	inputData := []struct{
		inputJSON    string
		statusCode   int
		errorMessage string
	}{
		{
			inputJSON: `{"display_name":"", "email":"test@gmail.com"}`,
			statusCode: http.StatusBadRequest,
			errorMessage: "field must not be empty",
		},
		{
			inputJSON: `{"display_name":"test", "email":"testgmail.com"}`,
			statusCode: http.StatusBadRequest,
			errorMessage: "field email is not valid",
		},
		{
			inputJSON: `{"display_name":"test", "email":""}`,
			statusCode: http.StatusBadRequest,
			errorMessage: "field must not be empty",
		},
		{
			inputJSON: `{"display_name":"test", "email":"test@gmail.com"}`,
			statusCode: http.StatusCreated,
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

		log.Println(w.StatusCode, v.statusCode)
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

func TestHandler_getUser(t *testing.T)  {

}