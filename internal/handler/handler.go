package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"strconv"
	"url_shortener/internal/cerrors"
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
		mJson, err := json.Marshal(map[string]string{"message": "Ошибка при получении ссылки"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

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
		mJson, err := json.Marshal(map[string]string{"message": "Ошибка при получении ссылок"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mJson)
		return
	}

	response, err := json.Marshal(links)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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
		mJson, err := json.Marshal(map[string]string{"message": "Ошибка при получении ссылки"})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(mJson)
		return
	}

	response, err := json.Marshal(link)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
}

func deleteLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	linkID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = service.DeleteLink(linkID)
	if err != nil {
		if errors.Is(cerrors.ErrNotFound, err) {
			mJson, err := json.Marshal(map[string]string{"message": fmt.Sprintf("Cсылка с ID = %d не найдена", linkID)})
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusNotFound)
			w.Write(mJson)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]int{"id": linkID})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(response)
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
