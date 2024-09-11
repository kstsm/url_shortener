package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"url_shortener/internal/service"
)

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/{shortened}", getOriginalByShortened)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/", getLinks)
	apiRouter.Get("/{id}", getLinkByID)
	apiRouter.Post("/link", createLink)
	apiRouter.Patch("/{id}", patchLink)
	apiRouter.Put("/{id}", putLink)
	apiRouter.Delete("/{id}", deleteLink)

	r.Mount("/link", apiRouter)

	return r
}

func getOriginalByShortened(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	shortened := chi.URLParam(r, "shortened")

	original, err := service.GetOriginalByShortened(shortened)
	if err != nil {
		mJson, _ := json.Marshal(map[string]string{"message": "Ошибка при получении ссылки"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mJson)
		return
	}

	http.Redirect(w, r, original, http.StatusSeeOther)
}

func getLinks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	links, err := service.GetLinks()
	if err != nil {
		mJson, _ := json.Marshal(map[string]string{"message": "Ошибка при получении ссылок"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mJson)
		return
	}

	response, err := json.Marshal(links)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func getLinkByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	linkID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	link, err := service.GetLinkByID(linkID)
	if err != nil {
		mJson, _ := json.Marshal(map[string]string{"message": "Ошибка при получении ссылки"})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mJson)
		return
	}

	response, err := json.Marshal(link)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func deleteLink(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
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

	w.Write(response)*/
}

func createLink(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
		var user models.User
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
		w.Write(response)*/
}

func patchLink(w http.ResponseWriter, r *http.Request) {

}

func putLink(w http.ResponseWriter, r *http.Request) {
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
