package web

import (
	"net"
	"bufio"
	"fmt"
	"net/http"
)

func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	zone := r.Form["type"]
	if len(zone) > 0 {
		conn, err := net.Dial("tcp","localhost:8000")
		if err != nil {
			send500(w,r)
			fmt.Println(err)
			return
		}
		what := r.Form["what"]
		var what_text string
		if len(what) > 0 {
			what_text = what[0]
		} else {
			what_text = ""
		}
		fmt.Fprintf(conn, zone[0]+"\n"+what_text)
		status, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			send500(w,r)
			fmt.Println(err)
			return
		}
		w.Write([]byte(status))
	} else {
		w.Write([]byte("You send me a version?"))
	}
}
