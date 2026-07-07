package randomapi

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

type RandomNumHandler struct{}

func NewRandomNumHandler(router *http.ServeMux) {
	handler := &RandomNumHandler{}
	router.HandleFunc("/roll", handler.RandomNum())
}

func (handler *RandomNumHandler) RandomNum() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.IntN(6)+1)
	}
}
