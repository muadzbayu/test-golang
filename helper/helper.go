package helper

import (
	"strings"

	"github.com/google/uuid"
)

type WebResponse[T any] struct {
	Code    string `json:"code"`
	Message T      `json:"message"`
	Data    T      `json:"data"`
}

// untuk reponse dalam bentuk object kosong
func EmptyObject() interface{} {
	return make(map[string]interface{})
}

// untuk reponse dalam bentuk array kosong
func EmptyArray() interface{} {
	return []interface{}{}
}

// Generate a unique session ID
func GenerateSessionID() string {
	randomUUID := uuid.New()
	return strings.Replace(randomUUID.String(), "-", "", -1)
}
