package qrpgweb

import (
	"net/http"
	"fmt"
)

func Main() {
	ParseTemplates()
	http.HandleFunc("/",IndexHandler)
	http.HandleFunc("/static/",StaticHandler)
	http.HandleFunc("/play/",PlayHandler)
	http.HandleFunc("/ajax/",AjaxHandler)
	port := ":8080"
	fmt.Printf("Serving content on %s\n",port)
	panic(http.ListenAndServe(port, nil))
}