package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

func homepage() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//index.vue
	}
}

func article() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//article
	}
}

func manageArticle() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//Manage Article
	}
}

func profile() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//profile
	}
}

func login() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//login
	}
}

func register() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		//register
	}
}



func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homepage())
	return mux
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main(){
	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", GetMux())
}