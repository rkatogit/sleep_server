package main

import (
	"fmt"
	"github.com/schollz/progressbar"
	"log"
	"net/http"
	"strconv"
	"time"
)

func count(s int) {
	bar := progressbar.New(s)
	for i := 0; i < s; i++ {
		bar.Add(1)
		time.Sleep(1 * time.Second)
	}
}

func HelloServer(w http.ResponseWriter, req *http.Request) {
	var t int
	t, _ = strconv.Atoi(req.URL.Query().Get("t"))
	s := time.Now()
	fmt.Println("\nRequest Received!")
	count(t)
	w.Header().Set("Content-Type", "text/plain")
	f := time.Now()
	d := fmt.Sprintf("[Request Received:]\n%s\n[Response sent:]\n%s", s.String(), f.String())
	w.Write([]byte(d))
	fmt.Println("\nResponse Finish!")
}

func main() {
	http.HandleFunc("/sleep", HelloServer)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
