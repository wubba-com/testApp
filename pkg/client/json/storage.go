package js

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type JsonStorage struct {
	pathStorage string
}

func (js *JsonStorage) path() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Printf(err.Error())
		return err.Error()
	}

	return filepath.Join(dir, js.pathStorage)
}

func (js *JsonStorage) Read() ([]byte, error) {
	return ioutil.ReadFile(js.path())
}

func (js *JsonStorage) Encode(v interface{}) error {
	b, err := js.Read()
	err = json.Unmarshal(b, v)
	if err != nil {
		return err
	}

	return nil
}

func (js *JsonStorage) Write(b []byte) (int, error) {
	err := ioutil.WriteFile(js.path(), b, fs.ModePerm)
	if err != nil {
		return 0, err
	}

	return len(b), nil
}

func (js *JsonStorage) Decode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}