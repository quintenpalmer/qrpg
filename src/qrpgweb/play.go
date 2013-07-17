package qrpgweb

import (
	"net/http"
)

func PlayHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "play", noData)
}
