package util

import (
	"log"
	"net/http"

	"github.com/ilgooz/eres"
	"github.com/ilgooz/formutils"
)

func ParseForm(w http.ResponseWriter, r *http.Request, out interface{}) (invalid bool) {
	invalids, err := formutils.ParseForm(r, out)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}

	return eres.New(w).SetFields(invalids).WeakSend()
}
