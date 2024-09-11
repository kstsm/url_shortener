package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "url_shortener"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

var storage map[int]User
var index int

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	for _, value := range storage {
		users = append(users, value)
	}
	response, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user, exists := storage[id]
	if !exists {
		mJson, err := json.Marshal(map[string]string{
			"message": fmt.Sprintf("Пользователь с ID = %d не найден", id),
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(mJson)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, exists := storage[id]
	if !exists {
		mJson, err := json.Marshal(map[string]string{
			"message": fmt.Sprintf("Пользователь с ID = %d не найден", id),
		})
		if err != nil {
			fmt.Println(err.Error())
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write(mJson)
		return
	}

	delete(storage, id)

	response, err := json.Marshal(map[string]string{
		"message": fmt.Sprintf("Пользователь с ID = %d удален", id),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Write(response)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.ID = index
	user.CreatedAt = time.Now()
	storage[index] = user
	index++

	_, err = db.Query(`
insert into links (original, shortened)
values ($1,$2);
`, user.Name, generateRandomString(6))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func patchUser(w http.ResponseWriter, r *http.Request) {

}

func putUser(w http.ResponseWriter, r *http.Request) {
	/*name := chi.URLParam(r, "name")
	  id := chi.URLParam(r, "id")
	  idToStr, _ := strconv.Atoi(id)
	  if _, patch := storage[idToStr]; patch {
	  	storage[idToStr] = name
	  	w.Write([]byte("Значение изменино"))
	  } else {
	  	w.Write([]byte("Ключ не найден "))
	  }*/
}

/*func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	delete(storage, id)
	w.Write([]byte("Удалено"))
}*/

var db *sql.DB

func init() {
	storage = make(map[int]User)
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer conn.Close()

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	db = conn
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/user", getUsers)
	r.Get("/user/{id}", getUserByID)
	r.Post("/user", createUser)
	r.Patch("/user/{id}", patchUser)
	r.Put("/user/{id}", putUser)
	r.Delete("/user/{id}", deleteUser)
	fmt.Println("Epta blya")
	http.ListenAndServe(":8090", r)
}
