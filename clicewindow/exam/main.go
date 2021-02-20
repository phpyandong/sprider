package main

import (
	"fmt"
	"net/http"
	ratelimit "sprider/clicewindow"
)
//https://github.com/sunanzhi/ratelimit-go
func main() {
	var (
		limitTime   = 10
		bucketCount = 5
		limitCount  = 200
	)
	slideWindow, _ := ratelimit.Init(limitTime, bucketCount, limitCount)

	http.HandleFunc("/ratelimit", func(w http.ResponseWriter, r *http.Request) {
		err := slideWindow.Limiting()
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			fmt.Println(err.Error())
		} else {
			w.Write([]byte("httpserver"))
		}
	})
	fmt.Println("Starting server ...")
	http.ListenAndServe(":9090", nil)
}
