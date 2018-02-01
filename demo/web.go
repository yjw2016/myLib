package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json" //序列化
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("%.2f刀", d) }

type MyHandler map[string]dollars

func (self MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		bytes, _ := json.Marshal(self)
		fmt.Fprintf(w, "%s\n", bytes)

	case "/price":
		//item := req.URL.Query().Get("item")
		item := req.Form.Get("item")
		price, ok := self[item]

		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "no such item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such page: %s\n", req.URL)
	}
}

func main() {
	var handler MyHandler = MyHandler{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe(":4000", handler))
}
