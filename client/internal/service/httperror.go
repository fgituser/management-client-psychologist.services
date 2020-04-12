package service

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime"
	"strings"

	"github.com/go-chi/render"
)

const (
	ErrInternal = "Internal Server Error"
	ErrNoAccess = "No Access"
	ErrNoAttachment = "Not attachment to psychologist"
)

func (rs *restserver) sendErrorJSON(w http.ResponseWriter, r *http.Request, httpStatusCode int, msg string, err error) {
	rs.logger.Errorf(errDetailsMsg(r, httpStatusCode, err, msg, httpStatusCode))
	render.Status(r, httpStatusCode)
	render.JSON(w, r, map[string]interface{}{"error": msg})
}

func errDetailsMsg(r *http.Request, httpStatusCode int, err error, details string, errCode int) string {
	uinfoStr := ""
	q := r.URL.String()
	if qun, e := url.QueryUnescape(q); e == nil {
		q = qun
	}

	srcFileInfo := ""
	if pc, file, line, ok := runtime.Caller(2); ok {
		fnameElems := strings.Split(file, "/")
		funcNameElems := strings.Split(runtime.FuncForPC(pc).Name(), "/")
		srcFileInfo = fmt.Sprintf("[%s:%d %s]", strings.Join(fnameElems[len(fnameElems)-3:], "/"),
			line, funcNameElems[len(funcNameElems)-1])
	}

	return fmt.Sprintf("%s - %v - %d (%d) - %s%s - %s",
		details, err, httpStatusCode, errCode, uinfoStr, q, srcFileInfo)
}
