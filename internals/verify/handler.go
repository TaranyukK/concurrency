package verify

import (
	"fmt"
	"go/concurrency/configs"
	"go/concurrency/pkg/response"
	"net/http"
)

type VerifyHandlerDeps struct {
	*configs.Config
}

type VerifyHandler struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data := SendResponse{}
		response.Json(w, data, 200)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Register")
	}
}
