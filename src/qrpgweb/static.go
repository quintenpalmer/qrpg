package qrpgweb

import (
	"net/http"
	"strings"
	"path/filepath"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w,r,filepath.Join("web",strings.Trim(r.URL.Path,"/")))
}
