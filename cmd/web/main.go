package main

import (
	"log"
	"net/http"
)

func main() {
	//http.NewServeMux() для инициализации нового рутера
	//Регистрируется как обработчик для URL-шаблонов.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", creatSnippet)

	//Функция http.ListenAndServe() используется для запуска нового веб-сервера.
	log.Println("Запуск веб-сервера http://127.0.0.1:4000")
	err := http.ListenAndServe("127.0.0.1:4000", mux)
	log.Fatal(err)
}
