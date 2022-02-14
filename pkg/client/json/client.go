package js
// js - Json Storage

func NewJsonClient(path string) *JsonStorage {
	return &JsonStorage{pathStorage: path}
}