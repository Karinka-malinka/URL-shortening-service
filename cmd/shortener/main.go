package main

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/speps/go-hashids"
)

var ShortUrl = make(map[string]string)

func main() {

	if err := run(); err != nil {
		panic(err)
	}

}

// функция run будет полезна при инициализации зависимостей сервера перед запуском
func run() error {

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, mainHandler)

	return http.ListenAndServe(`:8080`, mux)
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

		urlShort := "http://" + r.Host + "/" + shorteningUrl(string(body))

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(urlShort))

	case http.MethodGet:
		uri := strings.Split(r.RequestURI, "/")
		if len(uri) < 2 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		longUrl := getUrl(uri[1])
		if longUrl == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, longUrl, http.StatusTemporaryRedirect)
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func shorteningUrl(longUrl string) string {

	// получаем короткий url как хэш текущего времени
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	urlID, _ := h.Encode([]int{int(now.Unix())})

	ShortUrl[urlID] = string(longUrl)

	return urlID
}

func getUrl(shortUrl string) string {
	return ShortUrl[shortUrl]
}

/*Сервер должен быть доступен по адресу http://localhost:8080 и предоставлять два эндпоинта:
Эндпоинт с методом POST и путём /. Сервер принимает в теле запроса строку URL как text/plain и возвращает ответ с кодом 201 и сокращённым URL как text/plain.
Пример запроса к серверу:
POST / HTTP/1.1
Host: localhost:8080
Content-Type: text/plain

https://practicum.yandex.ru/

Пример ответа от сервера:
HTTP/1.1 201 Created
Content-Type: text/plain
Content-Length: 30

http://localhost:8080/EwHXdJfB

Эндпоинт с методом GET и путём /{id}, где id — идентификатор сокращённого URL (например, /EwHXdJfB). В случае успешной обработки запроса сервер возвращает ответ с кодом 307 и оригинальным URL в HTTP-заголовке Location.
Пример запроса к серверу:
GET /EwHXdJfB HTTP/1.1
Host: localhost:8080
Content-Type: text/plain

Пример ответа от сервера:
HTTP/1.1 307 Temporary Redirect
Location: https://practicum.yandex.ru/

На любой некорректный запрос сервер должен возвращать ответ с кодом 400.
*/
