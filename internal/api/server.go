package api

import (
	"fmt"
	"net/http"
)

func Start() error {
	fmt.Println("Starting up at port 8080")
	return http.ListenAndServe(":8080", NewRouter())
}
