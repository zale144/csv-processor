package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	db := newMockDB()
	serveHTTP(db)
}

func serveHTTP(db *mockDB)  {
	port := os.Getenv("CRM_HTTP_PORT")
	http.HandleFunc("/clear", clearUsersHandler(db))
	http.HandleFunc("/save", saveUsersHandler(db))
	http.HandleFunc("/get", getUsersHandler(db))
	log.Println("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type handler func (w http.ResponseWriter, r *http.Request)

func saveUsersHandler(db *mockDB) handler {
	return func (w http.ResponseWriter, r *http.Request) {

		users := make([]*user, 0)

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			http.Error(w, "no users were provided", http.StatusBadRequest)
			return
		}

		firstID, lastID := users[0].Id, users[len(users)-1].Id

		log.Printf("received users with id from %d to %d\n", firstID, lastID)

		if err := maybeError(); err != nil {
			log.Println("an unexpected error occurred")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, u := range users {
			db.save(u.Id, u)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
	}
}

func getUsersHandler(db *mockDB) handler {
	return func (w http.ResponseWriter, r *http.Request) {

		jsnUsr, err := json.Marshal(db.get())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		if _, err = w.Write(jsnUsr); err != nil {
			log.Println(err)
		}
	}
}

func clearUsersHandler(db *mockDB) handler {
	return func (w http.ResponseWriter, r *http.Request) {
		db.clear()
	}
}

func maybeError() error {
	results := make([]error, 3) // 33% chance of failing
	results[0] = errors.New("an unexpected error occurred")

	rand.Seed(time.Now().UnixNano())

	return results[rand.Intn(len(results))]
}

type user struct {
	Id                   int64    `json:"id"`
	FirstName            string   `json:"first_name"`
	LastName             string   `json:"last_name"`
	Email                string   `json:"email"`
	Phone                string   `json:"phone"`
}

type mockDB struct {
	users sync.Map
}

func newMockDB() *mockDB {
	return &mockDB{
		users: sync.Map{},
	}
}

func (m *mockDB) save(id int64, u *user)  {
	if _, ok := m.users.Load(id); ok {
		log.Println("DUPLICATE for ID: ", id)
	}
	m.users.Store(id, u)
}

func (m *mockDB) get() (u []*user) {
	m.users.Range(func(key, value interface{}) bool {
		u = append(u, value.(*user))
		return true
	})
	return
}

func (m *mockDB) clear()  {
	m.users = sync.Map{}
}
