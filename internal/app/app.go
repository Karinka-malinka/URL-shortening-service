package app

import (
	"time"

	"github.com/speps/go-hashids"
)

var ShortURL = make(map[string]string)

func ShorteningURL(longURL string) string {

	// получаем короткий url как хэш текущего времени
	hd := hashids.NewData()
	h, _ := hashids.NewWithData(hd)
	now := time.Now()
	urlID, _ := h.Encode([]int{int(now.Unix())})

	ShortURL[urlID] = string(longURL)

	return urlID
}

func ResolveURL(shortURL string) string {
	return ShortURL[shortURL]
}
