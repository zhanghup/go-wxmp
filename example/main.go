package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			values := req.URL.Query()
			signature := values.Get("signature")
			timestamp := values.Get("timestamp")
			nonce := values.Get("nonce")
			echostr := values.Get("echostr")

		case http.MethodPost:
		}


		res.Write([]byte("helloworld"))
	})
	http.ListenAndServe(":40018", nil)
}
