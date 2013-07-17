package qrpgweb

import (
	"net/http"
)

func send500(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("500 - Internal Server Error"))
}

func send400(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("400 - Bad Request"))
}

