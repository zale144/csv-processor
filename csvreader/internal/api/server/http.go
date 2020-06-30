package server

import (
	"fmt"
	"log"
	"net/http"
	"csvreader/internal/api/handler"
	"csvreader/internal/app/csv"
)

func ServeHTTP(port int64, up csv.UserProcessor)  {
	http.HandleFunc("/csv/user", handler.UserCSV(up))
	log.Println("HTTP Listening on localhost:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
