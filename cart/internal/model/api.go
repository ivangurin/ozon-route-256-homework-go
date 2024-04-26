package model

import "net/http"

type HttpAPIDescription struct {
	Handlers HttpApiHandlers
}

type HttpAPIHandler struct {
	Pattern string
	Handler http.HandlerFunc
}

type HttpApiHandlers []*HttpAPIHandler
