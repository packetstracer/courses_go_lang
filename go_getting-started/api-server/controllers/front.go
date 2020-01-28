package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers to avoid warnings
func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", uc)
	http.Handle("/users/", uc)
}

// EncodeResponseAsJSON implementation
func EncodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
