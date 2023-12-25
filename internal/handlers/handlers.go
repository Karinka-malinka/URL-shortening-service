package handlers

import (
	"io"
	"net/http"
	"strings"

	"github.com/URL-shortening-service/internal/app"
	"github.com/go-chi/chi/v5"
)

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func Run() error {

	r := chi.NewRouter()

	r.Get("/{id}", getURL)
	r.Post("/", shorteningURL)

	return http.ListenAndServe(`:8080`, r)
}

func getURL(w http.ResponseWriter, r *http.Request) {

	longURL := app.GetURL(chi.URLParam(r, "id"))
	if longURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

func shorteningURL(w http.ResponseWriter, r *http.Request) {

	if r.Body == http.NoBody {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	urlShort := "http://" + r.Host + "/" + app.ShorteningURL(string(body))

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(urlShort))
}

// функция webhook — обработчик HTTP-запроса
func mainHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:

		if r.Body == http.NoBody {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		body, err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		if len(body) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		urlShort := "http://" + r.Host + "/" + app.ShorteningURL(string(body))

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(urlShort))

	case http.MethodGet:
		uri := strings.Split(r.RequestURI, "/")
		if len(uri) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		longURL := app.GetURL(uri[1])
		if longURL == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
