package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_BKS/controllers"
	database "github.com/go-park-mail-ru/2023_1_BKS/db"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello)")
	})
	http.HandleFunc("/signup", controllers.Signup)

	database.InitDB()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
