package model

import "net/http"

type HttpAPIDescription struct {
	Handlers HttpApiHandlers
}

type HttpAPIHandler struct {
	Pattern string
	Handler func(http.ResponseWriter, *http.Request)
}

type HttpApiHandlers []*HttpAPIHandler
