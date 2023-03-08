package main

import (
	"fmt"
	"get/down"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: fmt.Sprintf(":8000"),
		Handler: func() http.Handler {
			mux := http.NewServeMux()

			mux.HandleFunc("/test", func(res http.ResponseWriter, req *http.Request) {
				url := req.FormValue("url")
				down.StratDu(url)
			},
			)
			return mux
		}(),
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("%v", err)
		fmt.Println("Hello world")
	} else {
		fmt.Println("Server closed!")
	}

}
