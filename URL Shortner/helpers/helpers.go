package helpers

import (
	"math/rand"
	"net/http"
	"os"
)

func ShortenURL(url string, r *http.Request) string {
	randomchars := os.Getenv("CHAR_SEQUENCE")
	newurl := r.URL.String() + string(randomchars[rand.Intn(len(randomchars))])
	return newurl
}
