package gohome

import (
	"net/http"
)

// NewEnvironment creates a new Environment object.
func NewEnvironment() *Environment {
	return &Environment{}
}

// Environment is the root datatype
type Environment struct {
}

// Run starts and runs the environment until shutdown or crash.
func (e *Environment) Run(address string) error {
	return http.ListenAndServe(address, http.HandlerFunc(e.handleHTTP))
}

func (e *Environment) handleHTTP(w http.ResponseWriter, r *http.Request) {
}
