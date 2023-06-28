package main

import (
	"demo/app"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){

	application := app.CreateApp()
	application.Router = mux.NewRouter()
	application.Setup()
	http.Handle("/",application.Router)
	log.Fatal(http.ListenAndServe(":8080",nil))

}
