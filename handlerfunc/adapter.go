package handlerfunc

import (
	"net/http"

	"github.com/0angelic0/aws-lambda-go-api-proxy/httpadapter"
)

type HandlerFuncAdapter = httpadapter.HandlerAdapter

func New(handlerFunc http.HandlerFunc) *HandlerFuncAdapter {
	return httpadapter.New(handlerFunc)
}
