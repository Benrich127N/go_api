package main

import (
	"log"
	"net/http"
)

type server struct {
	addr string
}

func (s *server) ServeHTTP(w http.ResponseWriter,r  *http.Request,)  {
	switch r.Method {
		case "GET":
			w.Write([]byte("index page"))
	}
}

func main() {
	s :=&server{addr:":8080"}
	if err := http.ListenAndServe(s.addr, s); err != nil {
		log.Fatal(http.ListenAndServe(s.addr, s))
	}
}