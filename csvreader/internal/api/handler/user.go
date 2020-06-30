package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"csvreader/internal/app/csv"
)

type handler func (w http.ResponseWriter, r *http.Request)

func UserCSV(up csv.UserProcessor) handler {
	return func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ret := up.Process(r.Context(), file)

		jsn, err := json.Marshal(ret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err = w.Write(jsn); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
