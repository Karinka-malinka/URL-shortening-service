package handlers

import (
	"io"
	"net/http"

	"github.com/URL-shortening-service/internal/app"
	"github.com/go-chi/chi/v5"
)

func ResolveURL(w http.ResponseWriter, r *http.Request) {

	longURL := app.GetURL(chi.URLParam(r, "id"))
	if longURL == "" {
		w.WriteHeader(http.StatusBadRequest)
		returng
	}
	http.Redirect(w, r, longURL, http.StatusTemporaryRedirect)
}

func ShorteningURL(baseURL string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		urlShort := baseURL + "/" + app.ShorteningURL(string(body))

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(urlShort))
	}
}
