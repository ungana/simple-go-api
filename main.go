package main

import (
	"fmt"
	"net/http"

	"ungana.com/api/routes"
)

func main() {
	http.HandleFunc("POST /item", routes.CreateItemHandler)
	fmt.Println("🚀 Listening on port port 8080")
	http.ListenAndServe(":8080", nil)
}
