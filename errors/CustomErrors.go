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
