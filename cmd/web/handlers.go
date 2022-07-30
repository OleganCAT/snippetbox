package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//Создается функция-обработчик домашней страницы "home"
func home(w http.ResponseWriter, r *http.Request) {
	//Проверка на catch-all
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Привет из Snippetbox"))
}

//Создается функция-обработчик для страницы "showSnippet которая извлекает параметр id из URL"
func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d...", id)
}

//Создается функция-обработчик для страницы "creatSnippet" которая отвечает только на POST запросы
func creatSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Метод запрещён", 405)
		return
	}
	w.Write([]byte("Создаёт заметки в  Snippetbox"))
}
