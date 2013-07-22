package web

import (
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", noData)
}
