package server

import (
	"net/http"

	"github.com/go-chi/render"
)

func SendErrorJSON(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error) {
	render.Status(r, httpStatusCode)
	render.JSON(w, r, err.Error())
}
