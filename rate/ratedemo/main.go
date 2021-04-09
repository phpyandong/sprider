package main

import (
	"time"
	"net/http"
	rate2 "sprider/rate"
	"log"
)

func main() {
	limiter := rate2.NewLimiter(rate2.Every(1000*time.Millisecond), 10)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {// do something
			log.Println("say hello")
			w.Write([]byte("say hello"))
		}else{
			w.Write([]byte("wait ......"))

		}
	})
	_ = http.ListenAndServe(":13100", nil)
}


