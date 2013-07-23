package web

import (
	"net"
	"bufio"
	"fmt"
	"net/http"
)

func AjaxHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	action := r.Form["type"]
	if len(action) > 0 {
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
		who := r.Form["who"]
		var who_text string
		if len(who) > 0 {
			who_text = who[0]
		} else {
			who_text = ""
		}
		where := r.Form["where"]
		var where_text string
		if len(where) > 0 {
			where_text = where[0]
		} else {
			where_text = ""
		}
		fmt.Fprintf(conn, action[0]+","+what_text+","+who_text+","+where_text)
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
