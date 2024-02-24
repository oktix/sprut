package main

import (
	"fmt"
	"net/http"
	"sprut/internal/service"
)

func main() {
	server := service.NewServerImpl()
	err := http.ListenAndServe("localhost:80", server)
	if err != nil {
		fmt.Println("Sprut happened: ", err)
	}
}
