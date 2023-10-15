package errors

import (
	"net/http"

	"github.com/charmbracelet/log"
)


func BadRequest(w http.ResponseWriter, err error, code int, msg string) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(msg)
	}
}

func InternalServerError(w http.ResponseWriter, err error, code int, msg string) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(msg)
	}
}

func Err(e error, msg string) {
	if e != nil {
		log.Error(msg)
	}
}

func Fatal(e error, msg string) {
	if e != nil {
		log.Fatal(e, msg)
	}
}

func GenericError[T any](e error, msg string, data T) (T, error) {
	if e != nil {
		log.Error(e, msg)
		return data, e
	}
	return data, nil
}
