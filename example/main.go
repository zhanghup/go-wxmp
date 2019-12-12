package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet {
			data, err := ioutil.ReadAll(req.Body)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(data)
		} else {

		}

		res.Write([]byte("helloworld"))
	})
	http.ListenAndServe(":40018", nil)
}
