package service

import (
	"net/http"

	"github.com/go-chi/render"
)

func (rs *restserver) sendErrorJSON(w http.ResponseWriter, r *http.Request, httpStatusCode int, err error) {
	rs.logger.Error(err)
	render.Status(r, httpStatusCode)
	render.JSON(w, r, err.Error())
}
