package gohome

import (
	"fmt"
	"net/http"
)

type Runner interface {
	Run() error
}

func restHandler(r Runner) http.HandlerFunc {
	id := ID(r)
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, id)
	}
}

func Run(r Runner) error {
	err := r.Run()
	if err != nil {
		return err
	}

	return http.ListenAndServe(":8080", restHandler(r))
}
