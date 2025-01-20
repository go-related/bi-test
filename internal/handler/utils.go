package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetParameter(r *http.Request, paramName string) string {
	return chi.URLParam(r, paramName)
}
